package strings

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "tes", args: args{[]string{"", "a"}}, want: ""},
		{name: "tes", args: args{[]string{"ab", "a"}}, want: "a"},
		{name: "tes", args: args{[]string{"flower", "flower", "flower", "flower"}}, want: "flower"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
