package binary

import "testing"

func Test_findTheDistanceValue1(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
		d    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: args{arr1: []int{1, 4, 2, 3}, arr2: []int{-4, -3, 6, 10, 20, 30}, d: 3}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDistanceValue1(tt.args.arr1, tt.args.arr2, tt.args.d); got != tt.want {
				t.Errorf("findTheDistanceValue1() = %v, want %v", got, tt.want)
			}
		})
	}
}
