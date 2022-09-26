package twopoints

import (
	"sort"
)

// 给定数组 people 。people[i]表示第 i 个人的体重 ，船的数量不限，每艘船可以承载的最大重量为 limit。
//
// 每艘船最多可同时载两人，但条件是这些人的重量之和最多为 limit。
//
// 返回 承载所有人所需的最小船数 。
//
// 示例 1：
//
// 输入：people = [1,2], limit = 3
// 输出：1
// 解释：1 艘船载 (1, 2)
// 示例 2：
//
// 输入：people = [3,2,2,1], limit = 3
// 输出：3
// 解释：3 艘船分别载 (1, 2), (2) 和 (3)
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/boats-to-save-people
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	res, l, r := 0, 0, len(people)-1
	for l+1 < r {
		res++
		if people[r] >= limit {
			r--
		} else if people[r]+people[l] > limit {
			r--
		} else if people[r]+people[l] <= limit {
			r--
			l++
		}
	}
	if r != l && people[l]+people[l+1] > limit {
		res += 2
	} else {
		res++
	}
	return res
}
