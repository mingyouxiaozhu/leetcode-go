package leetcode

import (
	. "leetcode-go/common"
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		//[9,9,9,9,9,9,9]
		//[9,9,9,9]
		{name: "test2add", args: args{GenerateListFrom([]int{9, 9, 9, 9, 9, 9, 9}), GenerateListFrom([]int{9, 9, 9, 9})}, want: GenerateListFrom([]int{8, 9, 9, 9, 0, 0, 0, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
