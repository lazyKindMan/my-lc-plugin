package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func nthUglyNumber(n int) int {
	p2, p3, p5 := 1, 1, 1
	product2, product3, product5 := 1, 1, 1
	p := 1
	ugly := make([]int, n+1)
	for p <= n {
		min := uglyMin(uglyMin(product2, product3), product5)
		ugly[p] = min
		p++
		if min == product2 {
			product2 = ugly[p2] * 2
			p2++
		}
		if min == product3 {
			product3 = ugly[p3] * 3
			p3++
		}
		if min == product5 {
			product5 = ugly[p5] * 5
			p5++
		}
	}
	return ugly[n]
}

func uglyMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func TestUglyNumberIi(t *testing.T) {
	// your test code here

}
