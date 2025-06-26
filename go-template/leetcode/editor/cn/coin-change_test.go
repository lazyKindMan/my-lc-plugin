package leetcode_solutions

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)

//func coinChange(coins []int, amount int) int {
//	memo := make([]int, amount+1)
//	for i := range memo {
//		memo[i] = -666
//	}
//	return dp(coins, amount, memo)
//}
//
//func dp(coins []int, amount int, memo []int) int {
//	if amount == 0 {
//		return 0
//	}
//	if amount < 0 {
//		return -1
//	}
//	if memo[amount] != -666 {
//		return memo[amount]
//	}
//	res := math.MaxInt32
//	for _, coin := range coins {
//		subProblem := dp(coins, amount-coin, memo)
//		if subProblem == -1 {
//			continue
//		}
//		res = min(res, subProblem+1)
//	}
//	if res == math.MaxInt32 {
//		memo[amount] = -1
//	} else {
//		memo[amount] = res
//	}
//	return memo[amount]
//}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	for i := 0; i < amount+1; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCoinChange(t *testing.T) {

	// your test code here
	coins := []int{2}
	println(coinChange(coins, 3))
}
