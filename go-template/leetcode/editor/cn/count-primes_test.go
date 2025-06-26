package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func countPrimes(n int) int {
	isPrimes := make([]int, n)
	for i := 0; i < n; i++ {
		isPrimes[i] = 1
	}
	ans := 0
	for i := 2; i < n; i++ {
		if isPrimes[i] == 1 {
			ans += 1
			if i*i < n {
				for j := i * i; j < n; j += i {
					isPrimes[j] = 0
				}
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCountPrimes(t *testing.T) {
	// your test code here

}
