package osInfo

import (
	"Auditheia/memory/constants"
	"reflect"
	"runtime"
	"testing"
)

func Test_getWindowsInfo_positiv_local1(t *testing.T) {
	tests := []struct {
		name      string
		wantInfos map[string]string
		wantErr   bool
	}{
		{
			name:      "Test1",
			wantInfos: map[string]string{"Caption": "Microsoft Windows 10 Pro ", "Version": "10.0.19041 "},
			wantErr:   false,
		},
		//{ just possible to test for wantErr=false??

		// TODO: Add test cases.
	}
	if runtime.GOOS == "windows" && constants.LOCAL_1 {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotInfos, err := getWindowsInfo()
				if (err != nil) != tt.wantErr {
					t.Errorf("getWindowsInfo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotInfos, tt.wantInfos) {
					t.Errorf("getWindowsInfo() gotInfos = %v, want %v", gotInfos, tt.wantInfos)
				}
			})
		}
	}

}

func Test_getWindowsInfo_positiv_local2(t *testing.T) {
	tests := []struct {
		name      string
		wantInfos map[string]string
		wantErr   bool
	}{
		{
			name:      "Test1",
			wantInfos: map[string]string{"Caption": "Microsoft Windows 10 Home ", "Version": "10.0.19043 "},
			wantErr:   false,
		},
		//{ just possible to test for wantErr=false??

		// TODO: Add test cases.
	}
	if runtime.GOOS == "windows" && constants.LOCAL_2 {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotInfos, err := getWindowsInfo()
				if (err != nil) != tt.wantErr {
					t.Errorf("getWindowsInfo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotInfos, tt.wantInfos) {
					t.Errorf("getWindowsInfo() gotInfos = %v, want %v", gotInfos, tt.wantInfos)
				}
			})
		}
	}

}

func Test_getWindowsInfo_positiv_latest(t *testing.T) {
	tests := []struct {
		name      string
		wantInfos map[string]string
		wantErr   bool
	}{
		{
			name:      "Test1",
			wantInfos: map[string]string{"Caption": "Microsoft Windows Server 2019 Datacenter ", "Version": "10.0.17763 "},
			wantErr:   false,
		},
		//{ just possible to test for wantErr=false??

		// TODO: Add test cases.
	}
	if runtime.GOOS == "windows" && constants.RUNNER && constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotInfos, err := getWindowsInfo()
				if (err != nil) != tt.wantErr {
					t.Errorf("getWindowsInfo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotInfos, tt.wantInfos) {
					t.Errorf("getWindowsInfo() gotInfos = %v, want %v", gotInfos, tt.wantInfos)
				}
			})
		}
	}

}

func Test_getWindowsInfo_positiv(t *testing.T) {
	tests := []struct {
		name      string
		wantInfos map[string]string
		wantErr   bool
	}{
		{
			name:      "Test1",
			wantInfos: map[string]string{"Caption": "Microsoft Windows Server 2016 Datacenter ", "Version": "10.0.14393 "},
			wantErr:   false,
		},
		//{ just possible to test for wantErr=false??

		// TODO: Add test cases.
	}
	if runtime.GOOS == "windows" && constants.RUNNER && !constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotInfos, err := getWindowsInfo()
				if (err != nil) != tt.wantErr {
					t.Errorf("getWindowsInfo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotInfos, tt.wantInfos) {
					t.Errorf("getWindowsInfo() gotInfos = %v, want %v", gotInfos, tt.wantInfos)
				}
			})
		}
	}

}

func Test_getWindowsInfo_negativ_local1(t *testing.T) {
	tests := []struct {
		name      string
		wantInfos map[string]string
		wantErr   bool
	}{
		{
			name:      "Test1",
			wantInfos: nil,
			wantErr:   true,
		},
		//{ just possible to test for wantErr=false??

		// TODO: Add test cases.
	}
	if runtime.GOOS != "windows" && constants.LOCAL_1 {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotInfos, err := getWindowsInfo()
				if (err != nil) != tt.wantErr {
					t.Errorf("getWindowsInfo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotInfos, tt.wantInfos) {
					t.Errorf("getWindowsInfo() gotInfos = %v, want %v", gotInfos, tt.wantInfos)
				}
			})
		}
	}
}

