package binary

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 请必须使用时间复杂度为 O(log n) 的算法。
// 示例 1:
// 输入: nums = [1,3,5,6], target = 5
// 输出: 2
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/search-insert-position
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
/**
  本题为easy，主要想通过这道题了解到二分查找关于退出条件的计算和判断方式。
  本题很简单，即找到第一个不大于当前值的index
*/
func searchInsert(nums []int, target int) int {
	//1 使用朴素二分查找
	return binary1(nums, target)
}

// 使用的是左右闭合的区别，没次查找的范围为[left,right],退出条件为left=right+1
// 找到结果返回的mid，如果没找到，返回 0 （本题范围为0，其他可能为-1）
func binarySearchBasic(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target { // 如果当前的值比target 大，则说明需要的数据在mid的左边，mid 已经计算过，所以为mid-1
			right = mid - 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] >= target) {
				return mid + 1
			}
			left = mid + 1
			//如果比target 小，则说明在
		}
	}
	return 0
}

// 上面的在找到相等的值的时候会返回继续查找最左边的值，是否可以不做这种多余的操作
func binarySearch1(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid // 如果我们找到这个值，就使用right记录当前的值，后续查找是从[left,mid]
		} else {
			left = mid + 1
		}
	}
	if target > nums[left] {
		return left + 1
	}
	return len(nums)
}

func binary1(nums []int, target int) int {
	i, j := 0, len(nums)
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if nums[h] <= target {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}

//
