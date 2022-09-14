package twopoints

import (
	. "leetcode-go/common"
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "testfor19", args: args{head: GenerateListFrom([]int{1, 2, 3, 4, 5}), n: 2}, want: GenerateListFrom([]int{1, 2, 3, 5})},
		{name: "testfor19", args: args{head: GenerateListFrom([]int{1, 2}), n: 2}, want: GenerateListFrom([]int{2})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEndV1(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
