package binary

import (
	"sort"
)

// 给你两个整数数组arr1，arr2和一个整数d，请你返回两个数组之间的距离值。
//
// 「距离值」定义为符合此距离要求的元素数目：对于元素arr1[i]，不存在任何元素arr2[j]满足 |arr1[i]-arr2[j]| <= d 。
//
// 示例 1：
//
// 输入：arr1 = [4,5,8], arr2 = [10,9,1,8], d = 2
// 输出：2
// 解释：
// 对于 arr1[0]=4 我们有：
// |4-10|=6 > d=2
// |4-9|=5 > d=2
// |4-1|=3 > d=2
// |4-8|=4 > d=2
// 所以 arr1[0]=4 符合距离要求
//
// 对于 arr1[1]=5 我们有：
// |5-10|=5 > d=2
// |5-9|=4 > d=2
// |5-1|=4 > d=2
// |5-8|=3 > d=2
// 所以 arr1[1]=5 也符合距离要求
//
// 对于 arr1[2]=8 我们有：
// |8-10|=2 <= d=2
// |8-9|=1 <= d=2
// |8-1|=7 > d=2
// |8-8|=0 <= d=2
// 存在距离小于等于 2 的情况，不符合距离要求
//
// 故而只有 arr1[0]=4 和 arr1[1]=5 两个符合距离要求，距离值为 2
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/find-the-distance-value-between-two-arrays
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
// bruce
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	b, n := 0, len(arr2)
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if abs(arr1[i], arr2[j]) > d {
				if j == n-1 {
					b++
				}
			} else {
				break
			}
		}

	}
	return b
}

// 找到一个值，先对arr2进行排序，之后遍历arr1，假设这个数为xx，因为 d >= 0，
// 如果arr2中存在一个数 yy 且满足 num−d≤y≤num+d，那么arr1中的数xx就不满足题目要求。
func findTheDistanceValue1(arr1 []int, arr2 []int, d int) int {
	res := 0
	sort.Ints(arr2)
	check := func(num int) bool {
		left, right := 0, len(arr2)
		for left < right {
			mid := left + (right-left)/2
			if abs(arr2[mid], num) > d {
				right = mid - 1
			} else {
				return false
			}
		}
		return true
	}

	for _, v := range arr1 {
		if check(v) {
			res++
		}
	}
	return res
}

func abs(i int, j int) int {
	tmp := i - j
	if tmp >= 0 {
		return tmp
	} else {
		return -tmp
	}
}
