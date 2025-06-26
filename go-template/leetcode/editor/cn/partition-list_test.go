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
func partition(head *ListNode, x int) *ListNode {
	dummy := ListNode{}
	listGreater := ListNode{}
	cur := head
	list1 := &dummy
	list2 := &listGreater
	for cur != nil {
		if cur.Val < x {
			list1.Next = cur
			list1 = list1.Next
		} else {
			list2.Next = cur
			list2 = list2.Next
		}
		temp := cur.Next
		cur.Next = nil
		cur = temp
	}
	list1.Next = listGreater.Next
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPartitionList(t *testing.T) {
	// your test code here
	head := CreateHead([]int{1, 4, 3, 2, 5, 2})
	t.Log(partition(head, 3))
}
