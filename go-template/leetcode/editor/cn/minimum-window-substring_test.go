package leetcode_solutions

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	m := make(map[byte]int)
	total_contain_left := 0
	for i := 0; i < len(t); i++ {
		key := t[i]
		val, ok := m[key]
		if ok {
			m[key] = val + 1
		} else {
			m[key] = 1
		}
		total_contain_left++
	}
	left, right := 0, 0
	res := ""
	resLength := len(s) + 1
	for right < len(s) {
		c := s[right]
		right++
		if val, ok := m[c]; ok {
			if val > 0 {
				total_contain_left--
			}
			m[c] = val - 1
		}
		for total_contain_left == 0 {
			if right-left < resLength {
				res = s[left:right]
				resLength = right - left
			}
			c1 := s[left]
			left++
			if valLeft, okLeft := m[c1]; okLeft {
				if valLeft >= 0 {
					total_contain_left++
				}
				m[c1] = valLeft + 1
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumWindowSubstring(t *testing.T) {
	// your test code here
	s := "a"
	t1 := "a"
	println(minWindow(s, t1))
}
