package twopoints

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		str string
		t   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: args{str: "ADOBECODEBANC", t: "ABC"}, want: "BANC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.str, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
