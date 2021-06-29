// Package executor provides all functions for executing the audits.
// This includes all functions to analyze the tasks and their commands/scripts,
// execute them and compare their outputs with the expected values.
//
package executor

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"Auditheia/memory"
	"Auditheia/memory/constants"
	"Auditheia/output"

	"github.com/sirupsen/logrus"
)

// ExecuteAudit executes the given memory.Audit, calls executeTask with every task saved in the given memory.Audit
// sets the audit status related to the last executed Task and calls memory.Audit.SetExecuted()
//
// returns an error if either audit is nil or audit.Tasks is nil or empty
//
func ExecuteAudit(audit *memory.Audit) error {
	if audit == nil {
		return fmt.Errorf("no audit specified: %v", audit)
	} else if audit.Tasks == nil || len(audit.Tasks) == 0 {
		return fmt.Errorf("empty task list")
	}
	for n := range audit.Tasks {
		output.Log.Infof("Executing Task %v...", n)
		err := executeTask(&audit.Tasks[n])
		if err != nil {
			output.Log.Errorln(err)
		}
	}

	audit.SetExecuted()

	return nil
}

// executeTask executes the given memory.Task by first determining the tasks type and calling appropriate handlers
// for the types 'constants.TASK_TYPE_COMMAND' and 'constants.TASK_TYPE_SCRIPT'.
// Upon task completion or error executeTask checks the tasks status and selects an appropriate subsequent task,
// which then is executed.
//
// returns an Error the given task is nil. Errors created due to unknown task type and unknown task status are
// filtered by executeTask and directly logged.
//
func executeTask(task *memory.Task) error {
	if task == nil {
		return fmt.Errorf("no task specified (nil)")
	}

	task.SetExecuted()
	var err error
	switch task.Type {
	case constants.TASK_TYPE_COMMAND:
		output.Log.Debugf("identified type: %s", task.Type)
		err = executeTaskTypeCommand(task)
	case constants.TASK_TYPE_SCRIPT:
		output.Log.Debugf("identified type: %s", task.Type)
		err = executeTaskTypeScript(task)
	default:
		err = fmt.Errorf("unknown task type: '%s'", task.Type)
		task.SetStatus(constants.STATUS_ERROR)
		task.SetStatusMessage(err.Error())
	}
	if err != nil {
		output.Log.Errorln(err)
		err = nil
	}
	var nextTask *memory.Task
	switch task.GetStatus() {
	case constants.STATUS_SUCCESS:
		nextTask = task.OnSuccess
	case constants.STATUS_FAIL:
		nextTask = task.OnFail
	case constants.STATUS_ERROR:
		nextTask = task.OnError
	default:
		output.Log.Warningf("unknown task status: %s, cannot determine subsequent task", task.GetStatus())
		nextTask = nil
	}
	if nextTask != nil {
		output.Log.Infof("Executing subsequent task")
		err = executeTask(nextTask)
	}
	return err
}

// executeTaskTypeCommand executes the given memory.Task of type constants.TASK_TYPE_COMMAND.
//
// returns an error if the given task is nil or buildCommandsFromString returns either an error or an empty list of
// exec.Cmd
//
func executeTaskTypeCommand(task *memory.Task) error {
	if task == nil {
		return fmt.Errorf("no task (nil)")
	}

	cmdList, err := buildCommandsFromString(task.Execute)
	if err != nil {
		task.SetStatus(constants.STATUS_ERROR)
		return err
	}
	if cmdList == nil || len(cmdList) == 0 {
		task.SetStatus(constants.STATUS_ERROR)
		return fmt.Errorf("commandList nil or empty")
	}

	var tmp string
	for i, s := range cmdList {
		fields := logrus.Fields{
			"command": strings.Join(s.Args, " "),
		}
		if i > 0 {
			s.Stdin = strings.NewReader(tmp) // set the old output as input for the new command
		}
		result, err := runCmd(s)
		result = cleanShellOutputs(result)
		tmp = result
		if err != nil {
			output.Log.WithFields(fields).Errorln(err)
			task.SetStatus(constants.STATUS_ERROR)
			task.SetStatusMessage(result)
			break
		}
		if i == len(cmdList)-1 { // check if it's the last command
			output.Log.Debugf("Saving Result of last Command ('%s')...", strings.Join(s.Args, " "))
			task.SetResult(result) // save result
			output.Log.Debugf("Checking Result of last Command ('%s')...", strings.Join(s.Args, " "))
			if checkResult(task.GetResult(), task.Expected, task.Operator) {
				task.SetStatus(constants.STATUS_SUCCESS)
			} else {
				task.SetStatus(constants.STATUS_FAIL)
			}
		}
		// put every output of every command into artefact if task != nil
		if memory.OptionsList.Verbosity >= constants.VERBOSITY_DEBUG {
			saveArtefact(task, s, []byte(result))
		} else if i == 0 { // put output of first command into artefact
			saveArtefact(task, s, []byte(result))
		}
	}

	return nil
}

