package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	var backtrack func(int)
	backtrack = func(start int) {
		if k == len(cur) {
			res = append(res, append([]int(nil), cur...))
			return
		}
		for i := start; i <= n; i++ {
			cur = append(cur, i)
			backtrack(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(1)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCombinations(t *testing.T) {
	// your test code here

}
