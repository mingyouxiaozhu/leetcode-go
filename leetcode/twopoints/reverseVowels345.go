package twopoints

// 给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
//
// 元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现。
//
// 示例 1：
//
// 输入：s = "hello"
// 输出："holle"
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/reverse-vowels-of-a-string
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func reverseVowels(s string) string {
	left, right, bytes := 0, len(s)-1, []byte(s)

	dic := make(map[byte]bool)
	vo := "aeiouAEIOU"
	for _, v := range []byte(vo) {
		dic[v] = true
	}
	for left < right {
		if !dic[bytes[left]] {
			left++
		} else if !dic[bytes[right]] {
			right--
		} else {
			bytes[left], bytes[right] = bytes[right], bytes[left]
			left++
			right--
		}
	}
	return string(bytes)
}