// runCmd executes the given exec.Cmd provided by the task or javascript, errors get logged in the logfile
//
// returns an error if s.CombinedOutput returns an error and always a string containing the stdout of the shell.
//
func runCmd(s exec.Cmd) (string, error) {

	output.Log.Debugf("Executing Command '%s'", strings.Join(s.Args, " "))
	stdout, err := s.CombinedOutput() // get both output + error code

	output.Log.Traceln("Checking for Errors of Command...")
	if err != nil {
		o := fmt.Sprintf("%v: %s", err, string(stdout)) // set error code + output of the error together
		return o, errors.New(o)
	}
	return string(stdout), nil
}

// checkResult checks whether the given result value matches any of the given expected values according to the
// given operator.
//
// returns true if the value matches any result according to the operator, false any other way
//
//  checkResult("result", "[first_expected, second_expected]", "not equals")
//
func checkResult(result string, expected []string, operator string) (matching bool) {
	output.Log.Infoln("Checking Task result...")
	result = strings.TrimSpace(result) // remove all leading and trailing white spaces

	for i, expect := range expected {
		expect = strings.TrimSpace(expected[i]) // remove all leading and trailing white spaces
		switch operator {
		case constants.OPERATOR_EQUALS:
			if strings.EqualFold(result, expect) { // check equal case sensitive ignored
				return true
			}
		case constants.OPERATOR_NOT_EQUALS:
			if !(strings.EqualFold(result, expect)) { // check not equal case sensitive ignored
				return true
			}
		case constants.OPERATOR_CONTAINS:
			if strings.Contains(result, expect) { // check string contains
				return true
			}
		case constants.OPERATOR_NOT_CONTAINS:
			if !strings.Contains(result, expect) {
				return true
			}
		case constants.OPERATOR_LESSER:
			iResult, err1 := strconv.ParseInt(result, 0, 64)
			if err1 != nil {
				output.Log.Errorf("%s in result of operator lesser", err1.Error())
			}

			iExpect, err2 := strconv.ParseInt(expect, 0, 64)
			if err2 != nil {
				output.Log.Errorf("%s in expected of operator lesser", err2.Error())
			}
			if err1 == nil && err2 == nil {
				return iResult < iExpect
			}
		case constants.OPERATOR_GREATER:

			iResult, err1 := strconv.ParseInt(result, 0, 64)
			if err1 != nil {
				output.Log.Errorf("%s in result of operator greater", err1.Error())
			}

			iExpect, err2 := strconv.ParseInt(expect, 0, 64)
			if err2 != nil {
				output.Log.Errorf("%s in expected of operator greater", err2.Error())
			}

			if err1 == nil && err2 == nil {
				return iResult > iExpect
			}
		default:
			output.Log.Errorln(fmt.Errorf("unknown operator: '%s'", operator))
			return false
		}
	}
	return false
}

// saveArtefact saves the stdout of the given command s as memory.Artefact inside the given memory.Task.
// The artefact name is generated with the prefix 'output' and from the commands arguments.
//
// returns false only if the runtime os is not supported, as building and using the commands depends on the runtime os
//
func saveArtefact(task *memory.Task, s exec.Cmd, stdout []byte) bool {
	// check arguments of cmd for possible files
	var paths []string
	switch runtime.GOOS {
	case "linux", "darwin":
		paths = detectArtefactFilesInArgs(s.Args[1:])
	case "windows":
		paths = detectArtefactFilesInArgs(s.Args)
	}

	for _, path := range paths {
		memory.AdditionalFiles[path] = true
	}
	// add command outputs to corresponding task (difference in artefact name generation depending on OS)
	output.Log.Debugf("Saving output of Command '%s' as Artefact", strings.Join(s.Args, " "))
	switch runtime.GOOS {
	case "windows":
		task.AddArtefact(memory.NewArtefact(strings.Join(append([]string{"output", filepath.Base(s.Path)},
			s.Args...),
			" "),
			string(stdout)))
	case "linux":
		task.AddArtefact(memory.NewArtefact(strings.Join(append([]string{"output"}, s.Args...), " "),
			string(stdout)))
	case "darwin":
		task.AddArtefact(memory.NewArtefact(strings.Join(append([]string{"output"}, s.Args...), " "),
			string(stdout)))
	case "default":
		return false
	}
	return true
}

