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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	// 这里最后如果没有相交 p1 和 p2 最后都会为nil 然后相等推出循环
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestIntersectionOfTwoLinkedLists(t *testing.T) {
	// your test code here

}
