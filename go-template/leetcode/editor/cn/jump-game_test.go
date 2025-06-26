package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func canJump(nums []int) bool {
	farthest := 0
	for i := 0; i < len(nums)-1; i++ {
		tmp := i + nums[i]
		if tmp > farthest {
			farthest = tmp
		}
		if farthest <= i {
			return false
		}
	}
	return farthest >= len(nums)-1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestJumpGame(t *testing.T) {
	// your test code here

}
