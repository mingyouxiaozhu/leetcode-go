package twopoints

// 给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。# 代表退格字符。
//
// 注意：如果对空文本输入退格字符，文本继续为空。
//
// 示例 1：
//
// 输入：s = "ab#c", t = "ad#c"
// 输出：true
// 解释：s 和 t 都会变成 "ac"。
// 示例 2：
//
// 输入：s = "ab##", t = "c#d#"
// 输出：true
// 解释：s 和 t 都会变成 ""。
// 示例 3：
//
// 输入：s = "a#c", t = "b"
// 输出：false
// 解释：s 会变成 "c"，但 t 仍然是 "b"。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/backspace-string-compare
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func backspaceCompare(s string, t string) bool {
	sa, ta := []byte{}, []byte{}
	for _, v := range []byte(s) {
		if v != '#' {
			sa = append(sa, v)
		} else if len(sa) > 0 {
			sa = sa[0 : len(sa)-1]
		}
	}
	for _, v := range []byte(t) {
		if v != '#' {
			ta = append(ta, v)
		} else if len(ta) > 0 {
			ta = ta[0 : len(ta)-1]
		}
	}
	if len(sa) != len(ta) {
		return false
	}
	for i := range ta {
		if sa[i] != ta[i] {
			return false
		}
	}

	return true
}

// v2
func backspaceCompareV2(s string, t string) bool {
	c1, c2 := len(s)-1, len(t)-1
	skip1, skip2 := 0, 0
	for c1 >= 0 && c2 >= 0 {
		if s[c1] == '#' {
			c1--
			skip1++
			continue
		}
		if t[c2] == '#' {
			c2--
			skip2++
			continue
		}
		if skip1 > 0 {
			c1--
			skip1--
			continue
		}
		if skip2 > 0 {
			c2--
			skip2--
			continue
		}
		if s[c1] != t[c2] {
			return false
		}
		c1--
		c2--

	}

	for c1 >= 0 {
		if s[c1] == '#' {
			skip1++
		} else {
			skip1--
			if skip1 < 0 {
				return false
			}
		}
		c1--
	}
	for c2 >= 0 {
		if t[c2] == '#' {
			skip2++
		} else {
			skip2--
			if skip2 < 0 {
				return false
			}
		}
		c2--
	}

	return true

}
