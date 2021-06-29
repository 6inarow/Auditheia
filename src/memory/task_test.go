package memory

import (
	"reflect"
	"testing"
)

func TestTask_GetResult(t *testing.T) {
	tests := []struct {
		name string
		task Task
		want string
	}{
		{name: "Test1", task: Task2_Builder("command", "echo Hello", []string{"Hello"}, "equal", nil, nil, nil, "This is my Result", "", "", "", true, ""), want: "This is my Result"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.task.GetResult(); got != tt.want {
				t.Errorf("Task.GetResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_GetStatus(t *testing.T) {
	tests := []struct {
		name string
		task Task
		want string
	}{
		{name: "Test1", task: Task2_Builder("command", "echo Hello", []string{"Hello"}, "equal", nil, nil, nil, "", "This is my Status", "", "", true, ""), want: "This is my Status"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.task.GetStatus(); got != tt.want {
				t.Errorf("Task.GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_GetStatusMessage(t *testing.T) {
	tests := []struct {
		name string
		task Task
		want string
	}{
		{name: "Test1", task: Task2_Builder("command", "echo Hello", []string{"Hello"}, "equal", nil, nil, nil, "", "", "This is my StatusMessage", "", true, ""), want: "This is my StatusMessage"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.task.GetStatusMessage(); got != tt.want {
				t.Errorf("Task.GetStatusMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_GetArtefacts(t *testing.T) {
	tests := []struct {
		name string
		task Task
		want []Artefact
	}{
		{name: "Test1", task: Task2_Builder("command", "echo Hello", []string{"Hello"}, "equal", nil, nil, nil, "", "", "", "This is my ArtefactName", true, "This is my ArtefactContent"), want: []Artefact{Artefact_Builder("This is my ArtefactName", "This is my ArtefactContent", "path")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.task.GetArtefacts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task.GetArtefacts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_GetExecuted(t *testing.T) {
	tests := []struct {
		name string
		task Task
		want bool
	}{
		{name: "Test1", task: Task2_Builder("command", "echo Hello", []string{"Hello"}, "equal", nil, nil, nil, "", "", "This is my StatusMessage", "", false, ""), want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.task.GetExecuted(); got != tt.want {
				t.Errorf("Task.GetExecuted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskDTO_transfer(t *testing.T) {
	type args struct {
		s Task
	}
	tests := []struct {
		name string
		d    *TaskDTO
		args args
	}{
		{name: "Test1", d: TaskDTOForTest_Builder("command", "echo Hello", nil, "Hello", "", ""), args: args{Task2_Builder("command", "echo Hello", []string{"Hello"}, "equal", nil, nil, nil, "", "", "This is my StatusMessage", "", false, "")}},
		{name: "Test2", d: TaskDTOForTest_Builder("command", "echo Hello", nil, "Hello", "", ""), args: args{Task2_Builder("command", "echo Hello", []string{"Hello"}, "equal", TaskForTest_Builder("command", "echo Hello", []string{"Hello"}, "equal", nil, nil, nil, "", "", "This is my StatusMessage", "", false, ""), nil, nil, "", "", "This is my StatusMessage", "", false, "")}},
		{name: "Test3", d: TaskDTOForTest_Builder("command", "echo Hello", nil, "Hello", "", ""), args: args{Task2_Builder("command", "echo Hello", []string{"Hello"}, "equal", nil, TaskForTest_Builder("command", "echo Hello", []string{"Hello"}, "equal", nil, nil, nil, "", "", "This is my StatusMessage", "", false, ""), nil, "", "", "This is my StatusMessage", "", false, "")}},
		{name: "Test4", d: TaskDTOForTest_Builder("command", "echo Hello", nil, "Hello", "", ""), args: args{Task2_Builder("command", "echo Hello", []string{"Hello"}, "equal", nil, nil, TaskForTest_Builder("command", "echo Hello", []string{"Hello"}, "equal", nil, nil, nil, "", "", "This is my StatusMessage", "", false, ""), "", "", "This is my StatusMessage", "", false, "")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.transfer(tt.args.s)
		})
	}
}
