package leetcode_solutions

import (
	"testing"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	var maxDepth func(root *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftMax := maxDepth(root.Left)
		rightMax := maxDepth(root.Right)
		diameter := leftMax + rightMax
		if diameter > maxDiameter {
			maxDiameter = diameter
		}
		return 1 + max(leftMax, rightMax)
	}
	maxDepth(root)
	return maxDiameter
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDiameterOfBinaryTree(t *testing.T) {
	// your test code here

}
