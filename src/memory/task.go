package memory

// Task struct contains all necessary fields of an audit task, noteworthy the task type (script or command),
// what to execute (e.g. script or command), what the expected results are,
// which operator shall be used to compare actual result and expected results and which follow up tasks shall be
// executed depending on the current task's status after execution.
//
// All results are saved as well, but private to prevent unmarshalling functions from accidentally altering result,
// status, execution status, etc.
//
type Task struct {
	Type          string   `json:"type" yaml:"type"`
	Execute       string   `json:"execute" yaml:"execute"`
	Expected      []string `json:"expected" yaml:"expected"`
	Operator      string   `json:"operator" yaml:"operator"`
	OnSuccess     *Task    `json:"on_success" yaml:"on_success"`
	OnFail        *Task    `json:"on_fail" yaml:"on_fail"`
	OnError       *Task    `json:"on_error" yaml:"on_error"`
	result        string
	status        string
	statusMessage string
	executed      bool
	artefacts     []Artefact
}

// GetResult returns the task's result.
//
func (task Task) GetResult() string {
	return task.result
}

// SetResult enables setting of the task's result.
func (task *Task) SetResult(result string) {
	task.result = result
}

// GetStatus returns the task's status.
//
func (task Task) GetStatus() string {
	return task.status
}

// SetStatus enables setting of the task's status.
//
func (task *Task) SetStatus(status string) {
	task.status = status
}

// GetStatusMessage returns the task's status message.
//
func (task Task) GetStatusMessage() string {
	return task.statusMessage
}

// SetStatusMessage enables setting of the task's status message.
//
func (task *Task) SetStatusMessage(message string) {
	task.statusMessage = message
}

// GetExecuted returns true, if the task has been executed.
//
func (task Task) GetExecuted() bool {
	return task.executed
}

// SetExecuted always sets the task's execution state true.
//
// EXECUTION STATE CANNOT BE REVERTED!
//
func (task *Task) SetExecuted() {
	task.executed = true
}

// GetArtefacts returns a slice with all artefacts contained by the task.
//
func (task Task) GetArtefacts() []Artefact {
	return task.artefacts
}

// AddArtefact adds the passed Artefact to the task's artefact slice.
//
func (task *Task) AddArtefact(artefact *Artefact) {
	task.artefacts = append(task.artefacts, *artefact)
}

// TaskDTO serves as a data transfer object containing the original task's type, executed string,
// which subsequent task was executed (if one exists), the original task's result, status and status message.
//
// The fields Following and StatusMessage will be omitted by marshallers if left empty.
type TaskDTO struct {
	Type     string `json:"type" yaml:"type"`
	Executed string `json:"executed" yaml:"executed"`

	Result        string   `json:"result" yaml:"result"`
	Expected      []string `json:"expected" yaml:"expected"`
	Operator      string   `json:"operator" yaml:"operator"`
	Status        string   `json:"status" yaml:"status"`
	StatusMessage string   `json:"status_message,omitempty" yaml:"status_message,omitempty"`
	Following     *TaskDTO `json:"following,omitempty" yaml:"following,omitempty"`
}

// transfer copies fields relevant for the report from the passed Task to the calling TaskDTO. As Task contains three
// fields in which another Task may be stored and the TaskDTO does not, the function looks into each of these Fields of
// Task and checks whether the following Task exists and has been executed. Only if the Field is not nil and the Task
// has been executed, it will be copied to the TaskDTO. The three fields are Checked in a specific order: OnSuccess,
// OnFail and OnError.
func (d *TaskDTO) transfer(s Task) {
	// transfer type
	d.Type = s.Type
	// transfer executed
	d.Executed = s.Execute
	// transfer result
	d.Result = s.result
	//transfer expected slice
	d.Expected = s.Expected
	//transfer operator
	d.Operator = s.Operator
	// transfer status
	d.Status = s.status
	// transfer status message
	d.StatusMessage = s.statusMessage
	// transfer next executed task
	var dFollowing = TaskDTO{}
	if s.OnSuccess != nil && s.OnSuccess.executed {
		dFollowing.transfer(*s.OnSuccess)
		d.Following = &dFollowing
	} else if s.OnFail != nil && s.OnFail.executed {
		dFollowing.transfer(*s.OnFail)
		d.Following = &dFollowing
	} else if s.OnError != nil && s.OnError.executed {
		dFollowing.transfer(*s.OnError)
		d.Following = &dFollowing
	}
}
