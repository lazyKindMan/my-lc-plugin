package leetcode_solutions

import "testing"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	p, q := head, head
	for i := 0; i < k; i++ {
		if q == nil {
			return head
		}
		q = q.Next
	}
	newHead := reverseNList(p, k)
	p.Next = reverseKGroup(q, k)
	return newHead
}

func reverseNList(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, cur, nxt *ListNode
	pre, cur, nxt = (*ListNode)(nil), head, head.Next
	for n > 0 {
		n--
		cur.Next = pre
		pre = cur
		cur = nxt
		if nxt != nil {
			nxt = nxt.Next
		}
	}
	head.Next = cur
	return pre
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseNodesInKGroup(t *testing.T) {
	reverseKGroup(CreateHead([]int{1, 2, 3, 4, 5}), 2)
	// your test code here

}
