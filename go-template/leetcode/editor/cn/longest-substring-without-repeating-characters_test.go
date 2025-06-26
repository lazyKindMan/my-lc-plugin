package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	left, right, res := 0, 0, 0
	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		for window[c] > 1 {
			d := s[left]
			left++
			if _, ok := window[d]; ok {
				window[d]--
			}
		}
		if right-left > res {
			res = right - left
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	// your test code here

}
