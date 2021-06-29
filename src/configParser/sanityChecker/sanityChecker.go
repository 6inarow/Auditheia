// Package sanityChecker provides functions to check Audits and Tasks for correctly set values
//
package sanityChecker

import (
	"os"
	"os/exec"
	"runtime"
	"strings"

	"Auditheia/memory"
	"Auditheia/memory/constants"
	"Auditheia/output"

	"github.com/sirupsen/logrus"
)

// CheckAuditList checks all audits in memory.AuditList for potentially incorrect values via calling CheckAudit for
// each of them.
//
// returns false if at least one audit or their tasks contain incorrect values
//
func CheckAuditList() (sane bool) {
	sane = true
	output.Log.Infoln("sanity check of all parsed audits requested:")
	for n := range memory.AuditList {
		output.Log.Infof("checking audit %d", n)
		if !CheckAudit(&memory.AuditList[n]) {
			sane = false
		}
	}
	output.Log.Infoln("finished sanity check of all parsed audits")
	return
}

var sanityFieldsAudit = logrus.Fields{
	"sanity_check": "audit",
}

// CheckAudit checks the given audit for none empty fields. If the audit contains tasks,
// each task is checked via CheckTask.
//
// returns true if the audit itself and none of its tasks contain invalid fields
//
func CheckAudit(audit *memory.Audit) (sane bool) {
	sane = true
	if audit.Name == "" {
		output.Log.WithFields(sanityFieldsAudit).Warningf("empty audit name")
		sane = false
	}
	if len(audit.Tasks) == 0 {
		output.Log.WithFields(sanityFieldsAudit).Warningf("empty task list")
		sane = false
	} else {
		for n := range audit.Tasks {
			output.Log.Infof("checking task %d", n)
			if !CheckTask(&audit.Tasks[n]) {
				sane = false
			}
		}
	}

	return
}

var sanityFieldsTask = logrus.Fields{
	"sanity_check": "task",
}

// CheckTask checks the task for none empty fields and whether the contents of Task.Execute are possibly invalid (
// currently only checks if the defined executables exist). If the task contains other tasks, each task is checked.
//
// returns true if no invalid values have been found
//
func CheckTask(task *memory.Task) (sane bool) {
	sane = true
	if !validType(task.Type) {
		output.Log.WithFields(sanityFieldsTask).Warningf("invalid task type: '%s'", task.Type)
		sane = false
	}
	if task.Execute == "" {
		output.Log.WithFields(sanityFieldsTask).Warningf("no executable command or script defined")
		sane = false
	} else if task.Type == constants.TASK_TYPE_COMMAND && !validCommand(task.Execute) {
		output.Log.WithFields(sanityFieldsTask).Warningf("defined command invalid: '%s'", task.Execute)
		sane = false
	} /*else if task.Type == constants.TASK_TYPE_SCRIPT && !validScript(task.Execute) {
		output.Log.WithFields(sanityFieldsTask).Warningf("defined script invalid")
		sane = false
	}*/

	if len(task.Expected) == 0 {
		output.Log.WithFields(sanityFieldsTask).Warningf("no expected values defined")
		sane = false
	}
	if !validOperator(task.Operator) {
		output.Log.WithFields(sanityFieldsTask).Warningf("invalid operator: '%s'", task.Operator)
		sane = false
	}

	if task.OnSuccess != nil && !CheckTask(task.OnSuccess) {
		sane = false
	}
	if task.OnFail != nil && !CheckTask(task.OnFail) {
		sane = false
	}
	if task.OnError != nil && !CheckTask(task.OnError) {
		sane = false
	}
	return
}

// CheckAdditionalFiles checks whether the defined additional files are valid paths (and exist)
//
// returns true if the path exists on the system
//
func CheckAdditionalFiles(files map[string]bool) bool {
	sane := true
	for path := range files {
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			output.Log.Warningf("path does not exist: '%s'", path)
			sane = false
		}
	}
	return sane
}

// validOperator checks whether the passed operator is known to the program (i.e.
// equals one of the defined operators in the Package constants).
//
// returns true only if the operator is known
//
func validOperator(operator string) bool {
	switch operator {
	case constants.OPERATOR_EQUALS, constants.OPERATOR_NOT_EQUALS, constants.OPERATOR_CONTAINS, constants.OPERATOR_NOT_CONTAINS, constants.OPERATOR_GREATER, constants.OPERATOR_LESSER:
		return true
	default:
		return false
	}
}

// validType checks whether the passed task type is known to the program (i.e.
// equals one of the defined types in the Package constants)
//
// returns true if the type is known
//
func validType(taskType string) bool {
	switch taskType {
	case constants.TASK_TYPE_COMMAND, constants.TASK_TYPE_SCRIPT:
		return true
	default:
		return false
	}
}

// validCommand checks whether the command's executable is found on the system.
//
// returns true if the executable is found, false if not
//
func validCommand(command string) bool {
	switch runtime.GOOS {
	case "darwin":
		return validDarwinCommand(command)
	case "linux":
		return validLinuxCommand(command)
	case "windows":
		return validWindowsCommand(command)
	default:

	}
	return false
}

// validDarwinCommand checks whether the command for MacOS can be executed by executing 'open -Ra' with the executables
// name as an argument. Example for the command cat:
//
//  /bin/sh -c open -Ra cat
//
// returns true if execution of the above command does not return an error
//
func validDarwinCommand(command string) (sane bool) {
	sane = true
	cmds := memory.SeparateAtChar(command, '|', true)
	for _, cmd := range cmds {
		if err := exec.Command("/bin/sh", "-c", "open", "-Ra", strings.Fields(cmd)[0]).Run(); err != nil {
			output.Log.WithFields(sanityFieldsTask).Errorf("command not found: '%s'", strings.Fields(cmd)[0])
			sane = false
		}
	}
	return
}

// validLinuxCommand checks whether the command can be executed on linux by executing 'command -v' with the
// executables name as an argument. Example for the command cat:
//
//  /bin/sh -c command -v cat
//
// returns true if execution of the above command does not return an error
//
func validLinuxCommand(command string) (sane bool) {
	sane = true
	cmds := memory.SeparateAtChar(command, '|', true)
	for _, cmd := range cmds {
		if err := exec.Command("/bin/sh", "-c", "command -v "+strings.Fields(cmd)[0]).Run(); err != nil {
			output.Log.WithFields(sanityFieldsTask).Errorf("command not found: '%s'", strings.Fields(cmd)[0])
			sane = false
		}
	}
	return
}

// validWindowsCommand checks whether the command for Windows is an executable in %PATH% based on the app of a string
// ( e.g. powershell, cmd, findstr... )
//
// The function uses the where.exe executable available since Windows Vista in a cmd execution
//
func validWindowsCommand(command string) (sane bool) {
	cmds := memory.SeparateAtChar(command, '|', true)
	for _, cmd := range cmds {
		if memory.CheckRegex(strings.Fields(cmd)[0], `(?i)(powershell\b)`) {
			sane = true
			return
		} else if memory.CheckRegex(strings.Fields(cmd)[0], `(?i)(cmd\b)`) {
			sane = true
		} else if err := exec.Command("cmd", "/c", "where ", strings.Fields(cmd)[0]).Run(); err != nil {
			output.Log.WithFields(sanityFieldsTask).Errorf("command not found: '%s'", strings.Fields(cmd)[0])
			sane = false
		}
	}
	return
}

// validScript checks the given string for valid ECMAScript 5.1 Syntax.
//
// returns true if no syntax violations have been found
//
// Deprecated: will most likely not be implemented
//
func validScript(script string) bool {
	// TODO
	return false
}
