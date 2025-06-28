package leetcode_solutions

import "testing"

//leetcode submit region begin(Prohibit modification and deletion)
func numberOfSubstrings(s string) int {
	ans := 0
	cnt := []int{0, 0, 0}
	left := 0
	for _, c := range s {
		cnt[c - 'a'] ++
		for cnt[0] >0 && cnt[1] > 0 && cnt[2] > 0 {
			cnt[s[left] - 'a'] --
			left ++
		}
		ans += left
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNumberOfSubstringsContainingAllThreeCharacters(t *testing.T) {
	// your test code here

}