func Test_getWindowsInfo_negativ_local2(t *testing.T) {
	tests := []struct {
		name      string
		wantInfos map[string]string
		wantErr   bool
	}{
		{
			name:      "Test1",
			wantInfos: nil,
			wantErr:   true,
		},
		//{ just possible to test for wantErr=false??

		// TODO: Add test cases.
	}
	if runtime.GOOS != "windows" && constants.LOCAL_2 {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotInfos, err := getWindowsInfo()
				if (err != nil) != tt.wantErr {
					t.Errorf("getWindowsInfo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotInfos, tt.wantInfos) {
					t.Errorf("getWindowsInfo() gotInfos = %v, want %v", gotInfos, tt.wantInfos)
				}
			})
		}
	}
}

func Test_getWindowsInfo_negativ_latest(t *testing.T) {
	tests := []struct {
		name      string
		wantInfos map[string]string
		wantErr   bool
	}{
		{
			name:      "Test1",
			wantInfos: nil,
			wantErr:   true,
		},
		//{ just possible to test for wantErr=false??

		// TODO: Add test cases.
	}
	if runtime.GOOS != "windows" && constants.RUNNER && constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotInfos, err := getWindowsInfo()
				if (err != nil) != tt.wantErr {
					t.Errorf("getWindowsInfo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotInfos, tt.wantInfos) {
					t.Errorf("getWindowsInfo() gotInfos = %v, want %v", gotInfos, tt.wantInfos)
				}
			})
		}
	}
}

func Test_getWindowsInfo_negativ(t *testing.T) {
	tests := []struct {
		name      string
		wantInfos map[string]string
		wantErr   bool
	}{
		{
			name:      "Test1",
			wantInfos: nil,
			wantErr:   true,
		},
		//{ just possible to test for wantErr=false??

		// TODO: Add test cases.
	}
	if runtime.GOOS != "windows" && constants.RUNNER && !constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotInfos, err := getWindowsInfo()
				if (err != nil) != tt.wantErr {
					t.Errorf("getWindowsInfo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotInfos, tt.wantInfos) {
					t.Errorf("getWindowsInfo() gotInfos = %v, want %v", gotInfos, tt.wantInfos)
				}
			})
		}
	}
}

func Test_windowsInfo_name_local1(t *testing.T) {
	type fields struct {
		infos map[string]string
	}
	tests := []struct {
		name     string
		fields   fields
		wantName string
		wantErr  bool
	}{
		{
			name:     "Test1",
			fields:   fields{map[string]string{"Caption": "Windows 10 Pro "}},
			wantName: "Windows 10 Pro ",
			wantErr:  false,
		},
		{
			name:     "Test2",
			fields:   fields{map[string]string{"Caption": "UNKNOWN_NAME"}},
			wantName: "UNKNOWN_NAME",
			wantErr:  false,
		},
		// just possible to test for wantErr=false??
		// TODO: Add test cases.
	}

	if runtime.GOOS == "windows" && constants.LOCAL_1 {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				w := windowsInfo{
					infos: tt.fields.infos,
				}
				gotName, err := w.name()
				if (err != nil) != tt.wantErr {
					t.Errorf("name() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if gotName != tt.wantName {
					t.Errorf("name() gotName = %v, want %v", gotName, tt.wantName)
				}
			})
		}
	}
}

