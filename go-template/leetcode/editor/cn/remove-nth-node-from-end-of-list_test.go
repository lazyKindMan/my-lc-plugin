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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	p := head
	for n > 0 {
		p = p.Next
		n--
	}
	q := dummy
	for p != nil {
		p = p.Next
		q = q.Next
	}
	q.Next = q.Next.Next
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRemoveNthNodeFromEndOfList(t *testing.T) {
	input := CreateHead([]int{1, 2, 3, 4, 5})
	expected := CreateHead([]int{1, 2, 3, 5})
	output := removeNthFromEnd(input, 2)
	if !IsEqual(output, expected) {
		t.Errorf("Expected %v but got %v", ListToSlice(expected), ListToSlice(output))
	}
}