// executeTaskTypeScript executes the given memory.Task of type constants.TASK_TYPE_SCRIPT,
// calling RunScript for the actual script execution. Sets the result of the given task if RunScript reports an error.
//
// returns an error if the given task is nil or RunScript returns an error.
//
func executeTaskTypeScript(task *memory.Task) error {
	if task == nil {
		return fmt.Errorf("no task (nil)")
	}

	err := RunScript(task)
	output.Log.Debugln("setting task status")
	if err != nil {
		task.SetStatus(constants.STATUS_ERROR)
		task.SetStatusMessage(err.Error())
		return err
	} else {
		if checkResult(task.GetResult(), task.Expected, task.Operator) {
			task.SetStatus(constants.STATUS_SUCCESS)
		} else {
			task.SetStatus(constants.STATUS_FAIL)
		}
	}

	return nil
}

// runCommandFromScript as opposed to executeTaskTypeCommand does not get its command string directly from the task,
// but from an additional parameter, which enables the call of the function and executing commands,
// which where not directly defined inside the task, but its artefacts can be saved to the given task either way.
// This is necessary as the script execution allows passing shell commands from the script to the go runtime,
// which shall be executed, but whose artefacts must be saved outside of the goja runtime.
//
// returns an empty string and an according error if buildCommandsFromString reports an error or the returned cmdList
// is empty.
//
// returns otherwise the cleaned output (stdout) and error of runCmd
//
func runCommandFromScript(command string, task *memory.Task) (string, error) {
	cmdList, err := buildCommandsFromString(command)
	if err != nil {
		return "", err
	}
	if cmdList == nil {
		return "", fmt.Errorf("empty command list")
	}

	var tmp string
	var endResult string

	for i, s := range cmdList {
		if i > 0 {
			s.Stdin = strings.NewReader(tmp) // set the old output as input for the new command
		}
		result, err := runCmd(s)
		result = cleanShellOutputs(result)
		tmp = result
		if err != nil {
			output.Log.Errorln(err)
			return "", err
		}
		if i == len(cmdList)-1 { // check if it's the last command
			output.Log.Debugf("Saving Result of last Command into return Result ('%s')...", strings.Join(s.Args, " "))
			endResult = result
		}

		// put every output of every command into artefact if task != nil
		if memory.OptionsList.Verbosity >= constants.VERBOSITY_DEBUG {
			saveArtefact(task, s, []byte(result))
		} else if i == 0 { // put output of first command into artefact
			saveArtefact(task, s, []byte(result))
		}
	}
	return endResult, nil
}

// buildCommandsFromString receives a command and calls the corresponded function for the specific runtime OS
//
// returns a slice of exec.Cmd and an error
//
//  buildCommandsFromString("cmd echo test")
//
func buildCommandsFromString(command string) ([]exec.Cmd, error) {
	switch runtime.GOOS {
	case "windows":
		return buildCommandsFromStringWindows(command)
	case "linux":
		return buildCommandsFromStringLinux(command)
	case "darwin":
		return buildCommandsFromStringMacOS(command)
	default:
		return nil, fmt.Errorf("os not supported: %s", runtime.GOOS)
	}
}

