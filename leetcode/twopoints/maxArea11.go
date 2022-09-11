package twopoints

// 给定一个长度为 n 的整数数组height。有n条垂线，第 i 条线的两个端点是(i, 0)和(i, height[i])。
//
// 找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
//
// 返回容器可以储存的最大水量。
//
// 说明：你不能倾斜容器。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/container-with-most-water
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func maxArea(height []int) int {
	var min = func(a int, b int) int {
		if a > b {
			return b
		} else {
			return a
		}
	}
	var max = func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	result, left, right := 0, 0, len(height)-1
	for left < right {
		result = max(result, (right-left)*min(height[left], height[right]))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}
