package twopoints

import (
	"math"
	"sort"
)

// 给你一个长度为 n 的整数数组nums和 一个目标值target。请你从 nums 中选出三个整数，使它们的和与target最接近。
//
// 返回这三个数的和。
//
// 假定每组输入只存在恰好一个解。
// 示例 1：
//
// 输入：nums = [-1,2,1,-4], target = 1
// 输出：2
// 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
// 示例 2：
//
// 输入：nums = [0,0,0], target = 1
// 输出：0
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/3sum-closest
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 三数只和必须要考虑到距离为判断，和为结果，中间过程只需要判断距离，而不需要考虑到result
func threeSumClosest(nums []int, target int) int {
	result, n, sum := math.MaxInt, len(nums), math.MaxInt
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			tmp := nums[left] + nums[right] + nums[i]
			dist := abs(tmp, target)
			if dist < sum {
				sum = dist
				result = tmp
			}
			if tmp == target {
				return tmp
			} else if tmp > target {
				right--
			} else {
				left++
			}
		}
	}
	return result
}

func threeSumClosestV1(nums []int, target int) int {
	result, n, diff := math.MaxInt, len(nums), math.MaxInt
	sort.Ints(nums)
	update := func(d int, sum int) {
		if abs(d, 0) < diff {
			diff = d
			result = sum
		}
	}

	for i := 0; i < n; i++ {
		for j, k := i+1, n-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			update(abs(sum, target), sum)
			if sum == target {
				return sum
			}
			if sum > target {
				k--
				for k > j && nums[k] == nums[k+1] {
					k--
				}
			} else {
				j++
				for k > j && nums[j] == nums[j-1] {
					j++
				}
			}
		}
	}
	return result
}

func abs(i int, j int) int {
	res := i - j
	if res < 0 {
		return -res
	} else {
		return res
	}
}
