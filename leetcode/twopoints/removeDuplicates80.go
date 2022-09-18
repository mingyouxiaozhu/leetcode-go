package twopoints

// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
//
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

// 输入：nums = [1,1,1,2,2,3]
// 输出：5, nums = [1,1,2,2,3]
// 解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。 不需要考虑数组中超出新长度后面的元素。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
// 开始是自己想办法做个count计算个数，其实完全没必要啊，主要是有序,所以只需要判断当前的值是否和left-2 的值相等就可以，此时left为当前的index
func removeDuplicates(nums []int) int {
	left, right := 0, 0
	for ; right < len(nums); right++ {
		if right < 2 || nums[right] != nums[left-2] {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}
