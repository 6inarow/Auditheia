package memory

import (
	"io"
	"os"
	"os/exec"
	"syscall"
)

func Artefact_Builder(name string, content string, filePath string) Artefact {

	artefact := &Artefact{
		Name:    name,
		Content: content,
	}
	return *artefact
}

func CmdForTest_Builder(path string, args []string, env []string, dir string, stdout io.Writer, stderr io.Writer, extraFiles []*os.File, sysProcAttr *syscall.SysProcAttr, process *os.Process, processState *os.ProcessState) exec.Cmd {
	cmd := exec.Cmd{
		Path:         path,
		Args:         args,
		Env:          env,
		Dir:          dir,
		Stdout:       stdout,
		Stderr:       stderr,
		ExtraFiles:   extraFiles,
		SysProcAttr:  sysProcAttr,
		Process:      process,
		ProcessState: processState}

	return cmd
}

func TaskDTOForTest_Builder(Type string, executed string, following *TaskDTO, result string, status string, statusMessage string) *TaskDTO {
	taskDTO := &TaskDTO{
		Type:          Type,
		Executed:      executed,
		Following:     following,
		Result:        result,
		Status:        status,
		StatusMessage: statusMessage,
	}
	return taskDTO
}

func TaskDTO_Builder(Type string, executed string, following *TaskDTO, result string, status string, statusMessage string) TaskDTO {

	taskDTO := &TaskDTO{
		Type:          Type,
		Executed:      executed,
		Following:     following,
		Result:        result,
		Status:        status,
		StatusMessage: statusMessage,
	}
	return *taskDTO
}

func AuditDTOForTest_Builder(name string, tasks []TaskDTO, status string, statusMessage string) *AuditDTO {
	audit := &AuditDTO{
		Name:          name,
		Tasks:         tasks,
		Status:        status,
		StatusMessage: statusMessage}

	return audit
}

func AuditForTest_Builder(name string, tasks []Task, status string, statusMessage string) *Audit {
	audit := &Audit{Name: name, Tasks: tasks}
	audit.SetStatus(status)
	audit.SetStatusMessage(statusMessage)
	audit.SetExecuted()
	return audit
}

func Audit_Builder(name string, tasks []Task, status string, statusMessage string) Audit {

	audit := &Audit{Name: name, Tasks: tasks}
	audit.SetStatus(status)
	audit.SetStatusMessage(statusMessage)
	audit.SetExecuted()
	return *audit
}

func Task_Builder(Type string, execute string, expected []string, operator string, onSuccess *Task,
	onFail *Task, onError *Task, setResult string, setStatus string, setStatusMessage string,
	artifactName string, artifactIsFile bool, artifactContent string) *Task {

	var task *Task = &Task{
		Type:          Type,
		Execute:       execute,
		Expected:      expected,
		Operator:      operator,
		OnSuccess:     onSuccess,
		OnFail:        onFail,
		OnError:       onError,
		result:        setResult,
		status:        setStatus,
		statusMessage: setStatusMessage,
		executed:      true,
	}

	task.AddArtefact(&Artefact{
		Name:    artifactName,
		Content: artifactContent,
	})

	return task
}

func TaskForTest_Builder(Type string, execute string, expected []string, operator string, onSuccess *Task,
	onFail *Task, onError *Task, setResult string, setStatus string, setStatusMessage string,
	artifactName string, artifactIsFile bool, artifactContent string) *Task {

	task := &Task{
		Type:      Type,
		Execute:   execute,
		Expected:  expected,
		Operator:  operator,
		OnSuccess: onSuccess,
		OnFail:    onFail,
		OnError:   onError}
	task.SetResult(setResult)
	task.SetStatus(setStatus)
	task.SetStatusMessage(setStatusMessage)
	task.SetExecuted()
	task.AddArtefact(&Artefact{
		Name:    artifactName,
		Content: artifactContent,
	})
	return task
}

// just built for memory_test count task
func Task2_Builder(Type string, execute string, expected []string, operator string, onSuccess *Task,
	onFail *Task, onError *Task, setResult string, setStatus string, setStatusMessage string,
	artifactName string, artifactIsFile bool, artifactContent string) Task {

	task := &Task{
		Type:      Type,
		Execute:   execute,
		Expected:  expected,
		Operator:  operator,
		OnSuccess: onSuccess,
		OnFail:    onFail,
		OnError:   onError}
	task.SetResult(setResult)
	task.SetStatus(setStatus)
	task.SetStatusMessage(setStatusMessage)
	task.SetExecuted()
	task.AddArtefact(&Artefact{
		Name:    artifactName,
		Content: artifactContent,
	})
	return *task
}