// buildCommandsFromStringWindows receives a command and builds a commandList for Windows
//
// the first word in the string has to be an executable either provided by absolut path or available in %PATH%!
//
// returns a slice of exec.Cmd and an error
//
//  buildCommandsFromStringWindows("cmd echo test")
//  buildCommandsFromStringWindows("type file.txt")
//
func buildCommandsFromStringWindows(command string) (commandList []exec.Cmd, err error) {
	if command == "" {
		return nil, fmt.Errorf("empty command")
	}

	// command = lsmod | grep whatever | grep whatever2
	var pipeList []string // pipeList[0] = "lsmod", pipeList[1] = "grep -a whatever",  pipeList[2] = "grep whatever2"
	var app string
	var powershell bool

	// var arg string
	var argList []string
	pipeList = memory.SeparateAtChar(command, constants.SHELL_PIPE, true)

	for i, s := range pipeList {

		// s = " grep whatever "
		s = strings.TrimSpace(s) // removes all leading and trailing white spaces

		stringsTmp := memory.SeparateAtChar(s, ' ', true) // split string by white spaces
		// stringsTmp = ["grep","whatever"]
		var cmd exec.Cmd
		output.Log.Tracef("Using following string for building the cmdList: %s", s)
		// first option build powershell command recursive without piping the input
		if i == 0 {

			if memory.CheckRegex(constants.REGEX_POWERSHELL, stringsTmp[0]) { // check powershell or ps,
				// uppercase/lowercase ignored (can handle " ' ")
				output.Log.Debugln("Powershell at the start of the string detected")
				powershell = true
			}

			if memory.CheckRegex(constants.REGEX_WINDOWS_OPTION_SLASH_C, s) { // check if string contains /C /c
				output.Log.Debugln("/C detected")
				app = stringsTmp[0]                          // save the app (powershell, cmds etc.)
				argList = append(argList, stringsTmp[1:]...) // save the args (/c /C)
				cmd = *exec.Command(app)
				cmd.Args = argList
			} else {
				output.Log.Debugln("/C not detected adding to string")
				app = stringsTmp[0]             // save the app (powershell, cmds etc.)
				argList = append(argList, "/C") // set the arg to /C so powershell and cmd executes the command directly
				cmd = *exec.Command(app)
				argList = append(argList, stringsTmp[1:]...)
				cmd.Args = argList
			}

		} else { // first run is done
			if powershell {
				output.Log.Debugln("Powershell detected, building corresponding commands")
				cmd = *exec.Command(app)                                // set app (as we need it all the time)
				argList = append(argList, string(constants.SHELL_PIPE)) // set pipe between commands
				argList = append(argList, stringsTmp...)                // append stringsTmp to argList slice
				cmd.Args = argList
			} else { // first run is done and no powershell
				output.Log.Debugln("No Powershell detected building corresponding cmd")
				cmd = *exec.Command(app)                           // set app (as we need it all the time)
				stringsTmp = prependString(stringsTmp, argList[0]) // set arg[0] at the first index of stringsTmp
				cmd.Args = stringsTmp
			}
		}
		commandList = append(commandList, cmd)
	}
	return
}

// prependString is the opposite of append() it adds a string at index 0 of the current string slice.
// If multiple strings are given the last string is at index 0
// prependString returns the updated slice. It is therefore necessary to store the
// result of prepend, often in the variable holding the slice itself.
//
// 	slice = prependString(slice, element1, element2)
// 	slice = prependString(slice, element1)
//
func prependString(source []string, add ...string) []string {

	for _, s := range add {
		source = append(source, "")
		copy(source[1:], source)
		source[0] = s
	}
	return source
}

// buildCommandsFromStringLinux receives a command and builds a commandList for Linux
//
// the first word in the string has to be an executable! (echo, cat ...)
//
// returns a slice of exec.Cmd and an error
//
//  buildCommandsFromStringLinux("echo test")
//
func buildCommandsFromStringLinux(command string) (commandList []exec.Cmd, err error) {

	if command == "" {
		return nil, fmt.Errorf("empty command")
	}

	// command = lsmod | grep whatever | grep whatever2
	var pipeList []string // pipeList[0] = "lsmod", pipeList[1] = "grep -a whatever",  pipeList[2] = "grep whatever2"
	output.Log.Debugln("Split command string by | ")

	pipeList = memory.SeparateAtChar(command, constants.SHELL_PIPE, true)

	for _, s := range pipeList {
		output.Log.Tracef("Using following string for building the cmdList: %s", s)
		output.Log.Debugln("Building commandList for Linux")
		// s = " grep whatever "
		s = strings.TrimSpace(s) // removes all leading and trailing white spaces
		// s ="grep whatever"
		stringsTmp := memory.SeparateAtChar(s, ' ', false) // split string by white spaces

		if memory.CheckRegex(constants.REGEX_UNIX_SHELL_BASH, stringsTmp[0]) {
			if !memory.CheckRegex(constants.REGEX_UNIX_OPTION_HYPHEN_C, stringsTmp[1]) {
				output.Log.Debugln("-C not detected")
				tmp := stringsTmp[0]
				stringsTmp[0] = "-c"
				stringsTmp = prependString(stringsTmp, tmp)
			}
		}
		// stringsTmp = ["grep","whatever"]
		cmd := *exec.Command(stringsTmp[0]) // stringsTmp[0] == executed app (lsmod, grep, cmd...)
		cmd.Args = stringsTmp

		commandList = append(commandList, cmd)
	}
	return
}

