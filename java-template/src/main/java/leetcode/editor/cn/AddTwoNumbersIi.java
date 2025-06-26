package leetcode.editor.cn;

import java.util.*;
import leetcode.editor.common.*;

public class AddTwoNumbersIi {

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
        public ListNode addTwoNumbers(ListNode l1, ListNode l2){
            // 利用栈的先进后出思路避免链表反转（这也是很多操作计算的思路）
            ListNode dummy = new ListNode(-1);
            Stack<Integer> stack1 = new Stack<>();
            Stack<Integer> stack2 = new Stack<>();
            while (l1 != null) {
                stack1.push(l1.val);
                l1 = l1.next;
            }
            while (l2 != null) {
                stack2.push(l2.val);
                l2 = l2.next;
            }
            int carry = 0;
            while (!stack1.isEmpty() || !stack2.isEmpty() || carry > 0) {
                int sum = carry;
                if (!stack1.isEmpty()) {
                    sum += stack1.pop();
                }
                if (!stack2.isEmpty()) {
                    sum += stack2.pop();
                }
                carry = sum / 10;
                dummy.next = new ListNode(sum % 10, dummy.next);
            }
            return dummy.next;
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)
    class Solution1 {
        public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
            l1 = reverse(l1);
            l2 = reverse(l2);
            ListNode ans = new ListNode(-1);
            ListNode p = l1, q = l2, x = ans;
            int carry = 0;
            while (p != null || q!= null || carry >= 1) {
                int sum = carry;
                if (p != null) {
                    sum += p.val;
                    p = p.next;
                }
                if (q != null) {
                    sum += q.val;
                    q = q.next;
                }
                int pos = sum % 10;
                carry = sum / 10;
                x.next = new ListNode(pos);
                x = x.next;
            }
            return reverse(ans.next);
        }

        public ListNode reverse(ListNode head) {
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
    }
    
    public static void main(String[] args) {

    }
}