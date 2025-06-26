package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMoveZeroes(t *testing.T) {
	// your test code here

}
