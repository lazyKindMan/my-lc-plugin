package leetcode_solutions

import (
	"sort"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	var track []int
	sort.Ints(nums)
	traceback(nums, &track, &res, 0)
	return res
}

func traceback(nums []int, track *[]int, res *[][]int, start int) {
	*res = append(*res, append([]int(nil), *track...))
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		*track = append(*track, nums[i])
		traceback(nums, track, res, i+1)
		*track = (*track)[:len(*track)-1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSubsetsIi(t *testing.T) {
	// your test code here

}
