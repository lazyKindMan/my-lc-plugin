package leetcode_solutions

import (
	"sort"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var trace []int
	var sum int
	sort.Ints(candidates)
	traceback1(candidates, &res, &trace, &sum, target, 0)
	return res
}

func traceback1(candidates []int, res *[][]int, trace *[]int, sum *int, target int, start int) {
	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		if *sum+candidates[i] <= target {
			*trace = append(*trace, candidates[i])
			*sum += candidates[i]
			if *sum == target {
				*res = append(*res, append([]int(nil), *trace...))
			} else {
				traceback1(candidates, res, trace, sum, target, i+1)
			}
			*trace = (*trace)[:len(*trace)-1]
			*sum -= candidates[i]
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCombinationSumIi(t *testing.T) {
	// your test code here
	var candidates = []int{10, 1, 2, 7, 6, 1, 5}
	combinationSum2(candidates, 8)
}
