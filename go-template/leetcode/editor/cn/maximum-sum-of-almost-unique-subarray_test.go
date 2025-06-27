package leetcode_solutions

import "testing"

// leetcode submit region begin(Prohibit modification and deletion)
func maxSum(nums []int, m int, k int) int64 {
	cnt := map[int]int{}
	sum := int64(0)
	ans := int64(0)
	for i := 0; i < len(nums); i++ {
		sum += int64(nums[i])
		cnt[nums[i]]++
		if i < k-1 {
			continue
		}
		if len(cnt) >= m && sum > ans {
			ans = sum
		}
		out := nums[i-k+1]
		cnt[out]--
		if cnt[out] == 0 {
			delete(cnt, out)
		}
		sum -= int64(out)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumSumOfAlmostUniqueSubarray(t *testing.T) {
	// your test code here

}
