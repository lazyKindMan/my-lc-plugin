package leetcode_solutions

import (
	"testing"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	return nil
}

func reverseN(head *ListNode, k int) {
	var pre, cur, next *ListNode
	pre, cur, next = nil, head, head.Next
	for k > 0 {
		cur.Next = pre
		pre = cur
		cur = next
		if next != nil {
			next = next.Next
		}
		k--
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseLinkedListIi(t *testing.T) {
	// your test code here
	reverseBetween(CreateHead([]int{1, 2, 3, 4, 5}), 2, 4)
}
