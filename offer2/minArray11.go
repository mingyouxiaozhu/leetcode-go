package offer2

// 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
// 给你一个可能存在 重复 元素值的数组 numbers ，它原来是一个升序排列的数组，并按上述情形进行了一次旋转。请返回旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一次旋转，该数组的最小值为1。
// 示例 1：
// 输入：[3,4,5,1,2]
// 输出：1
// 示例 2：
// 输入：[2,2,2,0,1]
// 输出：0
// 注意：本题与主站 154 题相同：https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii/

func minArray(arr []int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	start, end := 0, n-1
	for start < end {
		mid := (start + end) / 2
		if arr[mid] > arr[end] { //此处开始是arr[mid] > arr[start],二分查找不要这么做
			start = mid + 1
		} else if arr[mid] < arr[end] {
			end = mid
		} else {
			end--
		}
	}
	return arr[start]
}

// 查找满足 x ≥ target 的下界的下标
func LowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] >= target { // 这里的比较运算符与题目要求一致
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left // 返回下界的下标
}
