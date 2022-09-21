package twopoints

import "testing"
import . "leetcode-go/common"

func Test_reorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test61", args: args{head: GenerateListFrom([]int{1, 2, 3, 4})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)
		})
	}
}
