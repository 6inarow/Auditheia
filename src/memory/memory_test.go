package memory

import (
	"reflect"
	"testing"
)

func TestCheckRegex(t *testing.T) {
	type args struct {
		regex string
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test1", args: args{"[a-zA-Z]\\w{3,7}", "AB_cd"}, want: true},
		{name: "Test2", args: args{"[a-zA-Z]\\w{3,7}", "00"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckRegex(tt.args.regex, tt.args.input); got != tt.want {
				t.Errorf("CheckRegex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCleanString(t *testing.T) {
	type args struct {
		input      string
		expression string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test1", args: args{"<javascript>Hallo ich bin ein String<javascript>", "<javascript>"}, want: "Hallo ich bin ein String"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CleanString(tt.args.input, tt.args.expression); got != tt.want {
				t.Errorf("CleanString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountAuditsAndTasks(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 int
	}{
		{
			name:  "Test1",
			want:  0,
			want1: 0,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CountAuditsAndTasks()
			if got != tt.want {
				t.Errorf("CountAuditsAndTasks() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CountAuditsAndTasks() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSeparateAtChar(t *testing.T) {
	type args struct {
		input          string
		c              rune
		preserveQuotes bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test1",
			args: args{"Halli_Hallo", '_', true},
			want: []string{"Halli", "Hallo"},
		},
		{
			name: "Test2",
			args: args{"Halli|Hallo|Hallöchen", '|', true},
			want: []string{"Halli", "Hallo", "Hallöchen"},
		},
		{ //different output with preserveQuotes false?? whats the difference
			name: "Test3",
			args: args{"Halli|Hallo|Hallöchen", '|', false},
			want: []string{"Halli", "Hallo", "Hallöchen"},
		},
		{
			name: "Test4",
			args: args{"Halli|Hallo|Hallöchen", 'ö', true},
			want: []string{"Halli|Hallo|Hall", "chen"},
		},
		{
			name: "Test5",
			args: args{"Halli|Hallo|Hal löchen", ' ', true},
			want: []string{"Halli|Hallo|Hal", "löchen"},
		},
		{
			name: "Test6",
			args: args{"I´m String with|'single quote'", '|', true},
			want: []string{"I´m String with", "'single quote'"},
		},
		{
			name: "Test7",
			args: args{"I´m String with-\"double quote\"", '-', true},
			want: []string{"I´m String with", "\"double quote\""},
		},
		{
			name: "Test8",
			args: args{"I´m String with|'single quote'", '|', false},
			want: []string{"I´m String with", "single quote"},
		},
		{
			name: "Test9",
			args: args{"I´m String with-\"double quote\"", '-', false},
			want: []string{"I´m String with", "double quote"},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SeparateAtChar(tt.args.input, tt.args.c, tt.args.preserveQuotes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SeparateAtChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countTasks(t *testing.T) {
	type args struct {
		tasks []Task
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{[]Task{Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, "")}},
			want: 1,
		},
		{
			name: "Test2",
			args: args{[]Task{Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, "")}},
			want: 5,
		},
		{
			name: "Test3",
			args: args{[]Task{Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""),
				Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""),
				Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""),
				Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, "")}},
			want: 25,
		},
		{
			name: "Test4",
			args: args{[]Task{Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""),
				Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""),
				Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""),
				Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""),
				Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""),
				Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""),
				Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""),
				Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""),
				Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""),
				Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""),
				Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""),
				Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""),
				Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""),
				Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""),
				Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""),
				Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, ""), Task2_Builder("script", "let result='SameSame';", []string{"SameSame"}, "equals", nil, nil, nil, "", "", "", "", true, "")}},
			want: 100,
		},
		{
			name: "Test5",
			args: args{[]Task{}},
			want: 0,
		},

		{ //task.onSuccess
			name: "Test6",
			args: args{[]Task{Task2_Builder("command", "", []string{"SameSame"}, "equals", TaskForTest_Builder("Command", "", []string{"Same"}, "equals", nil, nil, nil, "NotSameSame", "", "", "", true, ""), nil, nil, "SameSame", "", "", "", true, "")}},
			want: 2,
		},
		{ //task.onFail
			name: "Test7",
			args: args{[]Task{Task2_Builder("command", "", []string{"SameSame"}, "equals", nil, TaskForTest_Builder("Command", "", []string{"Same"}, "equals", nil, nil, nil, "NotSameSame", "", "", "", true, ""), nil, "SameSame", "", "", "", true, "")}},
			want: 2,
		},
		{ //task.onError expected auf nil
			name: "Test8",
			args: args{[]Task{Task2_Builder("command", "", []string{"SameSame"}, "equals", nil, nil, TaskForTest_Builder("Command", "", nil, "equals", nil, nil, nil, "", "", "", "", true, ""), "SameSame", "", "", "", true, "")}},
			want: 2,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTasks(tt.args.tasks...); got != tt.want {
				t.Errorf("countTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}
