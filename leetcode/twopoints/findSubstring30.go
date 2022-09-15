package twopoints

// 给定一个字符串s和一个字符串数组words。words中所有字符串 长度相同。
//
// s中的 串联子串 是指一个包含words中所有字符串以任意顺序排列连接起来的子串。
//
// 例如，如果words = ["ab","cd","ef"]， 那么"abcdef"，"abefcd"，"cdabef"，"cdefab"，"efabcd"， 和"efcdab" 都是串联子串。"acdbef" 不是串联子串
// 因为他不是任何words排列的连接。
// 返回所有串联字串在s中的开始索引。你可以以 任意顺序 返回答案。
//
// 示例 1：
//
// 输入：s = "barfoothefoobarman", words = ["foo","bar"]
// 输出：[0,9]
// 解释：因为 words.length == 2 同时 words[i].length == 3，连接的子字符串的长度必须为 6。
// 子串 "barfoo" 开始位置是 0。它是 words 中以 ["bar","foo"] 顺序排列的连接。
// 子串 "foobar" 开始位置是 9。它是 words 中以 ["foo","bar"] 顺序排列的连接。
// 输出顺序无关紧要。返回 [9,0] 也是可以的。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/substring-with-concatenation-of-all-words
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
// s 中必须由words拼接起来的，没次判断可以将数字切分为word的大小，然后判断是否相等
func findSubstring(s string, words []string) []int {
	res := []int{}
	wordCount, tmpCount := make(map[string]int), make(map[string]int)
	for _, w := range words {
		wordCount[w]++
	}
	sl, wsl, wl := len(s), len(words)*len(words[0]), len(words[0])
	for i := 0; i < sl-wsl+1; i++ {
		has, count := false, len(words)
		for j := 0; j < wsl; j += wl {
			left, right := i+j, i+j+wl
			if wordCount[s[left:right]] > 0 && tmpCount[s[left:right]] < wordCount[s[left:right]] {
				has = true
				tmpCount[s[left:right]]++
				count--
			} else {
				has = false
			}
		}
		tmpCount = make(map[string]int)
		if has && count == 0 {
			res = append(res, i)
		}

	}
	return res
}
