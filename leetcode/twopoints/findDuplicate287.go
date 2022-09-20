package twopoints

// 给定一个包含n + 1 个整数的数组nums ，其数字都在[1, n]范围内（包括 1 和 n），可知至少存在一个重复的整数。
//
// 假设 nums 只有 一个重复的整数 ，返回这个重复的数 。
//
// 你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。
// 示例 1：
//
// 输入：nums = [1,3,4,2,2]
// 输出：2
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/find-the-duplicate-number
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func findDuplicate(nums []int) int {
	slow, quick := 0, nums[0]
	for nums[slow] != nums[quick] {
		slow = nums[slow]
		quick = nums[nums[quick]]
	}
	quick = 0
	for nums[quick] != nums[slow] {
		quick = nums[quick]
		slow = nums[slow]
	}
	return nums[slow]
}

// 此道题和链表有环有点类似，如 [1,3,4,2,4] 可以看作是一个链表index 为listNode value，里面的值为next的index
// 0-->1 1-->3 2-->4 3-->2 4--->4
//0-->1-->3--->2--->4-->4-->4
