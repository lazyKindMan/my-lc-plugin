package leetcode_solutions

import (
	"math"
	"testing"
)

//leetcode submit region begin(Prohibit modification and deletion)
func findMaxAverage(nums []int, k int) float64 {
	ans := math.MinInt
	sum := 0
	for i, num := range nums {
		sum += num
		if i < k - 1 {
			continue
		}
		if sum > ans {
			ans = sum
		}
		sum -= nums[i - k + 1]
	}
	return float64(ans) / float64(k)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumAverageSubarrayI(t *testing.T) {
	// your test code here

}
