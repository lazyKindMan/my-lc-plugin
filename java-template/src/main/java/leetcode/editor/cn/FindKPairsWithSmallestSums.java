package leetcode.editor.cn;

import java.util.*;
import leetcode.editor.common.*;

public class FindKPairsWithSmallestSums {

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public List<List<Integer>> kSmallestPairs(int[] nums1, int[] nums2, int k) {
            List<List<Integer>> ans = new ArrayList<>();
            PriorityQueue<int[]> pq = new PriorityQueue<>(
                    (o1, o2) -> o1[0] + o1[1] - o2[0] - o2[1]
            );
            for (int i = 0; i < nums1.length; i++) {
                pq.add(new int[]{nums1[i], nums2[0], i, 0});
            }
            while (!pq.isEmpty() && k > 0) {
                int[] minPair = pq.poll();
                int nums1Index = minPair[2];
                int nums2Index = minPair[3];
                if (nums2Index < nums2.length - 1) {
                    pq.offer(new int[] {nums1[nums1Index], nums2[nums2Index + 1], nums1Index, nums2Index + 1});
                }
                ans.add(Arrays.asList(minPair[0], minPair[1]));
                k --;
            }
            return ans;
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

    
    public static void main(String[] args) {
        Solution solution = new FindKPairsWithSmallestSums().new Solution();
        // put your test code here
        int[] nums1 = new int[] {1, 7, 11};
        int[] nums2 = new int[] {2, 4, 6};
        solution.kSmallestPairs(nums1, nums2, 3);
    }
}