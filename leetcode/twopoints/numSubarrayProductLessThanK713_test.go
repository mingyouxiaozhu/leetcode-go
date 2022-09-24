package twopoints

import "testing"

func Test_numSubarrayProductLessThanK(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: args{
			nums: []int{10, 5, 2, 6},
			k:    100,
		}, want: 8},
		{name: "test", args: args{
			nums: []int{10, 5, 2, 6},
			k:    0,
		}, want: 0},
		{name: "test", args: args{
			nums: []int{10, 5, 2, 0},
			k:    0,
		}, want: 0},
		{name: "test", args: args{
			nums: []int{10, 9, 10, 4, 3, 8, 3, 3, 6, 2, 10, 10, 9, 3},
			k:    19,
		}, want: 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarrayProductLessThanK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("numSubarrayProductLessThanK() = %v, want %v", got, tt.want)
			}
		})
	}
}
