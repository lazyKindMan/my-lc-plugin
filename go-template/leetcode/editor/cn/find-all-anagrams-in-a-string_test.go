package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	need, window := make(map[rune]int), make(map[rune]int)
	for _, c := range p {
		need[c]++
	}
	valid, left, right := 0, 0, 0
	res := make([]int, 0)
	for right < len(s) {
		c := rune(s[right])
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}
			d := rune(s[left])
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindAllAnagramsInAString(t *testing.T) {
	// your test code here

}
