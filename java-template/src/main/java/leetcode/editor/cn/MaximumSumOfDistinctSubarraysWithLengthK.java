package leetcode.editor.cn;

import java.util.*;
import leetcode.editor.common.*;

public class MaximumSumOfDistinctSubarraysWithLengthK {

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public long maximumSubarraySum(int[] nums, int k) {
            long ans = 0;
            long sum = 0;
            int left = 0;
            Set<Integer> cnt = new HashSet<>();
            for (int i = 0; i < nums.length; i++) {
                int x = nums[i];
                sum += x;
                while (cnt.contains(x)) {
                    cnt.remove(nums[left]);
                    sum -= nums[left];
                    left++;
                }
                cnt.add(x);
                if (cnt.size() < k) {
                    continue;
                }
                ans = Math.max(ans, sum);
                sum -= nums[left];
                cnt.remove(nums[left]);
                left ++;
            }
            return ans;
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

    
    public static void main(String[] args) {
        Solution solution = new MaximumSumOfDistinctSubarraysWithLengthK().new Solution();
        // put your test code here
        solution.maximumSubarraySum(new int[]{1, 5, 4, 2, 9, 9, 9}, 3);
    }
}