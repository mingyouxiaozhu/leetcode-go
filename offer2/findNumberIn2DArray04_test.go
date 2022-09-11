package offer2

import (
	"fmt"
	"testing"
)

func TestFindNumberIn2DArray(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	matrix1 := [][]int{{-1, 3}}
	fmt.Println(findNumberIn2DArray(matrix, 0))
	fmt.Println(findNumberIn2DArray(matrix1, 3))
}
