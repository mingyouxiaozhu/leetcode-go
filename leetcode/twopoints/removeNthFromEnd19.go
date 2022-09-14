package twopoints

import . "leetcode-go/common"

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//
//
//示例 1：
//
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//示例 2：
//
//输入：head = [1], n = 1
//输出：[]
//示例 3：
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/remove-nth-node-from-end-of-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	quick, slow := head, head
	for i := 0; i < n; i++ {
		quick = quick.Next
	}

	for quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next
	}

	node := slow.Next //需要删除的节点
	if node == nil {
		return nil
	}
	if quick == nil {
		return slow.Next
	}
	slow.Next = node.Next
	return head
}

// 在遇到增加或者删除链表中的节点的时候，首先模拟一个虚拟节点出来
func removeNthFromEndV1(head *ListNode, n int) *ListNode {
	mock := &ListNode{Val: 0, Next: head}
	quick, slow := mock, mock
	for i := 0; i < n; i++ {
		quick = quick.Next
	}
	for quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next
	}
	slow.Next = slow.Next.Next
	return mock.Next
}
