package osInfo

import (
	"Auditheia/memory/constants"
	"reflect"
	"runtime"
	"testing"
)

func TestName_local1(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			want:    "Microsoft Windows 10 Pro ",
			wantErr: false,
		},
		// TODO: Add test cases.
	}

	if runtime.GOOS == "windows" && constants.LOCAL_1 {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := Name()
				if (err != nil) != tt.wantErr {
					t.Errorf("Name() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("Name() got = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func TestName_local2(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			want:    "Microsoft Windows 10 Home ",
			wantErr: false,
		},
		// TODO: Add test cases.
	}

	if runtime.GOOS == "windows" && constants.LOCAL_2 {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := Name()
				if (err != nil) != tt.wantErr {
					t.Errorf("Name() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("Name() got = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func TestName_windows_latest(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			want:    "Microsoft Windows Server 2019 Datacenter ",
			wantErr: false,
		},
		// TODO: Add test cases.
	}

	if runtime.GOOS == "windows" && constants.RUNNER && constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := Name()
				if (err != nil) != tt.wantErr {
					t.Errorf("Name() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("Name() got = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func TestName_windows(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			want:    "Microsoft Windows Server 2016 Datacenter ",
			wantErr: false,
		},
		// TODO: Add test cases.
	}

	if runtime.GOOS == "windows" && constants.RUNNER && !constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := Name()
				if (err != nil) != tt.wantErr {
					t.Errorf("Name() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("Name() got = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func TestName_darwin(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			want:    "Mac OS X",
			wantErr: false,
		},
		// TODO: Add test cases.
	}

	if runtime.GOOS == "darwin" && constants.RUNNER {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := Name()
				if (err != nil) != tt.wantErr {
					t.Errorf("Name() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("Name() got = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func TestName_linux_latest(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			want:    "Ubuntu",
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "linux" && constants.RUNNER && constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := Name()
				if (err != nil) != tt.wantErr {
					t.Errorf("Name() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("Name() got = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func TestName_linux(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			want:    "Ubuntu",
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "linux" && constants.RUNNER && !constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := Name()
				if (err != nil) != tt.wantErr {
					t.Errorf("Name() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("Name() got = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func TestVersion_local1(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			want:    "10.0.19041 ",
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "windows" && constants.LOCAL_1 {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := Version()
				if (err != nil) != tt.wantErr {
					t.Errorf("Version() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("Version() got = %v, want %v", got, tt.want)
				}
			})
		}
	}

}

func TestVersion_local2(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			want:    "10.0.19043 ",
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "windows" && constants.LOCAL_2 {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := Version()
				if (err != nil) != tt.wantErr {
					t.Errorf("Version() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("Version() got = %v, want %v", got, tt.want)
				}
			})
		}
	}

}

func TestVersion_windows_latest(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			want:    "10.0.17763 ",
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "windows" && constants.RUNNER && constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := Version()
				if (err != nil) != tt.wantErr {
					t.Errorf("Version() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("Version() got = %v, want %v", got, tt.want)
				}
			})
		}
	}

}

func TestVersion_windows(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			want:    "10.0.14393 ",
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "windows" && constants.RUNNER && !constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := Version()
				if (err != nil) != tt.wantErr {
					t.Errorf("Version() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("Version() got = %v, want %v", got, tt.want)
				}
			})
		}
	}

}

func TestVersion_linux_latest(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			want:    "20.04",
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "linux" && constants.RUNNER && constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := Version()
				if (err != nil) != tt.wantErr {
					t.Errorf("Version() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("Version() got = %v, want %v", got, tt.want)
				}
			})
		}
	}

}

func TestVersion_linux(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			want:    "18.04",
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "linux" && constants.RUNNER && !constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := Version()
				if (err != nil) != tt.wantErr {
					t.Errorf("Version() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("Version() got = %v, want %v", got, tt.want)
				}
			})
		}
	}

}

func TestVersion_darwin(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			want:    "10.15.7",
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "darwin" && constants.RUNNER {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := Version()
				if (err != nil) != tt.wantErr {
					t.Errorf("Version() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("Version() got = %v, want %v", got, tt.want)
				}
			})
		}
	}

}

