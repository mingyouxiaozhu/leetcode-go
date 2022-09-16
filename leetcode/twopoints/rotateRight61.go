package twopoints

import . "leetcode-go/common"

// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动k个位置。
//
// 示例 1：
//
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[4,5,1,2,3]
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/rotate-list
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func rotateRight(head *ListNode, k int) *ListNode {
	length, quick, slow, mock := 0, head, head, &ListNode{Val: 0, Next: head}
	for quick != nil {
		quick = quick.Next
		length++
	}
	if k > length {
		k = k % length
		if k == 0 {
			return head
		}
	}
	quick = head
	for k > 0 {
		quick = quick.Next
		k--
	}
	for quick.Next != nil {
		quick = quick.Next
		slow = slow.Next
	}
	quick.Next = mock.Next
	mock.Next = slow.Next
	slow.Next = nil

	return mock.Next
}

// 使用一个指针
func rotateRight100(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	} else {
		n, iter := 1, head
		for ; iter.Next != nil; iter = iter.Next {
			n++
		}
		if k %= n; k == 0 {
			return head
		} else {
			iter.Next = head
			for ; k < n; k++ {
				iter = iter.Next
			}
			head, iter.Next = iter.Next, nil
			return head
		}
	}
}
