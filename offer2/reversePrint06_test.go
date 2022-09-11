package offer2

import (
	"leetcode-go/common"
	"reflect"
	"testing"
)

func Test_reversePrint(t *testing.T) {
	type args struct {
		head *common.ListNode
	}
	var node1 = new(args)
	node1.head = new(common.ListNode)
	var node2 = new(common.ListNode)
	var node3 = new(common.ListNode)
	node1.head.Val = 10
	node2.Val = 9
	node3.Val = 8
	node1.head.Next = node2
	node2.Next = node3
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1", args: args{node1.head}, want: []int{8, 9, 10},
		},
		{
			name: "test2", args: args{nil}, want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrint(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrint() = %v, want %v", got, tt.want)
			}
		})
	}
}
