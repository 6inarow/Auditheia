package memory

// Audit contains the name of the audit in the config and the corresponding Tasks.
type Audit struct {
	Name          string `json:"name" yaml:"name"`
	Tasks         []Task `json:"tasks" yaml:"tasks"`
	status        string
	statusMessage string
	executed      bool
}

// GetStatus provided by Audit object returns the current Status of the audit
//
func (audit Audit) GetStatus() string {
	return audit.status
}

// SetStatus provided by Audit object sets the Status of the audit
//
func (audit *Audit) SetStatus(status string) {
	audit.status = status
}

// GetStatusMessage provided by Audit object returns the current status message of the audit
//
func (audit Audit) GetStatusMessage() string {
	return audit.statusMessage
}

// SetStatusMessage provided by Audit object sets the current status message of the audit
//
func (audit *Audit) SetStatusMessage(message string) {
	audit.statusMessage = message
}

// GetExecuted provided by Audit object returns true if the audit was executed
//
func (audit Audit) GetExecuted() bool {
	return audit.executed
}

// SetExecuted provided by Audit object sets the Audit execution state executed
//
// EXECUTION STATE CANNOT BE REVERTED!
//
func (audit *Audit) SetExecuted() {
	audit.executed = true
}

// AuditDTO serves as data transfer object, containing the original Audit's Name, Tasks (as TaskDTO),
// Status and StatusMessage.
//
type AuditDTO struct {
	Name          string    `json:"name" yaml:"name"`
	Tasks         []TaskDTO `json:"tasks" yaml:"tasks"`
	Status        string    `json:"status,omitempty" yaml:"status,omitempty"`
	StatusMessage string    `json:"status_message,omitempty" yaml:"status_message,omitempty"`
}

// GetFrom copies fields relevant for the report from the passed Audit to the calling AuditDTO.
// For each Task in the Audit a TaskDTO is created.
//
func (d *AuditDTO) GetFrom(s Audit) {
	d.transfer(s)
}

// transfer copies fields relevant for the report from the passed Audit to the calling AuditDTO.
// For each Task in the Audit the TaskDTO.transfer function is called as well.
//
func (d *AuditDTO) transfer(s Audit) {
	// transfer name
	d.Name = s.Name
	// transfer status
	d.Status = s.status
	// transfer status message
	d.StatusMessage = s.statusMessage
	// transfer tasks
	for _, sTask := range s.Tasks {
		var dTask = TaskDTO{}
		dTask.transfer(sTask)
		d.Tasks = append(d.Tasks, dTask)
	}
}
