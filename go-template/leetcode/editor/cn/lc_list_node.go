package leetcode_solutions

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateHead(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return head
}

func PrintList(l *ListNode) {
	cur := l
	for cur != nil {
		arrow := "->"
		if cur.Next == nil {
			arrow = ""
		}
		fmt.Printf("%d %s", cur.Val, arrow)
		cur = cur.Next
	}
}

func IsEqual(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1, l2 = l1.Next, l2.Next
	}
	return l1 == nil && l2 == nil
}

func ListToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
