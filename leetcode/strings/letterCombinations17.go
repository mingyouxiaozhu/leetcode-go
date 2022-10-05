package strings

// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
// 示例 1：
//
// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/letter-combinations-of-a-phone-number
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
var (
	letterMap = []string{
		" ",    //0
		"",     //1
		"abc",  //2
		"def",  //3
		"ghi",  //4
		"jkl",  //5
		"mno",  //6
		"pqrs", //7
		"tuv",  //8
		"wxyz", //9
	}
	res = []string{}
)

func letterCombinations(digits string) []string {
	if digits == "" {
		return res
	}
	var findCombination func(digits *string, index int, s string)
	findCombination = func(digits *string, index int, s string) {
		if index == len(*digits) {
			res = append(res, s)
			return
		}
		num := (*digits)[index]
		letter := letterMap[num-'0']
		for i := 0; i < len(letter); i++ {
			findCombination(digits, index+1, s+string(letter[i]))
		}
		return
	}
	findCombination(&digits, 0, "")
	return res
}

func letterCombinationsV1(digits string) []string {
	if digits == "" {
		return res
	}
	var helper func(digits string, index int, c string)
	helper = func(digits string, index int, c string) {
		if index == len(digits) {
			res = append(res, c)
			return
		}
		cur := letterMap[digits[index]-'0']
		for i := 0; i < len(cur); i++ {
			helper(digits, index+1, c+string(cur[i]))
		}
	}
	helper(digits, 0, "")
	return res
}
