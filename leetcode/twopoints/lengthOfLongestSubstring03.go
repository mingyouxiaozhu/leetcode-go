package twopoints

// 给定一个字符串 s ，请你找出其中不含有重复字符的最长子串的长度。
//
// 示例1:
//
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:
//
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:
//
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
//
//	请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。
//
// 0 <= s.length <= 5 * 104
// s 由英文字母、数字、符号和空格组成
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/longest-substring-without-repeating-characters
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func lengthOfLongestSubstring(s string) int {
	first, second, counter := 0, 0, [256]int{}
	charArray := []byte(s)
	res := 0
	for first < len(s) && second < len(s) {
		if counter[charArray[first]-'a'] == 0 {
			counter[charArray[first]-'a']++
			first++
		} else if second < first && counter[charArray[second]-'a'] == 1 {
			counter[charArray[second]-'a']--
			second++
		}
		res = max(res, first-second)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func lengthOfLongestSubstring0301(s string) int {

	freq := [256]int{}
	result, left, right := 0, 0, -1
	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}
