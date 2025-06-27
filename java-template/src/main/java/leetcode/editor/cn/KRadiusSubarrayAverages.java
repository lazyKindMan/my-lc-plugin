package leetcode.editor.cn;

import java.util.*;
import leetcode.editor.common.*;

public class KRadiusSubarrayAverages {

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int[] getAverages(int[] nums, int k) {
            int[] ans = new int[nums.length];
            long sum = 0;
            Arrays.fill(ans, -1);
            for (int i = 0; i < nums.length; i++) {
                sum += nums[i];
                if (i < 2 * k) {
                    continue;
                }
                ans[i - k] = (int) (sum / (2 * k + 1));
                sum -= nums[i - 2 * k];
            }
            return ans;
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

    
    public static void main(String[] args) {
        Solution solution = new KRadiusSubarrayAverages().new Solution();
        // put your test code here
        solution.getAverages(new int[]{7,4,3,9,1,8,5,2,6}, 3);
    }
}