func Test_windowsInfo_name_local2(t *testing.T) {
	type fields struct {
		infos map[string]string
	}
	tests := []struct {
		name     string
		fields   fields
		wantName string
		wantErr  bool
	}{
		{
			name:     "Test1",
			fields:   fields{map[string]string{"Caption": "Windows 10 Home "}},
			wantName: "Windows 10 Home ",
			wantErr:  false,
		},
		{
			name:     "Test2",
			fields:   fields{map[string]string{"Caption": "UNKNOWN_NAME"}},
			wantName: "UNKNOWN_NAME",
			wantErr:  false,
		},
		// just possible to test for wantErr=false??
		// TODO: Add test cases.
	}

	if runtime.GOOS == "windows" && constants.LOCAL_2 {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				w := windowsInfo{
					infos: tt.fields.infos,
				}
				gotName, err := w.name()
				if (err != nil) != tt.wantErr {
					t.Errorf("name() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if gotName != tt.wantName {
					t.Errorf("name() gotName = %v, want %v", gotName, tt.wantName)
				}
			})
		}
	}
}

func Test_windowsInfo_name_latest(t *testing.T) {
	type fields struct {
		infos map[string]string
	}
	tests := []struct {
		name     string
		fields   fields
		wantName string
		wantErr  bool
	}{
		{
			name:     "Test1",
			fields:   fields{map[string]string{"Caption": "Windows 10 Pro "}},
			wantName: "Windows 10 Pro ",
			wantErr:  false,
		},
		{
			name:     "Test2",
			fields:   fields{map[string]string{"Caption": "UNKNOWN_NAME"}},
			wantName: "UNKNOWN_NAME",
			wantErr:  false,
		},
		// just possible to test for wantErr=false??
		// TODO: Add test cases.
	}

	if runtime.GOOS == "windows" && constants.RUNNER && constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				w := windowsInfo{
					infos: tt.fields.infos,
				}
				gotName, err := w.name()
				if (err != nil) != tt.wantErr {
					t.Errorf("name() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if gotName != tt.wantName {
					t.Errorf("name() gotName = %v, want %v", gotName, tt.wantName)
				}
			})
		}
	}
}

func Test_windowsInfo_name(t *testing.T) {
	type fields struct {
		infos map[string]string
	}
	tests := []struct {
		name     string
		fields   fields
		wantName string
		wantErr  bool
	}{
		{
			name:     "Test1",
			fields:   fields{map[string]string{"Caption": "Windows 10 Pro "}},
			wantName: "Windows 10 Pro ",
			wantErr:  false,
		},
		{
			name:     "Test2",
			fields:   fields{map[string]string{"Caption": "UNKNOWN_NAME"}},
			wantName: "UNKNOWN_NAME",
			wantErr:  false,
		},
		// just possible to test for wantErr=false??
		// TODO: Add test cases.
	}

	if runtime.GOOS == "windows" && constants.RUNNER && !constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				w := windowsInfo{
					infos: tt.fields.infos,
				}
				gotName, err := w.name()
				if (err != nil) != tt.wantErr {
					t.Errorf("name() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if gotName != tt.wantName {
					t.Errorf("name() gotName = %v, want %v", gotName, tt.wantName)
				}
			})
		}
	}
}

func Test_windowsInfo_version_local1(t *testing.T) {
	type fields struct {
		infos map[string]string
	}
	tests := []struct {
		name        string
		fields      fields
		wantVersion string
		wantErr     bool
	}{
		{
			name:        "Test1",
			fields:      fields{map[string]string{"Version": "10.0.19041"}},
			wantVersion: "10.0.19041",
			wantErr:     false,
		},
		{
			name:        "Test2",
			fields:      fields{map[string]string{"Version": "UNKNOWN_VERSION"}},
			wantVersion: "UNKNOWN_VERSION",
			wantErr:     false,
		},
		{
			name:        "Test3",
			fields:      fields{nil},
			wantVersion: "10.0.19041 ",
			wantErr:     false,
		},
		// just possible to test for wantErr=false??
		// TODO: Add test cases.
	}
	if runtime.GOOS == "windows" && constants.LOCAL_1 {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				w := windowsInfo{
					infos: tt.fields.infos,
				}
				gotVersion, err := w.version()
				if (err != nil) != tt.wantErr {
					t.Errorf("version() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if gotVersion != tt.wantVersion {
					t.Errorf("version() gotVersion = %v, want %v", gotVersion, tt.wantVersion)
				}
			})
		}
	}

}

