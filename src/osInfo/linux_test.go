package osInfo

import (
	"Auditheia/memory/constants"
	"reflect"
	"runtime"
	"testing"
)

func Test_getLinuxInfo_latest(t *testing.T) {
	tests := []struct {
		name      string
		wantInfos map[string]string
		wantErr   bool
	}{
		{
			name: "Test1",
			wantInfos: map[string]string{"DISTRIB_CODENAME": "focal",
				"DISTRIB_DESCRIPTION": "Ubuntu 20.04.2 LTS", "DISTRIB_ID": "Ubuntu", "DISTRIB_RELEASE": "20.04"},
			wantErr: false,
		},

		// TODO: Add test cases.
	}
	if runtime.GOOS == "linux" && constants.RUNNER && constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotInfos, err := getLinuxInfo()
				if (err != nil) != tt.wantErr {
					t.Errorf("getLinuxInfo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotInfos, tt.wantInfos) {
					t.Errorf("getLinuxInfo() gotInfos = %v, want %v", gotInfos, tt.wantInfos)
				}
			})
		}
	}
}

func Test_getLinuxInfo(t *testing.T) {
	tests := []struct {
		name      string
		wantInfos map[string]string
		wantErr   bool
	}{
		{
			name: "Test1",
			wantInfos: map[string]string{"DISTRIB_CODENAME": "bionic",
				"DISTRIB_DESCRIPTION": "Ubuntu 18.04.5 LTS", "DISTRIB_ID": "Ubuntu", "DISTRIB_RELEASE": "18.04"},
			wantErr: false,
		},

		// TODO: Add test cases.
	}
	if runtime.GOOS == "linux" && constants.RUNNER && !constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotInfos, err := getLinuxInfo()
				if (err != nil) != tt.wantErr {
					t.Errorf("getLinuxInfo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotInfos, tt.wantInfos) {
					t.Errorf("getLinuxInfo() gotInfos = %v, want %v", gotInfos, tt.wantInfos)
				}
			})
		}
	}
}

func Test_linuxInfo_fullInfo_latest(t *testing.T) {
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
			fields:       fields{map[string]string{"DISTRIB_CODENAME": "focal", "DISTRIB_DESCRIPTION": "Ubuntu 20.04.2 LTS", "DISTRIB_ID": "Ubuntu", "DISTRIB_RELEASE": "20.04"}},
			wantCombined: "Ubuntu 20.04.2 LTS",
			wantErr:      false,
		},
		{
			name: "Test2",
			fields: fields{map[string]string{"noidetifier": "focal",
				"wrongversion": "Ubuntu 20.04.2 LTS", "idonotexist": "Ubuntu", "missingrelease": "20.04"}},
			wantCombined: "UNKNOWN DETAILS",
			wantErr:      true,
		},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "linux" && constants.RUNNER && constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				l := &linuxInfo{
					infos: tt.fields.infos,
				}
				gotCombined, err := l.fullInfo()
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

func Test_linuxInfo_fullInfo(t *testing.T) {
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
			name: "Test1",
			fields: fields{map[string]string{"DISTRIB_CODENAME": "bionic",
				"DISTRIB_DESCRIPTION": "Ubuntu 18.04.5 LTS", "DISTRIB_ID": "Ubuntu", "DISTRIB_RELEASE": "18.04"}},
			wantCombined: "Ubuntu 18.04.5 LTS",
			wantErr:      false,
		},
		{
			name: "Test2",
			fields: fields{map[string]string{"noidetifier": "bionic",
				"wrongversion": "Ubuntu 18.04.5 LTS", "idonotexist": "Ubuntu", "missingrelease": "18.04"}},
			wantCombined: "UNKNOWN DETAILS",
			wantErr:      true,
		},

		// TODO: Add test cases.
	}
	if runtime.GOOS == "linux" && constants.RUNNER && !constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				l := &linuxInfo{
					infos: tt.fields.infos,
				}
				gotCombined, err := l.fullInfo()
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

func Test_linuxInfo_name_latest(t *testing.T) {
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
			fields:   fields{map[string]string{"DISTRIB_CODENAME": "focal", "DISTRIB_DESCRIPTION": "Ubuntu 20.04.2 LTS", "DISTRIB_ID": "Ubuntu", "DISTRIB_RELEASE": "20.04"}},
			wantName: "Ubuntu",
			wantErr:  false,
		},
		{
			name: "Test2",
			fields: fields{map[string]string{"noidetifier": "focal",
				"wrongversion": "Ubuntu 20.04.2 LTS", "idonotexist": "Ubuntu", "missingrelease": "20.04"}},
			wantName: "UNKNOWN NAME",
			wantErr:  true,
		},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "linux" && constants.RUNNER && constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				l := &linuxInfo{
					infos: tt.fields.infos,
				}
				gotName, err := l.name()
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

func Test_linuxInfo_name(t *testing.T) {
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
			name: "Test1",
			fields: fields{map[string]string{"DISTRIB_CODENAME": "bionic",
				"DISTRIB_DESCRIPTION": "Ubuntu 18.04.5 LTS", "DISTRIB_ID": "Ubuntu", "DISTRIB_RELEASE": "18.04"}},
			wantName: "Ubuntu",
			wantErr:  false,
		},
		{
			name: "Test2",
			fields: fields{map[string]string{"noidetifier": "bionic",
				"wrongversion": "Ubuntu 18.04.5 LTS", "idonotexist": "Ubuntu", "missingrelease": "18.04"}},
			wantName: "UNKNOWN NAME",
			wantErr:  true,
		},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "linux" && constants.RUNNER && !constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				l := &linuxInfo{
					infos: tt.fields.infos,
				}
				gotName, err := l.name()
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

func Test_linuxInfo_version_latest(t *testing.T) {
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
			name: "Test1",
			fields: fields{map[string]string{"DISTRIB_CODENAME": "focal",
				"DISTRIB_DESCRIPTION": "Ubuntu 20.04.2 LTS", "DISTRIB_ID": "Ubuntu", "DISTRIB_RELEASE": "20.04"}},
			wantVersion: "20.04",
			wantErr:     false,
		},
		{
			name: "Test2",
			fields: fields{map[string]string{"noidetifier": "focal",
				"wrongversion": "Ubuntu 20.04.2 LTS", "idonotexist": "Ubuntu", "missingrelease": "20.04"}},
			wantVersion: "UNKNOWN VERSION",
			wantErr:     true,
		},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "linux" && constants.RUNNER && constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				l := &linuxInfo{
					infos: tt.fields.infos,
				}
				gotVersion, err := l.version()
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

func Test_linuxInfo_version(t *testing.T) {
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
			name: "Test1",
			fields: fields{map[string]string{"DISTRIB_CODENAME": "bionic",
				"DISTRIB_DESCRIPTION": "Ubuntu 18.04.5 LTS", "DISTRIB_ID": "Ubuntu", "DISTRIB_RELEASE": "18.04"}},
			wantVersion: "18.04",
			wantErr:     false,
		},
		{
			name: "Test2",
			fields: fields{map[string]string{"noidetifier": "bionic",
				"wrongversion": "Ubuntu 18.04.5 LTS", "idonotexist": "Ubuntu", "missingrelease": "18.04"}},
			wantVersion: "UNKNOWN VERSION",
			wantErr:     true,
		},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "linux" && constants.RUNNER && !constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				l := &linuxInfo{
					infos: tt.fields.infos,
				}
				gotVersion, err := l.version()
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
