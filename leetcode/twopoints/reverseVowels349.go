package twopoints

import "sort"

// 给定两个数组nums1和nums2 ，返回 它们的交集。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
//
// 示例 1：
//
// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
// 输出：[2]
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/intersection-of-two-arrays
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	res := []int{}
	for left, right := 0, 0; right < len(nums2) && left < len(nums1); {
		if nums1[left] < nums2[right] {
			left++
		} else if nums1[left] > nums2[right] {
			right++
		} else {
			res = append(res, nums1[left])
			left++
			right++
			for ; nums1[left] == nums1[left-1]; left++ {

			}
			for ; nums2[right] == nums2[right-1]; right++ {

			}
		}
	}
	return res
}
