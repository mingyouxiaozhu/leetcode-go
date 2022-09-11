package offer2

import "testing"

func Test_minArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: args{[]int{3, 4, 5, 1, 2, 2}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minArray(tt.args.arr); got != tt.want {
				t.Errorf("minArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerBound(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: args{[]int{1, 2, 3, 5, 5, 5, 6, 7, 9}, 5}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowerBound(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("LowerBound() = %v, want %v", got, tt.want)
			}
		})
	}
}
