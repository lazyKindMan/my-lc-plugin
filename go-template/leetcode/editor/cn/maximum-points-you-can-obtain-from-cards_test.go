package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	ans := 0
	s := 0
	for i := n - k; i < n+k; i++ {
		index := i % n
		s += cardPoints[index]
		if i < n - 1 {
			continue
		}
		if ans < s {
			ans = s
		}
		out := (i - k + 1) % n
		s -= cardPoints[out]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumPointsYouCanObtainFromCards(t *testing.T) {
	// your test code here

}
