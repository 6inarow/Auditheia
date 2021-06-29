package configParser

import (
	"testing"
)

func TestParseConfigFile(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "Test1", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ParseConfigFile(); (err != nil) != tt.wantErr {
				t.Errorf("ParseConfigFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_checkCorrectRootPermissions(t *testing.T) {

	tests := []struct {
		name    string
		want    bool
		wantErr bool
	}{
		{name: "Test1", want: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := checkCorrectRootPermissions()
			if (err != nil) != tt.wantErr {
				t.Errorf("checkCorrectRootPermissions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkCorrectRootPermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkCorrectRuntimeOS(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{name: "Test1", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkCorrectRuntimeOS(); got != tt.want {
				t.Errorf("checkCorrectRuntimeOS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseCLIFlags(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "Test1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParseCLIFlags()
		})
	}
}
