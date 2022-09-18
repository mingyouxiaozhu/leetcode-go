package twopoints

import . "leetcode-go/common"

// 给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
//
// 不允许修改 链表。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/linked-list-cycle-ii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 现在假设头节点到第一次相遇距离为N，第一次fast和slow 相遇距离环入口节点为S
// 此时，Slow 走了N+S 的距离
// 现在将fast 指向头指针，然后slow 和fast的速度一致前进，那么他们最后肯定会在第一次相遇（N+S）的地方再次遇到。
// 但是我们已经知道第一次相遇距离环入口的距离为S，所以当fast第二次进入环的时候，肯定也是走S的距离。
// 也就是说如果仍然需要在第一次相遇的地方再次相遇，那么fast进入环以后走的距离是S，slow 由于一直在环里走，所以当fast
// 进入环以后，他们肯定都是走S的距离，所以他们肯定有S的距离是重合着走的，所以第一次相遇的节点就是环入口。

func detectCycle(head *ListNode) *ListNode {
	h, fast, slow := HasCycle1(head)
	if !h {
		return nil
	} else {
		fast = head
		for fast != slow {
			fast = fast.Next
			slow = slow.Next
		}
	}
	return slow
}
