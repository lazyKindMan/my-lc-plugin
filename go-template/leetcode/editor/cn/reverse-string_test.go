package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func reverseString(s []byte) {
	n := len(s)
	left, right := 0, n-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseString(t *testing.T) {
	// your test code here

}
