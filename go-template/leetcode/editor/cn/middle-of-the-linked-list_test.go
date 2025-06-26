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
func middleNode(head *ListNode) *ListNode {
	// 这里用快慢指针就行
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func middleNode1(head *ListNode) *ListNode {
	// 普通方法，先统计一遍链表长度，然后用双指针找到中间点
	n := 0
	p := head
	for p != nil {
		n++
		p = p.Next
	}
	k := (n + 1) / 2
	p = head
	q := head
	for k > 0 {
		p = p.Next
		k--
	}
	for p != nil {
		p = p.Next
		q = q.Next
	}
	return q
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMiddleOfTheLinkedList(t *testing.T) {
	// your test code here

}
