package leetcode.editor.cn;

import java.util.*;
import leetcode.editor.common.*;

public class DiameterOfBinaryTree {

    //leetcode submit region begin(Prohibit modification and deletion)
    /**
     * Definition for a binary tree node.
     * public class TreeNode {
     *     int val;
     *     TreeNode left;
     *     TreeNode right;
     *     TreeNode() {}
     *     TreeNode(int val) { this.val = val; }
     *     TreeNode(int val, TreeNode left, TreeNode right) {
     *         this.val = val;
     *         this.left = left;
     *         this.right = right;
     *     }
     * }
     */
    class Solution {
        int maxDiameter = 0;
        public int diameterOfBinaryTree(TreeNode root) {
            maxDepth(root);
            return maxDiameter;
        }

        public int maxDepth(TreeNode root) {
            if (root == null) {
                return 0;
            }
            int maxLeft = maxDepth(root.left);
            int maxRight = maxDepth(root.right);
            int diameter = maxLeft + maxRight;
            if (maxDiameter < diameter) {
                maxDiameter = diameter;
            }
            return Math.max(maxLeft, maxRight) + 1;
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

    
    public static void main(String[] args) {
        Solution solution = new DiameterOfBinaryTree().new Solution();
        // put your test code here
        TreeNode root = TreeNode.createRoot(new Integer[]{1,2,3,4,5});
        System.out.println(solution.diameterOfBinaryTree(root));
    }
}