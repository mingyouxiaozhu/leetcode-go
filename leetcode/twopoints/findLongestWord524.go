package twopoints

import "sort"

// 给你一个字符串 s 和一个字符串数组 dictionary ，找出并返回dictionary 中最长的字符串，该字符串可以通过删除 s 中的某些字符得到。
// 如果答案不止一个，返回长度最长且字母序最小的字符串。如果答案不存在，则返回空字符串。
// 示例 1：
//
// 输入：s = "abpcplea", dictionary = ["ale","apple","monkey","plea"]
// 输出："apple"
// 示例 2：
//
// 输入：s = "abpcplea", dictionary = ["a","b","c"]
// 输出："a"
//
// 提示：
//
// 1 <= s.length <= 1000
// 1 <= dictionary.length <= 1000
// 1 <= dictionary[i].length <= 1000
// s 和 dictionary[i] 仅由小写英文字母组成
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/longest-word-in-dictionary-through-deleting
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func findLongestWord(s string, dictionary []string) string {
	count := map[byte][]int{}
	for i, v := range []byte(s) {
		count[v] = append(count[v], i)
	}

	check := func(count map[byte][]int, ss string) bool {
		maxIndex := -1
		indexCount := map[byte]int{}
		for i, v := range []byte(ss) {
			if count[v] == nil {
				i++
				i--
				return false
			} else {
				if count[v][len(count[v])-1] < maxIndex || indexCount[v] > len(count[v]) {
					return false
				}
				// 如果已经没有使用的index ，比当前的max 还要大，则返回false，否则
				for ii := 0; ii <= len(count[v]); ii++ {
					if ii == len(count[v]) {
						return false
					}
					if count[v][ii] > maxIndex {
						indexCount[v]++
						maxIndex = count[v][ii]
						break
					}
				}
			}
		}
		return true
	}
	res := ""
	sort.Strings(dictionary)
	for _, v := range dictionary {
		if check(count, v) {
			if len(v) > len(res) {
				res = v
			}
		}
	}

	return res
}
