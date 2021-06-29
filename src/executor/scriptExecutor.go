package executor

import (
	"fmt"
	"runtime"

	"Auditheia/memory"
	"Auditheia/output"

	"github.com/sirupsen/logrus"

	"github.com/dop251/goja"
)

// RunScript executes the scripts defined in the passed memory.Task, initializing the goja.Runtime,
// running the script inside and retrieving the result from the goja.
// Runtime and saving it again in the passed memory.Task.
//
// In case the goja.Runtime panics, recover() is called, logging the event to the logfile.
//
// returns an error if the setup of the goja.Runtime or running the script fails.
//
func RunScript(task *memory.Task) error {
	// defer recovering from panic in case the goja runtime panics
	defer func() {
		if r := recover(); r != nil {
			output.Log.Errorln("Panic during script execution:", r)
		}
	}()

	// initialize runtime
	output.Log.Debugln("initializing instance of goja runtime...")
	vm, err := initRuntime(task)
	if err != nil {
		return err
	}
	// run the script, discard returned value as a function for result retrieval will be implemented in initRuntime
	output.Log.Debugln("running script inside the goja runtime...")
	_, err = vm.RunString(task.Execute)
	if err != nil {
		return err
	}

	output.Log.Debugln("retrieving 'result' from goja runtime")
	result := vm.Get("result")
	if result == nil {
		return fmt.Errorf("result variable not defined in script")
	}

	output.Log.Debugln("setting result of script as task result")
	task.SetResult(result.String())

	return nil
}

// initRuntime creates the goja.Runtime and sets the variables and functions which then can be used inside
// the script.
//
// returns a pointer to a goja.Runtime and an error, if the goja.Runtime creation or setting variables/functions fail.
//
func initRuntime(task *memory.Task) (*goja.Runtime, error) {
	// create goja runtime
	output.Log.Debugln("creating new instance of goja runtime")
	vm := goja.New()
	// create console object
	c, err := newConsole(vm)
	if err != nil {
		return nil, err
	}
	// set console diversion (essentially adds the console object to the goja runtime)
	output.Log.Debugln("adding console diversion to goja runtime")
	if err := vm.Set("console", c); err != nil {
		return nil, err
	}
	// set the AuditheiaEngine "class", its functions will be available to JS as objects once a var is initiated with new AuditheiaEngine()
	// create auditheia object
	a, err := newAuditheia(vm, task)
	if err != nil {
		return nil, err
	}
	// add auditheia object to the goja runtime
	output.Log.Debugln("adding 'auditheia' object to goja runtime")
	if err := vm.Set("auditheia", a); err != nil {
		return nil, err
	}
	// return the goja runtime object
	return vm, nil
}

// auditheia is a struct which later serves as object inside the goja.Runtime.
// It additionally has fields containing the corresponding memory.Task, which the script originates from,
// to save artifacts at the correct place, and the current goja.Runtime.
//
type auditheia struct {
	correspondingTask *memory.Task
	vm                *goja.Runtime
}

// newAuditheia creates an auditheia object inside the passed goja.Runtime and sets functions/variables for the object.
//
// returns the goja.Object and an error, if setting of functions/variables fails.
//
func newAuditheia(vm *goja.Runtime, task *memory.Task) (*goja.Object, error) {
	a := &auditheia{correspondingTask: task, vm: vm}
	obj := vm.NewObject()
	if err := obj.Set("osExec", a.osExec); err != nil {
		return nil, err
	}
	if err := obj.Set("log", a.log); err != nil {
		return nil, err
	}
	if err := obj.Set("runtimeOS", runtime.GOOS); err != nil {
		return nil, err
	}
	return obj, nil
}

// osExec is a wrapper for the execution of system commands inside scripts via os/exec.
//
// returns an array containing an error if occurred as the first value and the command result as the second value.
//
func (a *auditheia) osExec(command string) []interface{} {
	output.Log.Debugf("script request execution of command: %s", command)
	result, err := runCommandFromScript(command, a.correspondingTask)
	if err != nil {
		output.Log.Tracef("error occured while trying to run command from script: %v", err)
		return []interface{}{err, result}
	} else {
		return []interface{}{nil, result}
	}
}

// log is a wrapper for logging of messages to the main programs logfiles.
// Adds an additional field stating the messages origin being the script engine.
// The passed integer 'level' determines the level in the logfiles and whether it will appear in the logfiles.
// See logging package for more details on log levels.
//
//  log(2, "this message has info level")
//
func (a *auditheia) log(level int, msg string) {
	// set the additional fields
	additionalFields := logrus.Fields{
		"origin": "script_engine",
	}
	// switch depending on passed verbosity level
	switch level {
	case 0:
		output.Log.WithFields(additionalFields).Errorln(msg)
	case 1:
		output.Log.WithFields(additionalFields).Warningln(msg)
	case 2:
		output.Log.WithFields(additionalFields).Infoln(msg)
	case 3:
		output.Log.WithFields(additionalFields).Debugln(msg)
	case 4:
		output.Log.WithFields(additionalFields).Traceln(msg)
	default:
		output.Log.WithFields(additionalFields).Warningln(msg)
	}
}

// console is a struct which later serves together with log as redirection for messages logged via "console.log"
// inside the goja.Runtime respectively the script.
//
type console struct{}

// newConsole creates an object which is later used to divert console outputs to the go console.
//
// returns a goja.Object and an error if the object setup fails
//
func newConsole(vm *goja.Runtime) (*goja.Object, error) {
	c := &console{}
	obj := vm.NewObject()
	if err := obj.Set("log", c.log); err != nil {
		return nil, err
	}
	return obj, nil
}

// log passes the given message to fmt (which prints it to stdout), adding a statement indicating the message origin
//
func (c *console) log(msg string) {
	fmt.Printf("Script Engine: %s\n", msg)
}
