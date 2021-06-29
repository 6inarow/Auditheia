package output

import (
	"Auditheia/memory/constants"
	"archive/zip"
	"reflect"
	"runtime"
	"testing"
)

func TestZipFolder_local1(t *testing.T) {
	type args struct {
		zipfolder string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test1", args: args{""}, wantErr: true},
		{name: "Test2", args: args{"C:\\Users\\Heiopei\\Documents\\ZipTest\\TestZipFolder.zip"}, wantErr: true},
		{name: "Test3", args: args{"C:\\Users\\Heiopei\\Documents\\ZipTest\\TestZipFolderUnziped\\"}, wantErr: false},
	}
	if runtime.GOOS == "windows" && constants.LOCAL_1 {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := ZipFolder(tt.args.zipfolder); (err != nil) != tt.wantErr {
					t.Errorf("ZipFolder() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func TestZipFolder_local2(t *testing.T) {
	type args struct {
		zipfolder string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test1", args: args{""}, wantErr: true},
		{name: "Test2", args: args{"C:\\Users\\spark\\Documents\\ZipTest\\TestZipFolder.zip"}, wantErr: true},
		{name: "Test3", args: args{"C:\\Users\\spark\\Documents\\ZipTest\\TestZipFolderUnziped\\"}, wantErr: false},
	}
	if runtime.GOOS == "windows" && constants.LOCAL_2 {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := ZipFolder(tt.args.zipfolder); (err != nil) != tt.wantErr {
					t.Errorf("ZipFolder() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}

}

func TestZipFolder_windows(t *testing.T) {
	type args struct {
		zipfolder string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// create zip and directory on OS
		{name: "Test1", args: args{""}, wantErr: true},
		{name: "Test2", args: args{"D:\\a\\Auditheia\\PackageOutputTest\\TestZipFolder.zip"}, wantErr: true},
		{name: "Test3", args: args{"D:\\NoPath\\To\\NoFile"}, wantErr: true},
		{name: "Test4", args: args{"D:\\a\\Auditheia\\PackageOutputTest\\TestZipFolderUnziped"}, wantErr: false},
	}
	if runtime.GOOS == "windows" && constants.RUNNER {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := ZipFolder(tt.args.zipfolder); (err != nil) != tt.wantErr {
					t.Errorf("ZipFolder() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}

}

func TestZipFolder_linux(t *testing.T) {
	type args struct {
		zipfolder string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test1", args: args{""}, wantErr: true},
		{name: "Test2", args: args{"/home/runner/work/Auditheia/PackageOutputTest/TestZipFolder.zip"}, wantErr: true},
		{name: "Test3", args: args{"/NoPath/To/NoFile"}, wantErr: true},
		{name: "Test4", args: args{"/home/runner/work/Auditheia/PackageOutputTest/TestZipFolderUnziped/"}, wantErr: false},
		//{name: "Test4", args: args{"/mnt/d/WRKPLC/Auditheia/PackageOutputTest/TestZipFolderUnziped/"}, wantErr: false},
		//{name: "Test4", args: args{"C:\\Users\\Heiopei\\Documents\\ZipTest\\ZipTestUngezipt\\Test2.txt"}, wantErr: true},
	}
	if runtime.GOOS == "linux" && constants.RUNNER {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := ZipFolder(tt.args.zipfolder); (err != nil) != tt.wantErr {
					t.Errorf("ZipFolder() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}

}

func TestZipFolder_macos(t *testing.T) {
	type args struct {
		zipfolder string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test1", args: args{""}, wantErr: true},
		{name: "Test2", args: args{"/Users/runner/work/Auditheia/PackageOutputTest/TestZipFolder.zip"}, wantErr: true},
		{name: "Test3", args: args{"/NoPath/To/NoFile"}, wantErr: true},
		{name: "Test4", args: args{"/Users/runner/work/Auditheia/PackageOutputTest/TestZipFolderUnziped/"}, wantErr: false},
		//{name: "Test4", args: args{"/mnt/d/WRKPLC/Auditheia/PackageOutputTest/TestZipFolderUnziped/"}, wantErr: false},
		//{name: "Test4", args: args{"C:\\Users\\Heiopei\\Documents\\ZipTest\\ZipTestUngezipt\\Test2.txt"}, wantErr: true},
	}
	if runtime.GOOS == "darwin" && constants.RUNNER {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := ZipFolder(tt.args.zipfolder); (err != nil) != tt.wantErr {
					t.Errorf("ZipFolder() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}

}

func Test_readDirectory_local1(t *testing.T) {
	type args struct {
		directory string
	}
	tests := []struct {
		name                string
		args                args
		wantDirectoryStruct []string
		wantErr             bool
	}{
		{name: "Test1", args: args{"C:\\Users\\Heiopei\\Documents\\ZipTest2\\readDirectory"}, wantDirectoryStruct: []string{"C:\\Users\\Heiopei\\Documents\\ZipTest2\\readDirectory\\Test.txt"}, wantErr: false},
		{name: "Test2", args: args{""}, wantDirectoryStruct: nil, wantErr: true},
	}
	if runtime.GOOS == "windows" && constants.LOCAL_1 {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotDirectoryStruct, err := readDirectory(tt.args.directory)
				if (err != nil) != tt.wantErr {
					t.Errorf("readDirectory() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotDirectoryStruct, tt.wantDirectoryStruct) {
					t.Errorf("readDirectory() = %v, want %v", gotDirectoryStruct, tt.wantDirectoryStruct)
				}
			})
		}
	}

}

func Test_readDirectory_local2(t *testing.T) {
	type args struct {
		directory string
	}
	tests := []struct {
		name                string
		args                args
		wantDirectoryStruct []string
		wantErr             bool
	}{
		{name: "Test1", args: args{"C:\\Users\\spark\\Documents\\ZipTest2\\readDirectory"}, wantDirectoryStruct: []string{"C:\\Users\\spark\\Documents\\ZipTest2\\readDirectory\\Test.txt"}, wantErr: false},
		{name: "Test2", args: args{""}, wantDirectoryStruct: nil, wantErr: true},
	}
	if runtime.GOOS == "windows" && constants.LOCAL_2 {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotDirectoryStruct, err := readDirectory(tt.args.directory)
				if (err != nil) != tt.wantErr {
					t.Errorf("readDirectory() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotDirectoryStruct, tt.wantDirectoryStruct) {
					t.Errorf("readDirectory() = %v, want %v", gotDirectoryStruct, tt.wantDirectoryStruct)
				}
			})
		}
	}

}

func Test_readDirectory_macos(t *testing.T) {
	type args struct {
		directory string
	}
	tests := []struct {
		name                string
		args                args
		wantDirectoryStruct []string
		wantErr             bool
	}{
		{name: "Test1", args: args{"/Users/runner/work/Auditheia/PackageOutputTest/readDirectoryDir"}, wantDirectoryStruct: []string{"/Users/runner/work/Auditheia/PackageOutputTest/readDirectoryDir/foo.txt"}, wantErr: false},
		{name: "Test2", args: args{""}, wantDirectoryStruct: nil, wantErr: true},
	}
	if runtime.GOOS == "darwin" && constants.RUNNER {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotDirectoryStruct, err := readDirectory(tt.args.directory)
				if (err != nil) != tt.wantErr {
					t.Errorf("readDirectory() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotDirectoryStruct, tt.wantDirectoryStruct) {
					t.Errorf("readDirectory() = %v, want %v", gotDirectoryStruct, tt.wantDirectoryStruct)
				}
			})
		}
	}

}

func Test_readDirectory_windows(t *testing.T) {
	type args struct {
		directory string
	}
	tests := []struct {
		name                string
		args                args
		wantDirectoryStruct []string
		wantErr             bool
	}{
		{name: "Test1", args: args{"D:\\a\\Auditheia\\PackageOutputTest\\readDirectoryDir"}, wantDirectoryStruct: []string{"D:\\a\\Auditheia\\PackageOutputTest\\readDirectoryDir\\foo.txt"}, wantErr: false},
		{name: "Test2", args: args{""}, wantDirectoryStruct: nil, wantErr: true},
	}

	if runtime.GOOS == "windows" && constants.RUNNER {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotDirectoryStruct, err := readDirectory(tt.args.directory)
				if (err != nil) != tt.wantErr {
					t.Errorf("readDirectory() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotDirectoryStruct, tt.wantDirectoryStruct) {
					t.Errorf("readDirectory() = %v, want %v", gotDirectoryStruct, tt.wantDirectoryStruct)
				}
			})
		}
	}
}

func Test_readDirectory_linux(t *testing.T) {
	type args struct {
		directory string
	}
	tests := []struct {
		name                string
		args                args
		wantDirectoryStruct []string
		wantErr             bool
	}{
		// create directory and foo.txt
		{name: "Test1", args: args{"/home/runner/work/Auditheia/PackageOutputTest/readDirectoryDir"}, wantDirectoryStruct: []string{"/home/runner/work/Auditheia/PackageOutputTest/readDirectoryDir/foo.txt"}, wantErr: false},
		{name: "Test2", args: args{""}, wantDirectoryStruct: nil, wantErr: true},
	}

	if runtime.GOOS == "linux" && constants.RUNNER {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotDirectoryStruct, err := readDirectory(tt.args.directory)
				if (err != nil) != tt.wantErr {
					t.Errorf("readDirectory() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotDirectoryStruct, tt.wantDirectoryStruct) {
					t.Errorf("readDirectory() = %v, want %v", gotDirectoryStruct, tt.wantDirectoryStruct)
				}
			})
		}
	}
}

func Test_zipFiles_local1(t *testing.T) {
	type args struct {
		filename string
		files    []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test1", args: args{"", []string{""}}, wantErr: true},
		{name: "Test2", args: args{"C:\\Users\\Heiopei\\Documents\\ZipTest\\TestZip2.zip", []string{"C:\\Users\\Heiopei\\Documents\\ZipTest\\Test.txt", "C:\\Users\\Heiopei\\Documents\\ZipTest\\Test2.txt"}}, wantErr: false},
		{name: "Test3", args: args{"", []string{"", ""}}, wantErr: true},
	}
	if runtime.GOOS == "windows" && constants.LOCAL_1 {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := zipFiles(tt.args.filename, tt.args.files); (err != nil) != tt.wantErr {
					t.Errorf("zipFiles() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}

}

func Test_zipFiles_local2(t *testing.T) {
	type args struct {
		filename string
		files    []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test1", args: args{"", []string{""}}, wantErr: true},
		{name: "Test2", args: args{"C:\\Users\\spark\\Documents\\ZipTest\\TestZip2.zip", []string{"C:\\Users\\spark\\Documents\\ZipTest\\Test.txt", "C:\\Users\\spark\\Documents\\ZipTest\\Test2.txt"}}, wantErr: false},
		{name: "Test3", args: args{"", []string{"", ""}}, wantErr: true},
	}
	if runtime.GOOS == "windows" && constants.LOCAL_2 {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := zipFiles(tt.args.filename, tt.args.files); (err != nil) != tt.wantErr {
					t.Errorf("zipFiles() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}

}

func Test_zipFiles_darwin(t *testing.T) {
	type args struct {
		filename string
		files    []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test1", args: args{"", []string{""}}, wantErr: true},
		{name: "Test2", args: args{"/Users/runner/work/Auditheia/PackageOutputTest/zipFilesTest.zip", []string{"/Users/runner/work/Auditheia/PackageOutputTest/foo.txt", "/Users/runner/work/Auditheia/PackageOutputTest/foo2.txt"}}, wantErr: false},
		{name: "Test3", args: args{"", []string{"", ""}}, wantErr: true},
	}
	if runtime.GOOS == "darwin" && constants.RUNNER {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := zipFiles(tt.args.filename, tt.args.files); (err != nil) != tt.wantErr {
					t.Errorf("zipFiles() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}

}

func Test_zipFiles_windows(t *testing.T) {
	type args struct {
		filename string
		files    []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test1", args: args{"", []string{""}}, wantErr: true},
		{name: "Test2", args: args{"D:\\a\\Auditheia\\PackageOutputTest\\zipFilesTest.zip", []string{"D:\\a\\Auditheia\\PackageOutputTest\\foo.txt", "D:\\a\\Auditheia\\PackageOutputTest\\foo2.txt"}}, wantErr: false},
		//{name: "Test3", args: args{"D:\\a\\Auditheia\\PackageOutputTest!'`\\zumZüppen.zip", []string{"D:\\a\\Auditheia\\PackageOutputTest!'`\\zumZüppen\\`Tüst!'.txt", "D:\\a\\Auditheia\\PackageOutputTest!'`\\Test2.txt"}}, wantErr: false},
	}
	if runtime.GOOS == "windows" && constants.RUNNER {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := zipFiles(tt.args.filename, tt.args.files); (err != nil) != tt.wantErr {
					t.Errorf("zipFiles() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}

}

func Test_zipFiles_linux(t *testing.T) {
	type args struct {
		filename string
		files    []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test1", args: args{"", []string{""}}, wantErr: true},
		{name: "Test2", args: args{"/home/runner/work/Auditheia/PackageOutputTest/zipFilesTest.zip", []string{"/home/runner/work/Auditheia/PackageOutputTest/foo.txt", "/home/runner/work/Auditheia/PackageOutputTest/foo2.txt"}}, wantErr: false},
	}
	if runtime.GOOS == "linux" && constants.RUNNER {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := zipFiles(tt.args.filename, tt.args.files); (err != nil) != tt.wantErr {
					t.Errorf("zipFiles() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func Test_addFileToZip(t *testing.T) {
	type args struct {
		zipWriter *zip.Writer
		filename  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test1", args: args{nil, ""}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := addFileToZip(tt.args.zipWriter, tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("addFileToZip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
