package twopoints

import "sort"

// 给你一个由 n 个整数组成的数组nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组[nums[a], nums[b], nums[c], nums[d]]（若两个四元组元素一一对应，则认为两个四元组重复）：
//
// 0 <= a, b, c, d< n
// a、b、c 和 d 互不相同
// nums[a] + nums[b] + nums[c] + nums[d] == target
// 你可以按 任意顺序 返回答案 。
//
// 示例 1：
//
// 输入：nums = [1,0,-1,0,-2,2], target = 0
// 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
// 示例 2：
//
// 输入：nums = [2,2,2,2,2], target = 8
// 输出：[[2,2,2,2]]
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/4sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result, n := [][]int{}, len(nums)
	for i := 0; i < n-3; i++ {
		for j := i + 1; j < n-2; j++ {

			left, right := j+1, n-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					for left < right && left > 0 && nums[left] == nums[left-1] {
						left++
					}
					for left < right && right < n-1 && nums[right] == nums[right+1] {
						right--
					}
				} else if sum > target {
					right--
				} else {
					left++
				}

			}
			for j < n-1 && nums[j+1] == nums[j] {
				j++
			}
		}
		for i < n-1 && nums[i+1] == nums[i] {
			i++
		}
	}
	return result
}
