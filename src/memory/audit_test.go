package memory

import (
	"testing"
)

func TestAudit_GetStatus(t *testing.T) {
	tests := []struct {
		name  string
		audit Audit
		want  string
	}{
		{name: "Test1", audit: Audit_Builder("", nil, "This is my Status", ""), want: "This is my Status"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.audit.GetStatus(); got != tt.want {
				t.Errorf("Audit.GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAudit_GetStatusMessage(t *testing.T) {
	tests := []struct {
		name  string
		audit Audit
		want  string
	}{
		{name: "Test1", audit: Audit_Builder("", nil, "", "This is my StatusMessage"), want: "This is my StatusMessage"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.audit.GetStatusMessage(); got != tt.want {
				t.Errorf("Audit.GetStatusMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAudit_GetExecuted(t *testing.T) {
	tests := []struct {
		name  string
		audit Audit
		want  bool
	}{
		{name: "Test1", audit: Audit_Builder("Name", nil, "", ""), want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.audit.GetExecuted(); got != tt.want {
				t.Errorf("Audit.GetExecuted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuditDTO_GetFrom(t *testing.T) {
	type args struct {
		s Audit
	}
	tests := []struct {
		name string
		d    *AuditDTO
		args args
	}{
		{name: "Test1", d: AuditDTOForTest_Builder("name", nil, "status", "StatusMessage"), args: args{Audit_Builder("", nil, "", "")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.GetFrom(tt.args.s)
		})
	}
}

func TestAuditDTO_transfer(t *testing.T) {
	type args struct {
		s Audit
	}
	tests := []struct {
		name string
		d    *AuditDTO
		args args
	}{
		{name: "Test1", d: AuditDTOForTest_Builder("name", nil, "status", "StatusMessage"), args: args{Audit_Builder("", nil, "", "")}},
		{name: "Test2", d: AuditDTOForTest_Builder("name", []TaskDTO{TaskDTO_Builder("command", "echo hello", nil, "hello", "", "")}, "status", "StatusMessage"), args: args{Audit_Builder("", []Task{Task2_Builder("command", "echo Hello", []string{"Hello"}, "equal", nil, nil, nil, "", "", "This is my StatusMessage", "", false, "")}, "", "")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.transfer(tt.args.s)
		})
	}
}
