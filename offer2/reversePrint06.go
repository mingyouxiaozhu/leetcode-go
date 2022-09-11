package offer2

import "leetcode-go/common"

func reversePrint(head *common.ListNode) []int {
	if head == nil {
		return nil
	}
	result := []int{head.Val}

	for tail := head.Next; tail != nil; tail = tail.Next {
		result = append([]int{tail.Val}, result...)
	}
	return result
}