func Test_windowsInfo_version_local2(t *testing.T) {
	type fields struct {
		infos map[string]string
	}
	tests := []struct {
		name        string
		fields      fields
		wantVersion string
		wantErr     bool
	}{
		{
			name:        "Test1",
			fields:      fields{map[string]string{"Version": "10.0.19043"}},
			wantVersion: "10.0.19043",
			wantErr:     false,
		},
		{
			name:        "Test2",
			fields:      fields{map[string]string{"Version": "UNKNOWN_VERSION"}},
			wantVersion: "UNKNOWN_VERSION",
			wantErr:     false,
		},
		{
			name:        "Test3",
			fields:      fields{nil},
			wantVersion: "10.0.19043 ",
			wantErr:     false,
		},
		// just possible to test for wantErr=false??
		// TODO: Add test cases.
	}
	if runtime.GOOS == "windows" && constants.LOCAL_2 {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				w := windowsInfo{
					infos: tt.fields.infos,
				}
				gotVersion, err := w.version()
				if (err != nil) != tt.wantErr {
					t.Errorf("version() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if gotVersion != tt.wantVersion {
					t.Errorf("version() gotVersion = %v, want %v", gotVersion, tt.wantVersion)
				}
			})
		}
	}

}

func Test_windowsInfo_version_latest(t *testing.T) {
	type fields struct {
		infos map[string]string
	}
	tests := []struct {
		name        string
		fields      fields
		wantVersion string
		wantErr     bool
	}{
		{
			name:        "Test1",
			fields:      fields{map[string]string{"Version": "10.0.19041"}},
			wantVersion: "10.0.19041",
			wantErr:     false,
		},
		{
			name:        "Test2",
			fields:      fields{map[string]string{"Version": "UNKNOWN_VERSION"}},
			wantVersion: "UNKNOWN_VERSION",
			wantErr:     false,
		},
		{
			name:        "Test3",
			fields:      fields{nil},
			wantVersion: "10.0.17763 ",
			wantErr:     false,
		},
		// just possible to test for wantErr=false??
		// TODO: Add test cases.
	}
	if runtime.GOOS == "windows" && constants.RUNNER && constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				w := windowsInfo{
					infos: tt.fields.infos,
				}
				gotVersion, err := w.version()
				if (err != nil) != tt.wantErr {
					t.Errorf("version() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if gotVersion != tt.wantVersion {
					t.Errorf("version() gotVersion = %v, want %v", gotVersion, tt.wantVersion)
				}
			})
		}
	}

}

func Test_windowsInfo_version(t *testing.T) {
	type fields struct {
		infos map[string]string
	}
	tests := []struct {
		name        string
		fields      fields
		wantVersion string
		wantErr     bool
	}{
		{
			name:        "Test1",
			fields:      fields{map[string]string{"Version": "10.0.19041"}},
			wantVersion: "10.0.19041",
			wantErr:     false,
		},
		{
			name:        "Test2",
			fields:      fields{map[string]string{"Version": "UNKNOWN_VERSION"}},
			wantVersion: "UNKNOWN_VERSION",
			wantErr:     false,
		},
		{
			name:        "Test3",
			fields:      fields{nil},
			wantVersion: "10.0.14393 ",
			wantErr:     false,
		},
		// just possible to test for wantErr=false??
		// TODO: Add test cases.
	}
	if runtime.GOOS == "windows" && constants.RUNNER && !constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				w := windowsInfo{
					infos: tt.fields.infos,
				}
				gotVersion, err := w.version()
				if (err != nil) != tt.wantErr {
					t.Errorf("version() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if gotVersion != tt.wantVersion {
					t.Errorf("version() gotVersion = %v, want %v", gotVersion, tt.wantVersion)
				}
			})
		}
	}

}
