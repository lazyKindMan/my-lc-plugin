package leetcode.editor.cn;

import java.util.*;
import leetcode.editor.common.*;

public class MergeKSortedLists {

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
        public ListNode mergeKLists(ListNode[] lists) {
            if (lists.length == 0) {
                return null;
            }
            ListNode dummy = new ListNode(-1);
            ListNode p = dummy;
            PriorityQueue<ListNode> hq = new PriorityQueue<>(
                    lists.length,
                    Comparator.comparingInt(a -> a.val)
            );
            for (ListNode list : lists) {
                if (list != null) {
                    hq.add(list);
                }
            }
            while (!hq.isEmpty()) {
                ListNode cur = hq.poll();
                p.next = cur;
                p = p.next;
                if (cur.next != null) {
                    hq.add(cur.next);
                }
            }
            return dummy.next;
        }

    }
    //leetcode submit region end(Prohibit modification and deletion)

    
    public static void main(String[] args) {
        Solution solution = new MergeKSortedLists().new Solution();
        // put your test code here
        
    }
}