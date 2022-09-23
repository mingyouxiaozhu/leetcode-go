package twopoints

import "testing"

func Test_findPairs(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: args{nums: []int{1, 3, 1, 5, 4}, k: 0}, want: 1},
		{name: "test", args: args{nums: []int{1, 2, 3, 4, 5}, k: 1}, want: 4},
		{name: "test", args: args{nums: []int{-1, -3, -1, -3}, k: 4}, want: 0},
		{name: "test", args: args{nums: []int{1, 1, 3, 2}, k: 2}, want: 1},
		{name: "test", args: args{nums: []int{1, 2, 4, 4, 3, 3, 0, 9, 2, 3}, k: 3}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPairsV1(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
