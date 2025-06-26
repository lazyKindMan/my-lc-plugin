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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var stkl1 []int
	var stkl2 []int
	for l1 != nil {
		stkl1 = append(stkl1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stkl2 = append(stkl2, l2.Val)
		l2 = l2.Next
	}

	dummy := &ListNode{-1, nil}
	carry := 0
	dummyNext := dummy.Next
	var sum int
	for len(stkl1) > 0 || len(stkl2) > 0 || carry > 0 {
		sum = carry
		n := len(stkl1)
		m := len(stkl2)
		if n > 0 {
			sum += stkl1[n-1]
			stkl1 = stkl1[:n-1]
		}
		if m > 0 {
			sum += stkl2[m-1]
			stkl2 = stkl2[:m-1]
		}
		carry = sum / 10
		dummy.Next = &ListNode{sum % 10, dummyNext}
		dummyNext = dummy.Next
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestAddTwoNumbersIi(t *testing.T) {

	// your test code here
	addTwoNumbers(CreateHead([]int{7, 2, 4, 3}), CreateHead([]int{5, 6, 4}))
}
