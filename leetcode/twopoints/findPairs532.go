package twopoints

import (
	"sort"
)

// 给你一个整数数组nums 和一个整数k，请你在数组中找出 不同的k-diff 数对，并返回不同的 k-diff 数对 的数目。
//
// k-diff数对定义为一个整数对 (nums[i], nums[j]) ，并满足下述全部条件：
//
// 0 <= i, j < nums.length
// i != j
// nums[i] - nums[j] == k
// 注意，|val| 表示 val 的绝对值。
//
// 示例 1：
//
// 输入：nums = [3, 1, 4, 1, 5], k = 2
// 输出：2
// 解释：数组中有两个 2-diff 数对, (1, 3) 和 (3, 5)。
// 尽管数组中有两个 1 ，但我们只应返回不同的数对的数量。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/k-diff-pairs-in-an-array
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func findPairs(nums []int, k int) int {
	sort.Ints(nums)
	check := func(a int, b int, k int) bool {
		if a-b == k {
			return true
		}
		if a-b == -k {
			return true
		}
		return false
	}
	res := 0
	for left := 0; left < len(nums); left++ {
		for right := left + 1; right < len(nums); right++ {
			for right < len(nums)-1 && nums[right] == nums[right+1] {
				right++
			}
			if check(nums[left], nums[right], k) {
				res++
			}

		}
		for left < len(nums)-1 && nums[left] == nums[left+1] {
			left++
		}
	}
	return res
}

func findPairsV1(nums []int, k int) int {
	if k < 0 || len(nums) == 0 {
		return 0
	}
	var count int
	m := make(map[int]int, len(nums))
	for _, value := range nums {
		m[value]++
	}
	for key := range m {
		if k == 0 && m[key] > 1 {
			count++
			continue
		}
		if k > 0 && m[key+k] > 0 {
			count++
		}
	}
	return count
}
func findPairsV2(nums []int, k int) (ans int) {
	sort.Ints(nums)
	y, n := 0, len(nums)
	for x, num := range nums {
		if x == 0 || num != nums[x-1] {
			for y < n && (nums[y] < num+k || y <= x) {
				y++
			}
			if y < n && nums[y] == num+k {
				ans++
			}
		}
	}
	return
}
