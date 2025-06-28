package leetcode.editor.cn;

import java.util.*;
import leetcode.editor.common.*;

public class FindSubsequenceOfLengthKWithTheLargestSum {

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        List<Integer> path = new ArrayList<>();
        int sum = 0;
        int ansSum = Integer.MIN_VALUE;
        int[] ans = null;
        public int[] maxSubsequence(int[] nums, int k) {
            dfs(nums, 0, k);
            return ans;
        }

        public void dfs(int[] nums, int start, int k) {
            if (path.size() == k) {
                if (sum > ansSum) {
                    ansSum = sum;
                    ans = path.stream().mapToInt(Integer::intValue).toArray();
                }
                return;
            }
            if (start == nums.length) {
                return;
            }
            path.add(nums[start]);
            sum += nums[start];
            dfs(nums, start + 1, k);
            path.remove(path.size() - 1);
            sum -= nums[start];

            dfs(nums, start + 1, k);
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

    
    public static void main(String[] args) {
        Solution solution = new FindSubsequenceOfLengthKWithTheLargestSum().new Solution();
        // put your test code here
        solution.maxSubsequence(new int[]{2, 1, 3, 3}, 2);
    }
}