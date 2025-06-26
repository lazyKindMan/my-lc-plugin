package leetcode.editor.cn;

import java.util.*;
import leetcode.editor.common.*;

public class ReverseNodesInKGroup {

    //leetcode submit region begin(Prohibit modification and deletion)
    /**
     * Definition for singly-linked list.
     * public class ListNode {
     *     int val;
     *     ListNode next;
     *     ListNode() {}
     *     ListNode(int val) { this.val = val; }
     *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
     * }
     */
    class Solution {
        public ListNode reverseKGroup(ListNode head, int k) {
            if (k == 1) {
                return head;
            }
            ListNode dummy = new ListNode(-1, head);
            ListNode p = dummy, q = dummy;
            int n = k;
            while (q != null) {
                if (n == 0) {
                   ListNode reverseHead = reverseN(p.next, k);
                   q = p.next;
                   p.next = reverseHead;
                   p = q;
                   n = k;
                } else {
                    n --;
                    q = q.next;
                }
            }
            return dummy.next;
        }

        public ListNode reverseN(ListNode head, int n) {
            if (head == null || head.next == null) {
                return head;
            }
            ListNode pre = null, cur = head, nxt = head.next;
            while (n > 0) {
                n --;
                cur.next = pre;
                pre = cur;
                cur = nxt;
                if (nxt != null) {
                    nxt = nxt.next;
                }
            }
            head.next = cur;
            return pre;
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

    
    public static void main(String[] args) {
        Solution solution = new ReverseNodesInKGroup().new Solution();
        // put your test code here
        solution.reverseKGroup(ListNode.createHead(new int[] {1, 2}), 2);
    }
}