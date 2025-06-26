package leetcode.editor.cn;

import java.util.*;
import leetcode.editor.common.*;

public class PartitionList {

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
        public ListNode partition(ListNode head, int x) {
            ListNode p1 = new ListNode(-1);
            ListNode dummy1 = p1;
            ListNode p2 = new ListNode(-1);
            ListNode dummy2 = p2;
            ListNode cur = head;
            while (cur != null) {
                if (cur.val < x) {
                    p1.next = cur;
                    p1 = p1.next;
                } else {
                    p2.next = cur;
                    p2 = p2.next;
                }
                cur = cur.next;
            }
            p1.next = dummy2.next;
            p2.next = null;
            return dummy1.next;
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

    
    public static void main(String[] args) {
        Solution solution = new PartitionList().new Solution();
        // put your test code here
        
    }
}