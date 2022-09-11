package offer2

import "testing"
import "fmt"

func TestFindRepeatNumber(t *testing.T) {
	nums := []int{3, 4, 2, 1, 1, 0}
	a := findRepeatNumber(nums)
	fmt.Println(a)
}
