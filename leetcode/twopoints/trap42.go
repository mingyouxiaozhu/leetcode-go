package twopoints

// 给定n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
// 示例 1：
//
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
// 示例 2：
//
// 输入：height = [4,2,0,3,2,5]
// 输出：9
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/trapping-rain-water
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func trap(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))
	dp := make([]int, len(height))
	tmp := 0
	for i := 0; i < len(height); i++ {
		tmp = max(tmp, height[i])
		left[i] = tmp
	}
	tmp = 0
	for i := len(height) - 1; i >= 0; i-- {
		tmp = max(tmp, height[i])
		right[i] = tmp
	}
	res := 0
	for i := 1; i < len(height)-1; i++ {
		dp[i] = min(left[i-1], right[i+1]) - height[i]
		if dp[i] > 0 {
			res += dp[i]
		}
	}
	return res
}
