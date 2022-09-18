package twopoints

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: args{[]int{1, 1, 1, 2, 2, 3}}, want: 5},
		{name: "test", args: args{[]int{1, 1, 2, 2, 3, 3, 3, 3, 4}}, want: 6},
		{name: "test", args: args{[]int{1, 1, 2, 2, 3, 3, 3, 3, 4}}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
