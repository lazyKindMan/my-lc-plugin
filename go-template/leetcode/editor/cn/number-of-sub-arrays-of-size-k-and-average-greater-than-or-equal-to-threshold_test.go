package leetcode_solutions

import "testing"

//leetcode submit region begin(Prohibit modification and deletion)
func numOfSubarrays(arr []int, k int, threshold int) int {
    ans := 0
	sum := 0
	threshold = threshold * k
	for i, num := range arr {
		sum += num
		if i < k - 1 {
			continue
		}
		if sum >= threshold {
			ans ++
		}
		sum -= arr[i - k + 1]
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)


func TestNumberOfSubArraysOfSizeKAndAverageGreaterThanOrEqualToThreshold(t *testing.T) {
	// your test code here
	
}