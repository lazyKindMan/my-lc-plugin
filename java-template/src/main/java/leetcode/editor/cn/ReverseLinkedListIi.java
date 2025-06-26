package leetcode.editor.cn;

import java.util.*;

import com.sun.security.auth.UnixNumericUserPrincipal;
import leetcode.editor.common.*;

public class ReverseLinkedListIi {

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
        public ListNode reverseBetween(ListNode head, int left, int right) {
            ListNode dummy = new ListNode(-1, head);
            ListNode leftHead = head, preLeftHead = dummy;
            for (int i = 1; i < left; i++) {
                leftHead = leftHead.next;
                preLeftHead = preLeftHead.next;
            }
            preLeftHead.next = reverseN(leftHead, right - left + 1);
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
        Solution solution = new ReverseLinkedListIi().new Solution();
        // put your test code here
        solution.reverseBetween(ListNode.createHead(new int[]{3, 5}), 1, 2);
    }
}