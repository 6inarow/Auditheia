package memory

import (
	"reflect"
	"testing"
)

func TestNewArtefact(t *testing.T) {
	type args struct {
		name    string
		content string
	}
	tests := []struct {
		name string
		args args
		want *Artefact
	}{
		{ //not sure if Tests works, Test ends up false if wrong name but takes 0.00s
			name: "Test1",
			args: args{"Artefact1", "Hallo"},
			want: &Artefact{Name: "Artefact1", Content: "Hallo"},
		},
		{
			name: "Test2",
			args: args{"Artefact2", "Hallölle"},
			want: &Artefact{Name: "Artefact2", Content: "Hallölle"},
		},
		{
			name: "Test3",
			args: args{},
			want: &Artefact{},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArtefact(tt.args.name, tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArtefact() = %v, want %v", got, tt.want)
			}
		})
	}
}
