package twopoints

// 给你一个字符串 s 和一个整数 k 。你可以选择字符串中的任一字符，并将其更改为任何其他大写英文字符。该操作最多可执行 k 次。
//
// 在执行上述操作后，返回包含相同字母的最长子字符串的长度。
//
// 示例 1：
//
// 输入：s = "ABAB", k = 2
// 输出：4
// 解释：用两个'A'替换为两个'B',反之亦然。
// 示例 2：
//
// 输入：s = "AABABBA", k = 1
// 输出：4
// 解释：
// 将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
// 子串 "BBBB" 有最长重复字母, 答案为 4。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/longest-repeating-character-replacement
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
// 此种方法无法解决ABBB的问题
func characterReplacement(s string, k int) int {
	if k >= len(s) {
		return len(s)
	}
	bytes, res := []byte(s), 0
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for left := 0; left < len(s); left++ {
		right := left
		c := k
		for right < len(s) {
			if bytes[left] != bytes[right] {
				if c == 0 {
					c = k
					break
				} else {
					c--
				}
			}
			right++
		}
		res = max(res, right-left)
		if left > 0 && bytes[left] == bytes[left-1] {
			left++
		}
	}

	return res
}

// 从left 到right中，被替换的值是非当前最多的值
func characterReplacementV1(s string, k int) int {
	res, left, mc, bs := 0, 0, 0, []byte(s)
	count := map[byte]int{}
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for right, v := range bs {
		count[v]++
		mc = max(mc, count[v])
		for right-left+1-mc > k {
			count[bs[left]]--
			left++
		}
		res = max(res, right-left+1)
		right++
	}
	return res
}
