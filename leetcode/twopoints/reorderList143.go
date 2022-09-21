package twopoints

//给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//
//L0 → L1 → … → Ln - 1 → Ln
//请将其重新排列后变为：
//
//L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
//不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/reorder-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
import . "leetcode-go/common"

func reorderList(head *ListNode) bool {
	quick, slow := head, head
	for quick != nil && quick.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next
	}
	mock := &ListNode{Val: -1, Next: nil}
	for slow != nil {
		tmp := slow.Next
		slow.Next = mock.Next
		mock.Next = slow
		slow = tmp
	}
	slow = head
	quick = mock.Next
	tail := mock
	for slow != nil && quick != nil && *slow != *quick {
		tail.Next = slow
		slow = slow.Next
		tail = tail.Next
		tail.Next = quick
		quick = quick.Next
		tail = tail.Next
	}
	if quick != nil {
		tail.Next = slow
	}
	if slow != nil {
		tail.Next = quick
	}

	return true
}
