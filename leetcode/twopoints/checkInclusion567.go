package twopoints

// 给你两个字符串s1和s2 ，写一个函数来判断 s2 是否包含 s1的排列。如果是，返回 true ；否则，返回 false 。
//
// 换句话说，s1 的排列之一是 s2 的 子串 。
//
// 示例 1：
//
// 输入：s1 = "ab" s2 = "eidbaooo"
// 输出：true
// 解释：s2 包含 s1 的排列之一 ("ba").
// 示例 2：
//
// 输入：s1= "ab" s2 = "eidboaoo"
// 输出：false
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/permutation-in-string
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var freq [256]int
	left, right, count := 0, 0, len(s1)
	for right < len(s2) {
		if freq[s2[right]-'a'] >= 1 {
			count--
		}
		freq[s2[right]-'a']--
		right++
		if count == 0 {
			return true
		}
		if right-left == len(s1) {
			if freq[s2[left]-'a'] >= 0 {
				count++
			}
			freq[s2[left]-'a']++
			left++
		}
	}
	return false
}

func checkInclusionV1(s1 string, s2 string) bool {
	if len(s1) == 0 || len(s1) > len(s2) {
		return false
	}
	var freq [26]int
	left, right, count := 0, 0, len(s1)
	for _, v := range []byte(s1) {
		freq[v-'a']++
	}

	for right < len(s2) {
		if freq[s2[right]-'a'] >= 1 {
			count--
		}
		freq[s2[right]-'a']--
		right++
		if count == 0 {
			return true
		}
		if right-left == len(s1) {
			if freq[s2[left]-'a'] >= 0 {
				count++
			}
			freq[s2[left]-'a']++
			left++
		}
	}
	return false
}
