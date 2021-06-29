package sanityChecker

import (
	"Auditheia/memory"
	"Auditheia/memory/constants"
	"runtime"
	"testing"
)

func TestCheckAdditionalFiles_local1(t *testing.T) {
	type args struct {
		files map[string]bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test1", args: args{files: map[string]bool{"D:\\WRKPLC\\Auditheia\\Anwendung\\test.md": true, "file.md": true}}, want: false},
	}
	if runtime.GOOS == "windows" && constants.LOCAL_1 {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := CheckAdditionalFiles(tt.args.files); got != tt.want {
					t.Errorf("CheckAdditionalFiles() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func TestCheckAdditionalFiles_local2(t *testing.T) {
	type args struct {
		files map[string]bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test1", args: args{files: map[string]bool{"C:\\Users\\spark\\go\\src\\Auditheia\\Anwendung\\test.md": true, "file.md": true}}, want: false},
	}
	if runtime.GOOS == "windows" && constants.LOCAL_2 {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := CheckAdditionalFiles(tt.args.files); got != tt.want {
					t.Errorf("CheckAdditionalFiles() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func TestCheckAdditionalFiles_windows(t *testing.T) {
	type args struct {
		files map[string]bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test1", args: args{files: map[string]bool{"D:\\a\\Auditheia\\Auditheia\\Anwendung\\test.md": true, "file.md": true}}, want: true},
	}
	if runtime.GOOS == "windows" && constants.RUNNER {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := CheckAdditionalFiles(tt.args.files); got != tt.want {
					t.Errorf("CheckAdditionalFiles() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func TestCheckAdditionalFiles_linux(t *testing.T) {
	type args struct {
		files map[string]bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test1", args: args{files: map[string]bool{"home/runner/work/Auditheia/Auditheia/Anwendung/test.md": true, "file.md": true}}, want: true},
	}

	if runtime.GOOS == "linux" && constants.RUNNER {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := CheckAdditionalFiles(tt.args.files); got != tt.want {
					t.Errorf("CheckAdditionalFiles() = %v, want %v", got, tt.want)
				}
			})
		}
	}

}

func TestCheckAdditionalFiles_macos(t *testing.T) {
	type args struct {
		files map[string]bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// edit Path
		{name: "Test1", args: args{files: map[string]bool{"home/runner/work/Auditheia/Auditheia/Anwendung/test.md": true, "file.md": true}}, want: true},
	}

	if runtime.GOOS == "darwin" && constants.RUNNER {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := CheckAdditionalFiles(tt.args.files); got != tt.want {
					t.Errorf("CheckAdditionalFiles() = %v, want %v", got, tt.want)
				}
			})
		}
	}

}

func TestCheckAudit(t *testing.T) {
	type args struct {
		audit *memory.Audit
	}
	tests := []struct {
		name     string
		args     args
		wantSane bool
	}{
		{name: "Test1", args: args{memory.AuditForTest_Builder("Audit1 ", nil, "Status", "Message")}, wantSane: false},

		{name: "Test2", args: args{memory.AuditForTest_Builder("Audit2", []memory.Task{*memory.Task_Builder("command", "cmd /c echo hello",
			[]string{"result"}, "equals", nil, nil, nil, "result", "status", "setStatus", "ArtifactName", true, "content")}, "Status", "Message")}, wantSane: true},

		{name: "Test3", args: args{memory.AuditForTest_Builder("", []memory.Task{*memory.Task_Builder("command", "echo hello",
			[]string{"result"}, "equals", nil, nil, nil, "result", "status", "setStatus", "ArtifactName", true, "content")}, "Status", "Message")}, wantSane: false},
	}
	for n, _ := range tests {
		t.Run(tests[n].name, func(t *testing.T) {
			if gotSane := CheckAudit(tests[n].args.audit); gotSane != tests[n].wantSane {
				t.Errorf("CheckAudit() = %v, want %v", gotSane, tests[n].wantSane)
			}
		})
	}
}

func TestCheckAuditList(t *testing.T) {
	tests := []struct {
		name     string
		wantSane bool
	}{
		{"Test1", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSane := CheckAuditList(); gotSane != tt.wantSane {
				t.Errorf("CheckAuditList() = %v, want %v", gotSane, tt.wantSane)
			}
		})
	}
}

func TestCheckTask(t *testing.T) {
	type args struct {
		task *memory.Task
	}
	tests := []struct {
		name     string
		args     args
		wantSane bool
	}{
		{
			name:     "Test1",
			args:     args{memory.TaskForTest_Builder("", "", []string{"SameSame"}, "not equals", nil, nil, nil, "SameSame", "", "", "", true, "")},
			wantSane: false,
		},
		{
			name:     "Test2",
			args:     args{memory.TaskForTest_Builder("", "", []string{"SameSame"}, "notOperator", nil, nil, nil, "SameSame", "", "", "", true, "")},
			wantSane: false,
		},
		{
			name:     "Test4",
			args:     args{memory.TaskForTest_Builder("", "", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, "")},
			wantSane: false,
		},
		{ //task.onSuccess
			name:     "Test5",
			args:     args{memory.TaskForTest_Builder("command", "", []string{"SameSame"}, "equals", memory.TaskForTest_Builder("Command", "", []string{"Same"}, "equals", nil, nil, nil, "NotSameSame", "", "", "", true, ""), nil, nil, "SameSame", "", "", "", true, "")},
			wantSane: false,
		},
		{ //task.onFail
			name:     "Test6",
			args:     args{memory.TaskForTest_Builder("command", "", []string{"SameSame"}, "equals", nil, memory.TaskForTest_Builder("Command", "", []string{"Same"}, "equals", nil, nil, nil, "NotSameSame", "", "", "", true, ""), nil, "SameSame", "", "", "", true, "")},
			wantSane: false,
		},
		{ //task.onError expected auf nil
			name:     "Test7",
			args:     args{memory.TaskForTest_Builder("command", "", []string{"SameSame"}, "equals", nil, nil, memory.TaskForTest_Builder("Command", "", nil, "equals", nil, nil, nil, "", "", "", "", true, ""), "SameSame", "", "", "", true, "")},
			wantSane: false,
		},
		{ //Test for == 0
			name:     "Test8",
			args:     args{memory.TaskForTest_Builder("", "", nil, "equals", nil, nil, nil, "SameSame", "", "", "", true, "")},
			wantSane: false,
		},
		//Test for task.* sane=false
		{
			name:     "Test9",
			args:     args{memory.TaskForTest_Builder("", "", []string{"SameSame"}, "equals", memory.TaskForTest_Builder("", "", []string{"SameSame"}, "equals", memory.TaskForTest_Builder("", "", []string{"SameSame"}, "not equals", nil, nil, nil, "SameSame", "", "", "", true, ""), nil, nil, "SameSame", "", "", "", true, ""), nil, nil, "SameSame", "", "", "", true, "")},
			wantSane: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSane := CheckTask(tt.args.task); gotSane != tt.wantSane {
				t.Errorf("CheckTask() = %v, want %v", gotSane, tt.wantSane)
			}
		})
	}
}

func Test_validScript(t *testing.T) {
	type args struct {
		script string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test1", args: args{""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validScript(tt.args.script); got != tt.want {
				t.Errorf("validScript() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validWindowsCommand(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name     string
		args     args
		wantSane bool
	}{
		{name: "Test1", args: args{command: "powershell /c echo hello | echo"}, wantSane: true},
		{name: "Test2", args: args{command: "cmd /c echo hello | echo"}, wantSane: true},
		{name: "Test2", args: args{command: "cmd /c NoCommand hello | NoCommand"}, wantSane: false},
	}

	if runtime.GOOS == "windows" {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if gotSane := validWindowsCommand(tt.args.command); gotSane != tt.wantSane {
					t.Errorf("validWindowsCommand() = %v, want %v", gotSane, tt.wantSane)
				}
			})
		}
	}

}

func Test_validLinuxCommand(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name     string
		args     args
		wantSane bool
	}{
		{name: "Test1", args: args{command: "echo hello | echo"}, wantSane: true},
		{name: "Test2", args: args{command: "NOTecho hello | NOTecho"}, wantSane: false},
	}

	if runtime.GOOS == "linux" {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if gotSane := validWindowsCommand(tt.args.command); gotSane != tt.wantSane {
					t.Errorf("validWindowsCommand() = %v, want %v", gotSane, tt.wantSane)
				}
			})
		}
	}

}

func Test_validDarwinCommand(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name     string
		args     args
		wantSane bool
	}{
		{name: "Test1", args: args{command: "echo hello | echo"}, wantSane: true},
		{name: "Test2", args: args{command: "NOTecho hello | NOTecho"}, wantSane: false},
	}

	if runtime.GOOS == "darwin" {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if gotSane := validWindowsCommand(tt.args.command); gotSane != tt.wantSane {
					t.Errorf("validWindowsCommand() = %v, want %v", gotSane, tt.wantSane)
				}
			})
		}
	}

}
