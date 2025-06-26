package leetcode.editor.cn;

import java.util.*;
import leetcode.editor.common.*;

public class ReverseLinkedList {

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
        /**
         * 我自己的写法，会产生临时环
         * @param head
         * @return
         */
        public ListNode reverseList1(ListNode head) {
            if (head == null) {
                return head;
            }
            ListNode dummy = new ListNode(-1, head);
            ListNode p = dummy, q = head;
            while (q != null) {
                ListNode next = q.next;
                q.next = p;
                p = q;
                q = next;
            }
            head.next = null;
            return p;
        }
        public ListNode reverseList2(ListNode head) {
            if (head == null) {
                return head;
            }
            ListNode pre = null, cur = head, next = head.next;
            while (cur != null) {
                cur.next = pre;
                pre = cur;
                cur = next;
                if (next != null) {
                    next = next.next;
                }
            }
            return pre;
        }
        public ListNode reverseList(ListNode head) {
            if (head == null || head.next == null) {
                return head;
            }
            ListNode lastNode = reverseList(head.next);
            head.next.next = head;
            head.next = null;
            return lastNode;
        }

    }
    //leetcode submit region end(Prohibit modification and deletion)

    
    public static void main(String[] args) {
        Solution solution = new ReverseLinkedList().new Solution();
        // put your test code here
        solution.reverseList(ListNode.createHead(new int[] {1, 2, 3, 4, 5}));
    }
}