package twopoints

// 给你一个整数数组 nums 和一个整数 k ，请你返回子数组内所有元素的乘积严格小于 k 的连续子数组的数目。
//
// 示例 1：
//
// 输入：nums = [10,5,2,6], k = 100
// 输出：8
// 解释：8 个乘积小于 100 的子数组分别为：[10]、[5]、[2],、[6]、[10,5]、[5,2]、[2,6]、[5,2,6]。
// 需要注意的是 [10,5,2] 并不是乘积小于 100 的子数组。
// 示例 2：
//
// 输入：nums = [1,2,3], k = 0
// 输出：0
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/subarray-product-less-than-k
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
// 1 <= nums.length <= 3 * 104
// 1 <= nums[i] <= 1000
// 0 <= k <= 106
func numSubarrayProductLessThanK(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right, count, s := 0, 0, 0, 1
	for left < len(nums) {
		if right < len(nums) && s*nums[right] < k {
			s = s * nums[right]
			right++
		} else if left == right {
			left++
			right++
		} else {
			count += right - left
			s = s / nums[left]
			left++
		}
	}
	return count
}
func numSubarrayProductLessThanKV1(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	res, left, right, prod := 0, 0, 0, 1
	for left < len(nums) {
		if right < len(nums) && prod*nums[right] < k {
			prod = prod * nums[right]
			right++
		} else if left == right {
			left++
			right++
		} else {
			res += right - left
			prod = prod / nums[left]
			left++
		}
	}
	return res
}
