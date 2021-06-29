package output

import (
	"Auditheia/memory"
	"reflect"
	"runtime"
	"testing"
)

func Test_report_ToJSON(t *testing.T) {
	type args struct {
		pretty bool
	}
	tests := []struct {
		name    string
		r       *report
		args    args
		want    string
		wantErr bool
	}{
		//{name: "Test1", r: nil, args: args{true}, want: "Test", wantErr: true},
		//	{name: "Test2", r: *memory.ReportForTest_Builder("", "", "", "", []memory.AuditDTO{*memory.AuditDTOForTest_Builder("", []memory.TaskDTO{*memory.TaskDTOForTest_Builder("", "", nil, "", "", "")}, "", "")}, ""), args: args{true}, want: "Test", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ToJSON(tt.args.pretty)
			if (err != nil) != tt.wantErr {
				t.Errorf("report.ToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("report.ToJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_report_WriteToFile(t *testing.T) {
	tests := []struct {
		name    string
		r       report
		wantErr bool
	}{
		//	{name: "Test1", r: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.WriteToFile(); (err != nil) != tt.wantErr {
				t.Errorf("report.WriteToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_report_reportEmpty(t *testing.T) {
	tests := []struct {
		name string
		r    *report
		want bool
	}{
		//{name: "Test1", r: nil, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.reportEmpty(); got != tt.want {
				t.Errorf("report.reportEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveArtefacts(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "Test1", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveArtefacts(); (err != nil) != tt.wantErr {
				t.Errorf("SaveArtefacts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_saveAdditionalFiles(t *testing.T) {
	type args struct {
		parentDir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test1", args: args{""}, wantErr: true},
		{name: "Test2", args: args{"C:\\WrongDirectory"}, wantErr: true},
		//{name: "Test3", args: args{"D:\\WRKPLC\\Auditheia\\saveAdditionalFiles"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := saveAdditionalFiles(tt.args.parentDir); (err != nil) != tt.wantErr {
				t.Errorf("saveAdditionalFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_copyFile_windows(t *testing.T) {
	type args struct {
		source      string
		destination string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test1", args: args{"source", "dest"}, wantErr: true},
		{name: "Test2", args: args{"source", "source"}, wantErr: true},
		{name: "Test3", args: args{"D:\\a\\Auditheia\\PackageOutputTest\\copyTest\\copyFrom.txt", "D:\\a\\Auditheia\\PackageOutputTest\\copyTest\\copyTo.txt"}, wantErr: false},
		{name: "Test4", args: args{"D:\\a\\Auditheia\\PackageOutputTest\\copyTest\\copyFrom.txt", "X:\\copyTo.txt"}, wantErr: true},
	}
	if runtime.GOOS == "windows" {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := copyFile(tt.args.source, tt.args.destination); (err != nil) != tt.wantErr {
					t.Errorf("copyFile() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func Test_copyFile_linux(t *testing.T) {
	type args struct {
		source      string
		destination string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test1", args: args{"source", "dest"}, wantErr: true},
		{name: "Test2", args: args{"source", "source"}, wantErr: true},
		{name: "Test3", args: args{"/home/runner/work/Auditheia/PackageOutputTest/copyTest/copyFrom.txt", "/home/runner/work/Auditheia/PackageOutputTest/copyTest/copyTo.txt"}, wantErr: false},
		{name: "Test4", args: args{"/home/runner/work/Auditheia/PackageOutputTest/copyTest/copyFrom.txt", "/etc/copyTo.txt"}, wantErr: true},
	}
	if runtime.GOOS == "linux" {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := copyFile(tt.args.source, tt.args.destination); (err != nil) != tt.wantErr {
					t.Errorf("copyFile() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func Test_copyFile_macOS(t *testing.T) {
	type args struct {
		source      string
		destination string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test1", args: args{"source", "dest"}, wantErr: true},
		{name: "Test2", args: args{"source", "source"}, wantErr: true},
		{name: "Test3", args: args{"/Users/runner/work/Auditheia/PackageOutputTest/copyTest/copyFrom.txt", "/Users/runner/work/Auditheia/PackageOutputTest/copyTest/copyTo.txt"}, wantErr: false},
		{name: "Test4", args: args{"/Users/runner/work/Auditheia/PackageOutputTest/copyTest/copyFrom.txt", "/etc/copyTo.txt"}, wantErr: true},
	}
	if runtime.GOOS == "darwin" {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := copyFile(tt.args.source, tt.args.destination); (err != nil) != tt.wantErr {
					t.Errorf("copyFile() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func Test_writeArtefactsToDirectory_windows(t *testing.T) {
	type args struct {
		artefacts  []memory.Artefact
		parentDirs []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test1", args: args{[]memory.Artefact{}, []string{"ParentsDir"}}, wantErr: false},
		{name: "Test2", args: args{nil, nil}, wantErr: false},
		{name: "Test3", args: args{[]memory.Artefact{memory.Artefact_Builder("Name", "Content", "Path")}, []string{"D:\\a\\Auditheia\\PackageOutputTest\\FilesToDir"}}, wantErr: false},
		{name: "Test4", args: args{[]memory.Artefact{memory.Artefact_Builder("Name", "Content", "Path")}, []string{"X:\\a\\Auditheia\\PackageOutputTest\\FilesToDir"}}, wantErr: true},
	}
	if runtime.GOOS == "windows" {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := writeArtefactsToDirectory(tt.args.artefacts, tt.args.parentDirs...); (err != nil) != tt.wantErr {
					t.Errorf("writeArtefactsToDirectory() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func Test_writeArtefactsToDirectory_linux(t *testing.T) {
	type args struct {
		artefacts  []memory.Artefact
		parentDirs []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test1", args: args{[]memory.Artefact{}, []string{"ParentsDir"}}, wantErr: false},
		{name: "Test2", args: args{nil, nil}, wantErr: false},
		{name: "Test3", args: args{[]memory.Artefact{memory.Artefact_Builder("Name", "Content", "Path")}, []string{"/home/runner/work/Auditheia/PackageOutputTest/FilesToDir"}}, wantErr: false},
		{name: "Test4", args: args{[]memory.Artefact{memory.Artefact_Builder("Name", "Content", "Path")}, []string{"/etc/runner/work/Auditheia/PackageOutputTest/FilesToDir"}}, wantErr: true},
	}
	if runtime.GOOS == "linux" {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := writeArtefactsToDirectory(tt.args.artefacts, tt.args.parentDirs...); (err != nil) != tt.wantErr {
					t.Errorf("writeArtefactsToDirectory() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func Test_writeArtefactsToDirectory_macOS(t *testing.T) {
	type args struct {
		artefacts  []memory.Artefact
		parentDirs []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test1", args: args{[]memory.Artefact{}, []string{"ParentsDir"}}, wantErr: false},
		{name: "Test2", args: args{nil, nil}, wantErr: false},
		{name: "Test3", args: args{[]memory.Artefact{memory.Artefact_Builder("Name", "Content", "Path")}, []string{"/Users/runner/work/Auditheia/PackageOutputTest"}}, wantErr: false},
		{name: "Test4", args: args{[]memory.Artefact{memory.Artefact_Builder("Name", "Content", "Path")}, []string{"/NoDir/PackageOutputTest"}}, wantErr: true},
	}
	if runtime.GOOS == "darwin" {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := writeArtefactsToDirectory(tt.args.artefacts, tt.args.parentDirs...); (err != nil) != tt.wantErr {
					t.Errorf("writeArtefactsToDirectory() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func Test_addTasks(t *testing.T) {
	type args struct {
		list []*memory.Task
		task *memory.Task
	}
	tests := []struct {
		name string
		args args
		want []*memory.Task
	}{
		{name: "Test1",
			args: args{[]*memory.Task{
				memory.TaskForTest_Builder("command", "", []string{}, "", nil, nil, nil, "", "", "", "", true, "")},
				memory.TaskForTest_Builder("script", "", []string{}, "", nil, nil, nil, "", "", "", "", true, "")},
			want: []*memory.Task{
				memory.TaskForTest_Builder("command", "", []string{}, "", nil, nil, nil, "", "", "", "", true, ""),
				memory.TaskForTest_Builder("script", "", []string{}, "", nil, nil, nil, "", "", "", "", true, "")},
		},
		{name: "Test2",
			args: args{[]*memory.Task{
				memory.TaskForTest_Builder("command", "echo Task1", []string{}, "", nil, nil, nil, "", "", "", "", true, "")},
				memory.TaskForTest_Builder("script", "echo Task2", []string{}, "",
					memory.TaskForTest_Builder("NotACommand", "echo Task3", []string{}, "", nil, nil, nil, "", "", "", "", true, ""),
					nil, nil, "", "", "", "", true, "")},
			want: []*memory.Task{
				memory.TaskForTest_Builder("command", "echo Task1", []string{}, "", nil, nil, nil, "", "", "", "", true, ""),
				memory.TaskForTest_Builder("script", "echo Task2", []string{}, "", memory.TaskForTest_Builder("NotACommand", "echo Task3", []string{}, "", nil, nil, nil, "", "", "", "", true, ""), nil, nil, "", "", "", "", true, ""),
				memory.TaskForTest_Builder("NotACommand", "echo Task3", []string{}, "", nil, nil, nil, "", "", "", "", true, "")},
		},
		{name: "Test3",
			args: args{[]*memory.Task{
				memory.TaskForTest_Builder("command", "echo Task1", []string{}, "", nil, nil, nil, "", "", "", "", true, "")},
				memory.TaskForTest_Builder("script", "echo Task2", []string{}, "", nil,
					memory.TaskForTest_Builder("NotACommand", "echo Task3", []string{}, "", nil, nil, nil, "", "", "", "", true, ""),
					nil, "", "", "", "", true, "")},
			want: []*memory.Task{
				memory.TaskForTest_Builder("command", "echo Task1", []string{}, "", nil, nil, nil, "", "", "", "", true, ""),
				memory.TaskForTest_Builder("script", "echo Task2", []string{}, "", nil, memory.TaskForTest_Builder("NotACommand", "echo Task3", []string{}, "", nil, nil, nil, "", "", "", "", true, ""), nil, "", "", "", "", true, ""),
				memory.TaskForTest_Builder("NotACommand", "echo Task3", []string{}, "", nil, nil, nil, "", "", "", "", true, "")},
		},
		{name: "Test4",
			args: args{[]*memory.Task{
				memory.TaskForTest_Builder("command", "echo Task1", []string{}, "", nil, nil, nil, "", "", "", "", true, "")},
				memory.TaskForTest_Builder("script", "echo Task2", []string{}, "", nil, nil,
					memory.TaskForTest_Builder("NotACommand", "echo Task3", []string{}, "", nil, nil, nil, "", "", "", "", true, ""),
					"", "", "", "", true, "")},
			want: []*memory.Task{
				memory.TaskForTest_Builder("command", "echo Task1", []string{}, "", nil, nil, nil, "", "", "", "", true, ""),
				memory.TaskForTest_Builder("script", "echo Task2", []string{}, "", nil, nil, memory.TaskForTest_Builder("NotACommand", "echo Task3", []string{}, "", nil, nil, nil, "", "", "", "", true, ""), "", "", "", "", true, ""),
				memory.TaskForTest_Builder("NotACommand", "echo Task3", []string{}, "", nil, nil, nil, "", "", "", "", true, "")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTasks(tt.args.list, tt.args.task); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateName(t *testing.T) {
	type args struct {
		parts []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test1", args: args{[]string{"parts"}}, want: "parts"},
		{name: "Test2", args: args{[]string{""}}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateName(tt.args.parts...); got != tt.want {
				t.Errorf("generateName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_report_updateElapsedTime(t *testing.T) {
	tests := []struct {
		name string
		r    *report
	}{
		//{name: "Test2", r: nil},
		//{name: "Test1", r: memory.ReportForTest_Builder2("Linux", "Robert-Paulson", "", "", []memory.AuditDTO{*memory.AuditDTOForTest_Builder("", []memory.TaskDTO{*memory.TaskDTOForTest_Builder("", "", nil, "", "", "")}, "", "")}, "")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.updateElapsedTime()
		})
	}
}

func TestGenerateReport(t *testing.T) {
	tests := []struct {
		name string
		want *report
	}{
		//{name: "Test1", want: *memory.ReportForTest_Builder("Linux", "Robert-Paulson", "", "", []memory.AuditDTO{*memory.AuditDTOForTest_Builder("", []memory.TaskDTO{*memory.TaskDTOForTest_Builder("", "", nil, "", "", "")}, "", "")}, ""), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateReport(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateReport() = %v, want %v", got, tt.want)
			}
		})
	}
}
