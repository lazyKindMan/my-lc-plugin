package leetcode_solutions

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func jump(nums []int) int {
	n := len(nums)
	end, farthest, minStep := 0, 0, 0
	for i := 0; i < n-1; i++ {
		farthest := max(i+nums[i], farthest)
		if end == i {
			minStep++
			end = farthest
		}
		if farthest >= n-1 {
			break
		}
	}
	return minStep
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func TestJumpGameIi(t *testing.T) {

	// your test code here
	fmt.Println(jump([]int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}))
}
