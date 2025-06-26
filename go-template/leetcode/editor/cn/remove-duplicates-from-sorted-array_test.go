package leetcode_solutions

import "testing"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

// leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates1(nums []int) int {
	// 写的太复杂了
	n := len(nums)
	r := n
	slow, fast := 0, 1
	for fast < n {
		for fast < n && nums[slow] == nums[fast] {
			fast++
			r--
		}
		if fast < n {
			nums[slow+1] = nums[fast]
		}
		slow++
		fast++
	}
	return r
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	// your test code here

}
