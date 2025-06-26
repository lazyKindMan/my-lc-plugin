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
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLinkedListCycleIi(t *testing.T) {
	// your test code here
	head := CreateHead([]int{3, 2, 0, -4})
	tail := head
	var entry *ListNode
	for tail.Next != nil {
		if tail.Val == 2 {
			entry = tail
		}
		tail = tail.Next
	}
	tail.Next = entry
	detectCycle(head)
}
