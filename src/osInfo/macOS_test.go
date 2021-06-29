package osInfo

import (
	"Auditheia/memory/constants"
	"reflect"
	"runtime"
	"testing"
)

func Test_getMacOSInfo(t *testing.T) {
	tests := []struct {
		name      string
		wantInfos map[string]string
		wantErr   bool
	}{
		{
			name:      "Test1",
			wantInfos: map[string]string{"BuildVersion": "19H1217", "ProductName": "Mac OS X", "ProductVersion": "10.15.7"},
			wantErr:   false,
		},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "darwin" && constants.RUNNER {

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotInfos, err := getMacOSInfo()
				if (err != nil) != tt.wantErr {
					t.Errorf("getMacOSInfo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotInfos, tt.wantInfos) {
					t.Errorf("getMacOSInfo() gotInfos = %v, want %v", gotInfos, tt.wantInfos)
				}
			})
		}
	}
}

func Test_macOsInfo_fullInfo(t *testing.T) {
	type fields struct {
		infos map[string]string
	}
	tests := []struct {
		name         string
		fields       fields
		wantCombined string
		wantErr      bool
	}{
		{
			name:         "Test1",
			fields:       fields{map[string]string{"BuildVersion": "19H1217", "ProductName": "Mac OS X", "ProductVersion": "10.15.7"}},
			wantCombined: "UNKNOWN DETAILS",
			wantErr:      true,
		},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "darwin" && constants.RUNNER {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				m := &macOsInfo{
					infos: tt.fields.infos,
				}
				gotCombined, err := m.fullInfo()
				if (err != nil) != tt.wantErr {
					t.Errorf("fullInfo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if gotCombined != tt.wantCombined {
					t.Errorf("fullInfo() gotCombined = %v, want %v", gotCombined, tt.wantCombined)
				}
			})
		}
	}
}

func Test_macOsInfo_name(t *testing.T) {
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
			fields:   fields{map[string]string{"ProductName": "MacOS", "ProductVersion": "10.15.7"}},
			wantName: "MacOS",
			wantErr:  false,
		},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "darwin" && constants.RUNNER {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				m := &macOsInfo{
					infos: tt.fields.infos,
				}
				gotName, err := m.name()
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

func Test_macOsInfo_version(t *testing.T) {
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
			fields:   fields{map[string]string{"ProductName": "MacOS", "ProductVersion": "10.15.7"}},
			wantName: "10.15.7",
			wantErr:  false,
		},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "darwin" && constants.RUNNER {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				m := &macOsInfo{
					infos: tt.fields.infos,
				}
				gotName, err := m.version()
				if (err != nil) != tt.wantErr {
					t.Errorf("version() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if gotName != tt.wantName {
					t.Errorf("version() gotName = %v, want %v", gotName, tt.wantName)
				}
			})
		}
	}
}
