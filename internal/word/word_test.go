package word

import "testing"

func TestToLower(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestToLower", args: args{s: "WANGCHENG"}, want: "wangcheng"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToLower(tt.args.s); got != tt.want {
				t.Errorf("ToLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUpper(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestToUpper", args: args{s: "wangcheng"}, want: "WANGCHENG"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUpper(tt.args.s); got != tt.want {
				t.Errorf("ToUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnderscoreToUpperCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestToUpper", args: args{s: "wang_cheng"}, want: "WangCheng"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnderscoreToUpperCamelCase(tt.args.s); got != tt.want {
				t.Errorf("UnderscoreToUpperCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnderscoreToLowerCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestToUpper", args: args{s: "wang_cheng"}, want: "wangCheng"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnderscoreToLowerCamelCase(tt.args.s); got != tt.want {
				t.Errorf("UnderscoreToLowerCameCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCamelCaseToUnderscore(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestToUpper", args: args{s: "wangCheng"}, want: "wang_cheng"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelCaseToUnderscore(tt.args.s); got != tt.want {
				t.Errorf("CamelCaseToUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}