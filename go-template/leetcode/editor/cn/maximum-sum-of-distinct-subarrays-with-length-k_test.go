package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func maximumSubarraySum(nums []int, k int) int64 {
	ans := int64(0)
	cnt := map[int]int{}
	s := int64(0)
	for i, num := range nums {
		cnt[num]++
		s += int64(num)
		if i < k-1 {
			continue
		}
		if len(cnt) == k && ans < s {
			ans = s
		}
		out := nums[i-k+1]
		s -= int64(out)
		cnt[out]--
		if cnt[out] == 0 {
			delete(cnt, out)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumSumOfDistinctSubarraysWithLengthK(t *testing.T) {
	// your test code here

}
