package twopoints

import "sort"

// 给你两个整数数组nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，
// 应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。
//
// 示例 1：
//
// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
// 输出：[2,2]
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/intersection-of-two-arrays-ii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func intersect(nums1 []int, nums2 []int) []int {
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
		}
	}
	return res
}

func intersectV1(nums1 []int, nums2 []int) []int {
	dic, res := map[int]int{}, []int{}
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	for _, v := range nums1 {
		dic[v]++
	}
	for _, v := range nums2 {
		if dic[v] > 0 {
			res = append(res, v)
			dic[v]--
		}
	}
	return res

}
