package leetcode

import . "leetcode-go/common"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	need, res, l1Tail, l2Tail := false, &ListNode{Val: 1}, l1, l2
	resTail := res
	var currentValue int
	for l1Tail != nil && l2Tail != nil {
		currentValue = l1Tail.Val + l2Tail.Val
		if need {
			currentValue += 1
		}
		if currentValue > 9 {
			need = true
			resTail.Next = &ListNode{Val: currentValue % 10}
		} else {
			need = false
			resTail.Next = &ListNode{Val: currentValue}
		}
		resTail = resTail.Next
		l1Tail = l1Tail.Next
		l2Tail = l2Tail.Next
	}

	if l1Tail != nil {
		addOne(l1Tail, 0, resTail, need)
	} else {
		addOne(l2Tail, 0, resTail, need)
	}
	return res.Next
}

func addOne(l2Tail *ListNode, currentValue int, resTail *ListNode, need bool) {
	//resTail := res
	for l2Tail != nil {
		currentValue = l2Tail.Val
		if need {
			currentValue += 1
		}
		if currentValue > 9 {
			need = true
			resTail.Next = &ListNode{Val: currentValue % 10}
		} else {
			need = false
			resTail.Next = &ListNode{Val: currentValue}
		}
		resTail = resTail.Next
		l2Tail = l2Tail.Next
	}
	if need {
		resTail.Next = &ListNode{Val: 1}
	}
}
