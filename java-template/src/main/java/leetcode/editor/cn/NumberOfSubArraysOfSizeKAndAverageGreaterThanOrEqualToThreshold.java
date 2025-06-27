package leetcode.editor.cn;

import java.util.*;
import leetcode.editor.common.*;

public class NumberOfSubArraysOfSizeKAndAverageGreaterThanOrEqualToThreshold {

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int numOfSubarrays(int[] arr, int k, int threshold) {
            int ans = 0;
            int sum = 0;
            threshold = threshold * k;
            for (int i = 0; i < arr.length; i++) {
                sum += arr[i];
                if (i < k - 1) {
                    continue;
                }
                if (sum >= threshold) {
                    ans += 1;
                }
                sum -= arr[i - k + 1];
            }
            return ans;
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

    
    public static void main(String[] args) {
        Solution solution = new NumberOfSubArraysOfSizeKAndAverageGreaterThanOrEqualToThreshold().new Solution();
        // put your test code here
        
    }
}