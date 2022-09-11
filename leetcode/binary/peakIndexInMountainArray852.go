package binary

/*
*
符合下列属性的数组 arr 称为 山脉数组 ：
arr.length >= 3
存在 i（0 < i< arr.length - 1）使得：
arr[0] < arr[1] < ... arr[i-1] < arr[i]
arr[i] > arr[i+1] > ... > arr[arr.length - 1]
给你由整数组成的山脉数组 arr ，返回任何满足 arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1] 的下标 i 。

示例 1：
输入：arr = [0,1,0]
输出：1

示例 2：
输入：arr = [0,2,1,0]
输出：1
示例 3：
输入：arr = [0,10,5,2]
输出：1
示例 4：
输入：arr = [3,4,5,1]
输出：2
示例 5：
输入：arr = [24,69,100,99,79,78,67,36,26,19]
输出：2
提示：
3 <= arr.length <= 104
0 <= arr[i] <= 106
题目数据保证 arr 是一个山脉数组
来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/peak-index-in-a-mountain-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	for left+1 < right {
		mid := left + (right-left)/2
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid] > arr[mid-1] && arr[mid] < arr[mid+1] {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}

func peakIndexInMountainArray100(arr []int) int {
	left, right := 0, len(arr)-1
	for left < right { // 因为题目必定有解，所有 left == right就是终止条件
		mid := int(uint(left+right) >> 1)
		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 暴力解法
func bruce(arr []int) int {
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			return i
		}
	}
	return 0
}