// buildCommandsFromStringMacOS receives a command and builds a commandList for MacOS
//
// the first word in the string has to be an executable! (echo, cat ...)
//
// returns a slice of exec.Cmd and an error
//
//  buildCommandsFromStringMacOS("echo test")
//
func buildCommandsFromStringMacOS(command string) (commandList []exec.Cmd, err error) {
	if command == "" {
		return nil, errors.New("command empty")
	}

	// command = lsmod | grep whatever | grep whatever2
	var pipeList []string // pipeList[0] = "lsmod", pipeList[1] = "grep -a whatever",  pipeList[2] = "grep whatever2"
	output.Log.Debugln("Split command string by | ")

	pipeList = memory.SeparateAtChar(command, constants.SHELL_PIPE, false) // separated by "|"

	for _, s := range pipeList {
		output.Log.Tracef("Using following string for building the cmdList: %s", s)
		output.Log.Debugln("Building commandList for MacOS")
		// s = " grep whatever "
		s = strings.TrimSpace(s) // removes all leading and trailing white spaces
		// s ="grep whatever"
		stringsTmp := memory.SeparateAtChar(s, ' ', false) // split string by white spaces

		if memory.CheckRegex(constants.REGEX_UNIX_SHELL_BASH, stringsTmp[0]) {
			if !memory.CheckRegex(constants.REGEX_UNIX_OPTION_HYPHEN_C, stringsTmp[1]) {
				output.Log.Debugln("-C not detected")
				tmp := stringsTmp[0]
				stringsTmp[0] = "-c"
				stringsTmp = prependString(stringsTmp, tmp)
			}
		}

		// stringsTmp = ["grep","whatever"]
		cmd := *exec.Command(stringsTmp[0]) // stringsTmp[0] == executed app (lsmod, grep, cmd...)
		cmd.Args = stringsTmp

		commandList = append(commandList, cmd)
	}
	return
}

// cleanShellOutputs removes first trailing new line (\n) in a string then carriage return (\r) in order to support removals for windows/linux/darwin
//
//  cleanShellOutputs("stringToClean\r\n")   result: "stringToClean"
//  cleanShellOutputs("stringToClean\n")     result: "stringToClean"
//  cleanShellOutputs("stringToClean\r")     result: "stringToClean"
//  cleanShellOutputs("stringToClean\n\r")   result: "stringToClean\n"
//
func cleanShellOutputs(toClean string) string {
	output.Log.Tracef("cleaning string: %s", toClean)
	toClean = strings.TrimSuffix(toClean, "\n")
	toClean = strings.TrimSuffix(toClean, "\r")
	return toClean
}

// detectArtefactFilesInArgs takes an exec.Cmd and iterates over its arguments,
// testing each for whether it is valid system file path.
//
// returns a slice with all as file paths approved strings
//
func detectArtefactFilesInArgs(args []string) []string {
	var paths []string
	for _, a := range args {
		a = memory.CleanString(a, `['"]`)
		// dont touch this magic!
		if checkValidFileCharacters(a) {
			info, err := os.Stat(a)
			if !os.IsNotExist(err) {
				if !info.IsDir() && info.Mode().IsRegular() {
					paths = append(paths, a)
				}
			}
			continue
		} else if filepath.IsAbs(a) {
			info, err := os.Stat(a)
			if !os.IsNotExist(err) {
				if !info.IsDir() && info.Mode().IsRegular() {
					paths = append(paths, a)
				}
			}
		}
	}
	return paths
}

// checkValidFileCharacters checks against invalid characters in a file name by supported OS
//
// windows: \\/'"?%*|<>
//
// linux: /
//
// macOS: /
//
// returns true if filename contains only valid file characters, false other ways or if OS is not supported
func checkValidFileCharacters(input string) bool {
	switch runtime.GOOS {
	case "windows":
		return !memory.CheckRegex(constants.REGEX_WINDOWS_INVALID_FILE_NAME_CHARACTERS, input)
	case "linux":
		return !memory.CheckRegex(constants.REGEX_UNIX_INVALID_FILE_NAME_CHARACTERS, input)
	case "darwin":
		return !memory.CheckRegex(constants.REGEX_UNIX_INVALID_FILE_NAME_CHARACTERS, input)
	default:
		return false
	}
}
