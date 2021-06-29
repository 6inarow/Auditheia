package executor

import (
	"Auditheia/memory"
	"os/exec"
	"reflect"
	"runtime"
	"testing"
)

func TestExecuteAudit(t *testing.T) {
	type args struct {
		audit *memory.Audit
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test1", args: args{memory.AuditForTest_Builder("Audit1", nil, "Status", "Message")}, wantErr: true},
		{name: "Test2", args: args{nil}, wantErr: true},
		{name: "Test3", args: args{memory.AuditForTest_Builder("Audit2", []memory.Task{*memory.Task_Builder("command", "echo hello", []string{"result"}, "equals", nil, nil, nil, "result", "status", "setStatus", "ArtifactName", true, "content")}, "Status", "Message")}, wantErr: false},
		//{name: "Test4", args: args{memory.AuditForTest_Builder("Audit3", []memory.Task{memory.TaskNilBuilder()}, "", "")}, wantErr: true},
		//{name: "Test4", args: args{memory.AuditForTest_Builder("Audit1", []memory.Task{memory.Task_Builder("command", "", []string{}, "operator", nil, nil, nil, "result", "status", "setStatus", "ArtifactName", true, "content", "path")}, "Status", "Message")}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExecuteAudit(tt.args.audit); (err != nil) != tt.wantErr {
				t.Errorf("ExecuteAudit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_executeTask(t *testing.T) {
	type args struct {
		task *memory.Task
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test1", args: args{memory.TaskForTest_Builder("command", "echo", []string{"SameSame"}, "equals", nil, nil, nil, "SameSame", "setStatus", "StatusMsg", "ArtifactName", true, "content")}, wantErr: false},
		{name: "Test2", args: args{memory.TaskForTest_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "setStatus", "StatusMsg", "ArtifactName", true, "content")}, wantErr: false},
		{name: "Test3", args: args{memory.TaskForTest_Builder("NotKnown", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "setStatus", "StatusMsg", "ArtifactName", true, "content")}, wantErr: false},
		{name: "Test4", args: args{memory.TaskForTest_Builder("command", "let result='SameSame';", []string{"SameSame"}, "equals", memory.TaskForTest_Builder("command", "echo", []string{"SameSame"}, "equals", nil, nil, nil, "SameSame", "setStatus", "StatusMsg", "ArtifactName", true, "content"), nil, nil, "", "setStatus", "StatusMsg", "ArtifactName", true, "content")}, wantErr: false},
		{name: "Test5", args: args{memory.TaskForTest_Builder("command", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "setStatus", "StatusMsg", "ArtifactName", true, "content")}, wantErr: false},
		{name: "Test6", args: args{nil}, wantErr: true},

		//{name: "Test4", args: args{TaskForTest_Builder("default", "test", []string{"SameSame"}, "operator", nil, nil, nil, "SameSame", "setStatus", "StatusMsg", "ArtifactName", true, "content", "path")}, wantErr: false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := executeTask(tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("executeTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_executeTaskTypeCommand(t *testing.T) {
	type args struct {
		task *memory.Task
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test1", args: args{memory.TaskForTest_Builder("command", "echo", []string{"SameSame"}, "equals", nil, nil, nil, "SameSame", "setStatus", "StatusMsg", "ArtifactName", true, "content")}, wantErr: false},
		{name: "Test2", args: args{memory.TaskForTest_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "setStatus", "StatusMsg", "ArtifactName", true, "content")}, wantErr: false},
		{name: "Test3", args: args{nil}, wantErr: true},
		{name: "Test4", args: args{memory.TaskForTest_Builder("command", "", []string{"SameSame"}, "equals", nil, nil, nil, "SameSame", "setStatus", "StatusMsg", "ArtifactName", true, "content")}, wantErr: true},
		{name: "Test5", args: args{memory.TaskForTest_Builder("command", "", []string{"SameSame"}, "equals", nil, nil, nil, "SameSame", "setStatus", "StatusMsg", "ArtifactName", true, "content")}, wantErr: true},
		{name: "Test5", args: args{memory.TaskForTest_Builder("command", "", []string{"SameSame"}, "equals", nil, nil, nil, "SameSame", "setStatus", "StatusMsg", "ArtifactName", true, "content")}, wantErr: true},
		//  {name: "Test6", args: args{memory.TaskForTest_Builder("script", "let result = nil;", []string{"SameSame"}, "equals", nil, nil, nil, "", "setStatus", "StatusMsg", "ArtifactName", true, "content", "path")}, wantErr: true},
		//	{name: "Test6", args: args{memory.TaskForTest_Builder("script", "let execTest = nil;let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "SameSame", "setStatus", "StatusMsg", "ArtifactName", true, "content", "path")}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := executeTaskTypeCommand(tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("executeTaskTypeCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_checkResult(t *testing.T) {
	type args struct {
		//task *memory.Task
		result   string
		expected []string
		operator string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		// Pos. & Neg. Tests für Operatoren equals/not/contains
		{name: "Test1_Operator_equals",
			args: args{result: "SameSame", expected: []string{"SameSame"}, operator: "equals"},
			want: true},
		{name: "Test2_Operator_not",
			args: args{result: "SameSame", expected: []string{"NotSameSame"}, operator: "not equals"},
			want: true},
		{name: "Test3_Operator_contains",
			args: args{result: "notsamesame", expected: []string{"samesame"}, operator: "contains"},
			want: true},
		{name: "Test4_Operator_equals_false",
			args: args{result: "SameSameNe", expected: []string{"SameSameNeg"}, operator: "equals"},
			want: false},
		{name: "Test5_Operator_not_false",
			args: args{result: "NotSameSameNeg", expected: []string{"NotSameSameNeg"}, operator: "not equals"},
			want: false},
		{name: "Test6_Operator_contains_false",
			args: args{result: "samesameNe", expected: []string{"samesameNeg"}, operator: "contains"},
			want: false},

		// Success und Fsail im selben Expected [Array] - Test
		{name: "Test7_Operator_equals",
			args: args{result: "SameSame", expected: []string{"SameSame", "NotSame"}, operator: "equals"},
			want: true},
		{name: "Test8_Operator_not",
			args: args{result: "SameSame", expected: []string{"NotSameSame", "SameSame"}, operator: "not equals"},
			want: true},
		{name: "Test9_Operator_contains",
			args: args{result: "notsamesame", expected: []string{"not", "ton"}, operator: "contains"},
			want: true},
		{name: "Test10_Operator_equals_false",
			args: args{result: "SameSame", expected: []string{"SameSameNeg", "SameSame"}, operator: "equals"},
			want: true},
		{name: "Test11_Operator_not_false",
			args: args{result: "NotSameSame", expected: []string{"NotSameSame", "SameSame"}, operator: "not equals"},
			want: true},
		{name: "Test12_Operator_contains_false",
			args: args{result: "samesameNe", expected: []string{"zz", "Ne"}, operator: "contains"},
			want: true},

		// Success und Fail im selben Expected [Array] - Test
		{name: "Test13_Operator_equals",
			args: args{result: "SameSame", expected: []string{"NotSame", "SameSame"}, operator: "equals"},
			want: true},
		{name: "Test14_Operator_not",
			args: args{result: "SameSame", expected: []string{"SameSame", "NotSameSame"}, operator: "not equals"},
			want: true},
		{name: "Tes15_Operator_contains",
			args: args{result: "notsamesame", expected: []string{"ton", "not"}, operator: "contains"},
			want: true},
		{name: "Test16_Operator_equals_false",
			args: args{result: "SameSame", expected: []string{"SameSame", "SameSameNeg"}, operator: "equals"},
			want: true},
		{name: "Test17_Operator_not_false",
			args: args{result: "NotSameSame", expected: []string{"SameSame", "NotSameSame"}, operator: "not equals"},
			want: true},
		{name: "Test18_Operator_contains_false",
			args: args{result: "samesameNe", expected: []string{"Ne", "zz"}, operator: "contains"},
			want: true},

		{name: "Test19_Operator_not_contains",
			args: args{result: "abcdefg", expected: []string{"zyx"}, operator: "not contains"},
			want: true},
		{name: "Test20_Operator_lesser",
			args: args{result: "100", expected: []string{"101"}, operator: "lesser"},
			want: true},
		{name: "Test21_Operator_greater",
			args: args{result: "100", expected: []string{"99"}, operator: "greater"},
			want: true},
		{name: "Test22_False_Operator",
			args: args{result: "SameSameNe", expected: []string{"SameSameNeg"}, operator: "NotAnOperator"},
			want: false},
		{name: "Test23_Operator_lesser_neg",
			args: args{result: "100", expected: []string{"1"}, operator: "lesser"},
			want: false},
		{name: "Test24_Operator_greater_neg",
			args: args{result: "100", expected: []string{"1000"}, operator: "greater"},
			want: false},
		{name: "Test25__Operator_not_contains_neg",
			args: args{result: "abcdefg", expected: []string{"abcdefg"}, operator: "not contains"},
			want: false},

		{name: "Test26_Operator_",
			args: args{result: "NoNumber", expected: []string{"101"}, operator: "lesser"},
			want: false},
		{name: "Test27_Operator_greater",
			args: args{result: "100", expected: []string{"NoNumber"}, operator: "greater"},
			want: false},
		{name: "Test28_Operator_lesser_neg",
			args: args{result: "100", expected: []string{"NoNumber"}, operator: "lesser"},
			want: false},
		{name: "Test29_Operator_greater_neg",
			args: args{result: "NoNumber", expected: []string{"1000"}, operator: "greater"},
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkResult(tt.args.result, tt.args.expected, tt.args.operator); got != tt.want {
				t.Errorf("checkResult() = %v, want %v", got, tt.want)

			}
		})
	}
}

func Test_saveArtefact(t *testing.T) {
	type args struct {
		task   *memory.Task
		s      exec.Cmd
		stdout []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test1", args: args{memory.TaskForTest_Builder("command1", "", []string{"testString"}, "not", nil, nil, nil, "testStringNot", "", "", "", true, ""),
			memory.CmdForTest_Builder("", []string{"hello"}, []string{"hello"}, "", nil, nil, nil, nil, nil, nil), []byte{0, 1, 0, 1}}, want: true},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := saveArtefact(tt.args.task, tt.args.s, tt.args.stdout); got != tt.want {
				t.Errorf("saveArtefact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_executeTaskTypeScript(t *testing.T) {
	type args struct {
		task *memory.Task
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test1", args: args{memory.TaskForTest_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, "")}, wantErr: false},
		{name: "Test2", args: args{nil}, wantErr: true},
		{name: "Test3", args: args{memory.TaskForTest_Builder("script", "FailedScript", []string{"SameSame"}, "equals", nil, nil, nil, "SameSame", "", "", "", true, "")}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := executeTaskTypeScript(tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("executeTaskTypeScript() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_runCommandFromScript(t *testing.T) {
	type args struct {
		command string
		task    *memory.Task
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Test1", args: args{command: "", task: memory.TaskForTest_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, "")}, want: "", wantErr: true},
		//	{name: "Test2", args: args{command: "echo hallo", task: memory.TaskForTest_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, "", "")}, want: "hallo", wantErr: false},
		//{name: "Test3", args: args{command: "whoami", task: memory.TaskForTest_Builder("script", "", []string{"SameSame"}, "", nil, nil, nil, "", "", "", "", true, "", "")}, want: "robert-paulson\\heiopei", wantErr: false},
		//{name: "Test3", args: args{command: "cmd /C echo hallo | echo", task: memory.TaskForTest_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, "", "")}, want: "hallo", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runCommandFromScript(tt.args.command, tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("runCommandFromScript() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("runCommandFromScript() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func Test_buildCommandsFromString(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name    string
		args    args
		want    []exec.Cmd
		wantErr bool
	}{
		// TODO: GOOS only resembles the OS of the Current Runner. Therefore it cant be Unit-Tested via GitHub Actions
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildCommandsFromString(tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildCommandsFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildCommandsFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildCommandsFromStringWindows(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name            string
		args            args
		wantCommandList []exec.Cmd
		wantErr         bool
	}{
		// TODO: Add test cases.
		{name: "Test1", args: args{""}, wantCommandList: nil, wantErr: true},
		//	{name: "Test2", args: args{"echo hallo"}, wantCommandList: []exec.Cmd{CmdForTest_Builder("C:\\Programme\\Git\\usr\\bin\\echo.exe", []string{"/C hallo"}, []string{}, "", nil, nil, nil, nil, nil, nil) }, wantErr: false},
		// TODO: Output and WantCommandList doesn´t match // On different OS Unit Test wont work
	}
	if runtime.GOOS == "windows" {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotCommandList, err := buildCommandsFromStringWindows(tt.args.command)
				if (err != nil) != tt.wantErr {
					t.Errorf("buildCommandsFromStringWindows() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotCommandList, tt.wantCommandList) {
					t.Errorf("buildCommandsFromStringWindows() = %v, want %v", gotCommandList, tt.wantCommandList)
				}
			})
		}
	}

}

func Test_prependString(t *testing.T) {
	type args struct {
		source []string
		add    []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "Test1", args: args{[]string{"Source"}, []string{"Append"}}, want: []string{"Append", "Source"}},
		{name: "Test2", args: args{[]string{""}, []string{"Butter"}}, want: []string{"Butter", ""}},
		{name: "Test3", args: args{[]string{"Brot", "Marmelade", "Butter"}, []string{"Eier"}}, want: []string{"Eier", "Brot", "Marmelade", "Butter"}},
		{name: "Test4", args: args{[]string{"Brot"}, []string{"Eier", "Märmelade", "Butter"}}, want: []string{"Butter", "Märmelade", "Eier", "Brot"}},

		// TODO: Check the boundary value
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prependString(tt.args.source, tt.args.add...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prependString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cleanShellOutputs(t *testing.T) {
	type args struct {
		toClean string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test1", args: args{"Hallo Welt\n"}, want: "Hallo Welt"},
		{name: "Test2", args: args{"Hallo Welt\nTest"}, want: "Hallo Welt\nTest"},
		{name: "Test3", args: args{"\\nHallo Welt"}, want: "\\nHallo Welt"},
		{name: "Test4", args: args{"\nHallo Welt"}, want: "\nHallo Welt"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanShellOutputs(tt.args.toClean); got != tt.want {
				t.Errorf("cleanShellOutputs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_detectArtefactFilesInArgs1(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectArtefactFilesInArgs(tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectArtefactFilesInArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
