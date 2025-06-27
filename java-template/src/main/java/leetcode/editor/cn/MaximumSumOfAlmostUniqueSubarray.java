package leetcode.editor.cn;

import java.util.*;
import leetcode.editor.common.*;

public class MaximumSumOfAlmostUniqueSubarray {

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public long maxSum(List<Integer> nums, int m, int k) {
            Map<Integer, Integer> cnt = new HashMap<>();
            long sum = 0;
            long ans = 0;
            for (int i = 0; i < nums.size(); i ++) {
                int key = nums.get(i);
                sum += key;
                cnt.merge(key, 1, Integer::sum);
                if (i < k - 1) {
                    continue;
                }
                if (cnt.keySet().size() >= m) {
                    ans = Math.max(ans, sum);
                }
                int out = nums.get(i - k + 1);
                cnt.put(out, cnt.get(out) - 1);
                if (cnt.get(out) == 0) {
                    cnt.remove(out);
                }
                sum -= out;
            }
            return ans;
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

    
    public static void main(String[] args) {
        Solution solution = new MaximumSumOfAlmostUniqueSubarray().new Solution();
        // put your test code here
        solution.maxSum(Arrays.asList(2,6,7,3,1,7), 3, 4);
    }
}