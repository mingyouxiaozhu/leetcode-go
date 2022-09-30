package strings

// 编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串""。
//
// 示例 1：
//
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
// 示例 2：
//
// 输入：strs = ["dog","racecar","car"]
// 输出：""
// 解释：输入不存在公共前缀。
//
// 提示：
//
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 仅由小写英文字母组成
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/longest-common-prefix
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func longestCommonPrefix(strs []string) string {
	//left, maxLength := 0, math.MaxInt
	//res := ""
	//for _, v := range strs {
	//	if len(v) == 0 {
	//		return ""
	//	}
	//	if maxLength > len(v) {
	//		maxLength = len(v)
	//	}
	//}
	//c := strs[0][left]
	//for _, v := range strs {
	//	if left >= maxLength || v[left] != c {
	//		break
	//	} else {
	//		left++
	//		c = v[left]
	//	}
	//
	//}
	//if left > 0 {
	//	res = strs[0][0:left]
	//}
	//return res
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]

}
