package leetcode_solutions

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	left, right := 0, n-1
	for left < right {
		if numbers[left]+numbers[right] == target {
			break
		} else if numbers[left]+numbers[right] < target {
			for left < right && numbers[left] == numbers[left+1] {
				left++
			}
			left++
		} else {
			for left < right && numbers[right] == numbers[right-1] {
				right--
			}
			right--
		}
	}
	return []int{left + 1, right + 1}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestTwoSumIiInputArrayIsSorted(t *testing.T) {

	// your test code here

}
