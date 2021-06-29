package executor

import (
	"Auditheia/memory"
	"reflect"
	"testing"

	"github.com/dop251/goja"
)

func TestRunScript(t *testing.T) {
	type args struct {
		task *memory.Task
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test1", args: args{memory.TaskForTest_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, "")}, wantErr: false},
		{name: "Test2", args: args{memory.TaskForTest_Builder("script", "let result=nil;", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, "")}, wantErr: true},
		{name: "Test3", args: args{memory.TaskForTest_Builder("script", "", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, "")}, wantErr: true},
		{name: "Test4", args: args{memory.TaskForTest_Builder("script", "DasIstKeinExec", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, "")}, wantErr: true},
		//{name: "Test5", args: args{TaskForTest_Builder("script", "DasIstKeinExec", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, "", "")}, wantErr: true},
		{name: "Test6", args: args{nil}, wantErr: false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RunScript(tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("RunScript() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_initRuntime(t *testing.T) {
	type args struct {
		task *memory.Task
	}
	tests := []struct {
		name    string
		args    args
		want    *goja.Runtime
		wantErr bool
	}{
		//	{name: "Test2", args: args{memory.TaskForTest_Builder("script", "", []string{}, "", nil, nil, nil, "", "", "", "", true, "", "")}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := initRuntime(tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("initRuntime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initRuntime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newAuditheia(t *testing.T) {
	type args struct {
		vm   *goja.Runtime
		task *memory.Task
	}
	tests := []struct {
		name    string
		args    args
		want    *goja.Object
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newAuditheia(tt.args.vm, tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("newAuditheia() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newAuditheia() = %v, want %v", got, tt.want)
			}
		})
	}
}
