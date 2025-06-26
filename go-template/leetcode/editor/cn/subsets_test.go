package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	res = append(res, path)
	var choose func(i int)
	choose = func(i int) {
		path = append(path, nums[i])
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		for j := i + 1; j < len(nums); j++ {
			choose(j)
		}
		path = path[:len(path)-1]
	}
	for i := 0; i < len(nums); i++ {
		choose(i)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSubsets(t *testing.T) {
	// your test code here
	input := []int{1, 2, 3}
	println(subsets(input))
}