func Test_keyValueParser(t *testing.T) {
	type args struct {
		data              string
		pairSeparator     string
		keyValueSeparator string
		cf1               stringCleanerFunc
		cf2               stringCleanerFunc
	}
	tests := []struct {
		name    string
		args    args
		wantOut map[string]string
		wantErr bool
	}{
		{
			name:    "Test1",
			args:    args{},
			wantOut: map[string]string{},
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := keyValueParser(tt.args.data, tt.args.pairSeparator, tt.args.keyValueSeparator, tt.args.cf1, tt.args.cf2)
			if (err != nil) != tt.wantErr {
				t.Errorf("keyValueParser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("keyValueParser() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

//func Test_newInfo(t *testing.T) {
//	tests := []struct {
//		name string
//		want info
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := newInfo(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("newInfo() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_unsupported_name(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			want:    "UNSUPPORTED OS",
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := unsupported{}
			got, err := d.name()
			if (err != nil) != tt.wantErr {
				t.Errorf("name() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("name() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unsupported_version(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			want:    "UNSUPPORTED OS",
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := unsupported{}
			got, err := d.version()
			if (err != nil) != tt.wantErr {
				t.Errorf("version() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("version() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFullInfo_windows_latest(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			want:    "UNKNOWN DETAILS",
			wantErr: true,
		},
		//{ Which output should TestFullInfo have so wantErr can be false?
		//	name:    "Test2",
		//	want:    "Caption: Microsoft Windows 10 Home  Version: 10.0.19043 ",
		//	wantErr: false,
		//},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "windows" && constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := FullInfo()
				if (err != nil) != tt.wantErr {
					t.Errorf("FullInfo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("FullInfo() got = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func TestFullInfo_windows(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			want:    "UNKNOWN DETAILS",
			wantErr: true,
		},
		//{ Which output should TestFullInfo have so wantErr can be false?
		//	name:    "Test2",
		//	want:    "Caption: Microsoft Windows 10 Home  Version: 10.0.19043 ",
		//	wantErr: false,
		//},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "windows" && !constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := FullInfo()
				if (err != nil) != tt.wantErr {
					t.Errorf("FullInfo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("FullInfo() got = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func TestFullInfo_linux_latest(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			want:    "Ubuntu 20.04.2 LTS",
			wantErr: false,
		},
		//{ Which output should TestFullInfo have so wantErr can be false?
		//	name:    "Test2",
		//	want:    "Caption: Microsoft Windows 10 Home  Version: 10.0.19043 ",
		//	wantErr: false,
		//},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "linux" && constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := FullInfo()
				if (err != nil) != tt.wantErr {
					t.Errorf("FullInfo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("FullInfo() got = %v, want %v", got, tt.want)
				}
			})
		}
	}

}

func TestFullInfo_linux(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			want:    "Ubuntu 18.04.5 LTS",
			wantErr: false,
		},
		//{ Which output should TestFullInfo have so wantErr can be false?
		//	name:    "Test2",
		//	want:    "Caption: Microsoft Windows 10 Home  Version: 10.0.19043 ",
		//	wantErr: false,
		//},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "linux" && !constants.LATEST_VERSION {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := FullInfo()
				if (err != nil) != tt.wantErr {
					t.Errorf("FullInfo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("FullInfo() got = %v, want %v", got, tt.want)
				}
			})
		}
	}

}

func TestFullInfo_macOS(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			want:    "UNKNOWN DETAILS",
			wantErr: true,
		},
		//{ Which output should TestFullInfo have so wantErr can be false?
		//	name:    "Test2",
		//	want:    "Caption: Microsoft Windows 10 Home  Version: 10.0.19043 ",
		//	wantErr: false,
		//},
		// TODO: Add test cases.
	}
	if runtime.GOOS == "darwin" {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := FullInfo()
				if (err != nil) != tt.wantErr {
					t.Errorf("FullInfo() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("FullInfo() got = %v, want %v", got, tt.want)
				}
			})
		}
	}

}

func Test_unsupported_fullInfo(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			want:    "UNSUPPORTED OS",
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := unsupported{}
			got, err := d.fullInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("fullInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fullInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
