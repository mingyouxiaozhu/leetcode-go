package twopoints

// 给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
// 我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
//
// 必须在不使用库的sort函数的情况下解决这个问题。
//
// 示例 1：
//
// 输入：nums = [2,0,2,1,1,0]
// 输出：[0,0,1,1,2,2]
// 示例 2：
//
// 输入：nums = [2,0,1]
// 输出：[0,1,2]
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/sort-colors
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
// 交换法，因为0，1，2 一共就三种，如果等于0,则将该值放在slow，如果等于1，则cur++，如果等于2，则将这个值和quick交换，并且cur的值不变
// 快排其实可以解决这个问题
// 将数组遍历一边，然后得出0的个数，1的个数和2的个数，直接丢进去也是可以的
func sortColors(nums []int) {
	slow, quick, cur := 0, 0, len(nums)-1
	for cur < quick {
		if nums[cur] == 0 {
			nums[cur], nums[slow] = nums[slow], nums[cur]
			cur++
			slow++
		} else if nums[cur] == 2 {
			nums[cur], nums[quick] = nums[quick], nums[cur]
			//cur++ 如果将quick的值和cur 交换，则不需要cur++
			quick--
		} else {
			cur++
		}
	}
}
