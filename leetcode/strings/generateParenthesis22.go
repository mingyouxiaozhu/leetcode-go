package strings

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
// 示例 1：
//
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
// 示例 2：
//
// 输入：n = 1
// 输出：["()"]
//
// 提示：
//
// 1 <= n <= 8
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/generate-parentheses
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := []string{}
	var findGenerateParenthesis func(lindex, rindex int, str string, res *[]string)
	findGenerateParenthesis = func(lindex, rindex int, str string, res *[]string) {
		if lindex == 0 && rindex == 0 {
			*res = append(*res, str)
			return
		}
		if lindex > 0 {
			findGenerateParenthesis(lindex-1, rindex, str+"(", res)
		}
		if rindex > 0 && lindex < rindex {
			findGenerateParenthesis(lindex, rindex-1, str+")", res)
		}
	}
	findGenerateParenthesis(n, n, "", &res)
	return res
}
