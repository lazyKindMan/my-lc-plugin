package leetcode_solutions

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func minimumRecolors(blocks string, k int) int {
	ans, white := 101, 0
	for i, block := range blocks {
		if block == 'W' {
			white++
		}
		if i < k-1 {
			continue
		}
		if ans > white {
			ans = white
		}
		if blocks[i-k+1] == 'W' {
			white--
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumRecolorsToGetKConsecutiveBlackBlocks(t *testing.T) {
	// your test code here

}
