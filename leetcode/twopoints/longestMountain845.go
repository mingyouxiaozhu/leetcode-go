package twopoints

// 845. 数组中的最长山脉
// 把符合下列属性的数组 arr 称为 山脉数组 ：
//
// arr.length >= 3
// 存在下标 i（0 < i < arr.length - 1），满足
// arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
// arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
// 给出一个整数数组 arr，返回最长山脉子数组的长度。如果不存在山脉子数组，返回 0 。
//
// 示例 1：
//
// 输入：arr = [2,1,4,7,3,2,5]
// 输出：5
// 解释：最长的山脉子数组是 [1,4,7,3,2]，长度为 5。
// 示例 2：
//
// 输入：arr = [2,2,2]
// 输出：0
// 解释：不存在山脉子数组。
func longestMountain(arr []int) int {
	n := len(arr)
	ans := 0
	left := 0
	for left+2 < n {
		right := left + 1
		if arr[left] < arr[left+1] {
			for right+1 < n && arr[right] < arr[right+1] {
				right++
			}
			if right < n-1 && arr[right] > arr[right+1] {
				for right+1 < n && arr[right] > arr[right+1] {
					right++
				}
				if right-left+1 > ans {
					ans = right - left + 1
				}
			} else {
				right++
			}
		}
		left = right
	}
	return ans
}

// 开始的想法，从左边开始找到生序，此时down 降序必须为0，然后生序结束，找到降序，此时up>0
// 然后计算结果，但是会存在不通过的值，即，先875，884，239，731，723，685
// 因为计算的时候已经是第二次生序的起始节点，所以i--
// 解法很丑陋
// beat100 的解法感觉就是先找生序，找到后在找降序。简单有效，开始卡在如何找生序上，起始一个简答的循环就ok
func long(nums []int) int {
	up, down, res := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if i > 0 && (down == 0 && nums[i] > nums[i-1]) {
			up++
			continue
		}
		if i > 0 && (up > 0 && nums[i] < nums[i-1]) {
			down++
			continue
		}
		if up != 0 && down != 0 {
			if down+up+1 > res {
				res = down + up + 1
			}
			i--
		}
		down = 0
		up = 0

	}
	if up != 0 && down != 0 {
		if down+up+1 > res {
			res = down + up + 1
		}
	}
	return res
}
