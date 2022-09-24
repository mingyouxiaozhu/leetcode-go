package twopoints

// 给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
//
// 请你找出符合题意的 最短 子数组，并输出它的长度。
//
// 示例 1：
//
// 输入：nums = [2,6,4,8,10,9,15]
// 输出：5
// 解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
// 示例 2：
//
// 输入：nums = [1,2,3,4]
// 输出：0
// 示例 3：
//
// 输入：nums = [1]
// 输出：0
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/shortest-unsorted-continuous-subarray
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
// 将数据分成3组
// 左边为生序，右边为降序，中间为乱序
// 找到左边最大的数，该数肯定比右边有序的数最小数要小，但是比左边都大，此时为乱序的end
// 找到右边最小的数，该数肯定比左边有序的数最大的数要大，但是比右边都小，此时为乱序的start
func findUnsortedSubarray(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	min, max := nums[0], nums[len(nums)-1]
	start, end := 0, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] < max {
			end = i
		} else {
			max = nums[i]
		}
		if nums[len(nums)-i-1] > min {
			start = len(nums) - i - 1
		} else {
			min = nums[len(nums)-i-1]
		}
	}
	return end - start + 1
}
