package output

import "testing"

func TestInitLog(t *testing.T) {
	tests := []struct {
		name string

		wantErr bool
	}{
		{name: "Test1", wantErr: false},
		//{name: "Test2", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitLog(); (err != nil) != tt.wantErr {
				t.Errorf("InitLog() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
