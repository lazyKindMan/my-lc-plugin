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
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Val: -101, Next: head}
	p, q := dummy, head
	for q != nil {
		if p.Val != q.Val {
			p.Next = q
			p = p.Next
		}
		q = q.Next
	}
	p.Next = nil
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRemoveDuplicatesFromSortedList(t *testing.T) {
	// your test code here

}
