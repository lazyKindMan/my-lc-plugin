package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, s2 string) bool {
	need, window := make(map[rune]int), make(map[rune]int)
	for _, c := range s1 {
		need[c]++
	}
	valid, left, right := 0, 0, 0
	for right < len(s2) {
		c := rune(s2[right])
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if need[c] == window[c] {
				valid++
			}
		}
		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			d := rune(s2[left])
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPermutationInString(t *testing.T) {
	// your test code here
	s1 := "abcc"
	s2 := "abcbc"
	println(checkInclusion(s1, s2))
}
