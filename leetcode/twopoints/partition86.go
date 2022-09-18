package twopoints

// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
//
// 你应当 保留 两个分区中每个节点的初始相对位置。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/partition-list
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
import . "leetcode-go/common"

func partition(head *ListNode, x int) *ListNode {
	mockS, mockM := &ListNode{Val: 0, Next: nil}, &ListNode{Val: 0, Next: nil}
	ts, tM := mockS, mockM
	t := head

	for t != nil {
		if t.Val < x {
			ts.Next = t
			ts = ts.Next
		} else {
			tM.Next = t
			tM = tM.Next
		}
		t = t.Next
	}
	ts.Next = mockM.Next
	tM.Next = nil
	return mockS.Next

}
