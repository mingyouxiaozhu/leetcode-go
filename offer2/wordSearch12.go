package offer2

// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
// 例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。
// https://leetcode.cn/problems/word-search/
func wordSearch(board [][]byte, word string) bool {
	chars := []byte(word)
	//n := len(chars)
	m, n := len(board), len(board[0])
	used := make([][]bool, m)
	for i := 0; i < len(used); i++ {
		used[i] = make([]bool, len(board[0]))
	}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if board[row][col] == chars[0] {
				if find(row, col, board, used, chars, 0) {
					return true
				}
			}
		}
	}
	return false
}

// [["A","B","C","E"]
// ,["S","F","C","S"],
// ["A","D","E","E"]]
// "SEE"
func find(m, n int, board [][]byte, used [][]bool, chars []byte, index int) bool {
	if index == len(chars) {
		return true
	}
	res := false
	if m >= 0 && n >= 0 && m < len(board) && n < len(board[0]) && !used[m][n] {
		if board[m][n] == chars[index] {
			used[m][n] = true
			for i := 0; i < 4; i++ { //0 up 1 left 2 under 3 right
				switch i {
				case 0:
					res = find(m-1, n, board, used, chars, index+1)
					break
				case 1:
					res = find(m, n-1, board, used, chars, index+1)
					break
				case 2:
					res = find(m+1, n, board, used, chars, index+1)
					break
				case 3:
					res = find(m, n+1, board, used, chars, index+1)
					break
				}
				if res {
					return res
				}
			}
			if !res {
				used[m][n] = false
			}

		}
	}
	return res
}

type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	vis := make([][]bool, h)
	for i := range vis {
		vis[i] = make([]bool, w)
	}
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] { // 剪枝：当前字符不匹配
			return false
		}
		if k == len(word)-1 { // 单词存在于网格中
			return true
		}
		vis[i][j] = true
		defer func() { vis[i][j] = false }() // 回溯时还原已访问的单元格
		for _, dir := range directions {
			if newI, newJ := i+dir.x, j+dir.y; 0 <= newI && newI < h && 0 <= newJ && newJ < w && !vis[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}
