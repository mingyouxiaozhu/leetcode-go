package twopoints

import "sort"

// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
//
// 你返回所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
// 示例 1：
//
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
// 解释：
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
// 不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
// 注意，输出的顺序和三元组的顺序并不重要。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/3sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 三个数相等，只能使用一个不动，另外两个指针来回查找的方式，需要去掉重复，采用当前数字计算完毕
// 跳过和当前数字相等的数
func threeSum(nums []int) [][]int {
	var result [][]int
	n := len(nums)
	sort.Ints(nums)
	// nums[i] <= 0 三数之和，第一个数大于0 ，那么后续就没必要继续比较了
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		left, right := i+1, n-1
		for left < right {
			if nums[left]+nums[right]+nums[i] == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				for left < n && nums[left] == nums[left-1] {
					left++
				}
				for right > 0 && nums[right-1] == nums[right] {
					right--
				}
			} else if nums[left]+nums[right]+nums[i] > 0 {
				right--
			} else if nums[left]+nums[right]+nums[i] < 0 {
				left++
			}
		}
		// 开始考虑到在计算当前的值以后开始对i 进行去重，其实该重复也只是在i=0 的时候不可以忽略
		for i < n-1 && nums[i+1] == nums[i] {
			i++
		}
	}
	return result
}
