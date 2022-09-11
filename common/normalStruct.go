package common

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenerateListFrom(nums []int) *ListNode {
	head := &ListNode{Val: 0}
	tail := head
	for _, num := range nums {
		tail.Next = &ListNode{Val: num}
		tail = tail.Next
	}
	return head.Next
}
