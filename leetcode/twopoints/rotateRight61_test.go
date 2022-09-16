package twopoints

import (
	"reflect"
	"testing"
)

import . "leetcode-go/common"

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test61", args: args{head: GenerateListFrom([]int{1, 2, 3, 4, 5}), k: 2}, want: GenerateListFrom([]int{4, 5, 1, 2, 3})},
		{name: "test61", args: args{head: GenerateListFrom([]int{1, 2, 3, 4, 5}), k: 2000}, want: GenerateListFrom([]int{4, 5, 1, 2, 3})},
		{name: "test61", args: args{head: GenerateListFrom([]int{1, 2, 3, 4, 5}), k: 5}, want: GenerateListFrom([]int{4, 5, 1, 2, 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
