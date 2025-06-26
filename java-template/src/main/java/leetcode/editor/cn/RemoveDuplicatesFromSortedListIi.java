package leetcode.editor.cn;

import java.util.*;
import leetcode.editor.common.*;

public class RemoveDuplicatesFromSortedListIi {

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
        public ListNode deleteDuplicates(ListNode head) {
            ListNode dummy = new ListNode(-1, head);
            ListNode slow = dummy, fast = head;
            while (fast != null) {
                if (fast.next != null && fast.val == fast.next.val) {
                    while (fast.next != null && fast.val == fast.next.val) {
                        fast = fast.next;
                    }
                    fast = fast.next;
                    if (fast == null) {
                        slow.next = null;
                    }
                } else {
                    slow.next = fast;
                    slow = slow.next;
                    fast = fast.next;
                }
            }
            return dummy.next;
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

    
    public static void main(String[] args) {
        Solution solution = new RemoveDuplicatesFromSortedListIi().new Solution();
        // put your test code here
        
    }
}