package leetcode.editor.cn;

import java.util.*;
import leetcode.editor.common.*;

public class MaximumPointsYouCanObtainFromCards {

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int maxScore(int[] cardPoints, int k) {
            int n = cardPoints.length;
            int m = n - k;
            int sum = Arrays.stream(cardPoints).sum();
            if (m == 0) {
                return sum;
            }
            int ans = Integer.MAX_VALUE;
            int s = 0;
            for (int i = 0; i < n; i++) {
                s += cardPoints[i];
                if (i < m - 1) {
                    continue;
                }
                ans = Math.min(ans, s);
                s -= cardPoints[i - m + 1];
            }
            return sum - ans;
        }

        public int maxScore1(int[] cardPoints, int k) {
            int ans = 0, s = 0;
            int n = cardPoints.length;
            for (int i = n - k; i < n + k; i++) {
                int index = i % n;
                s += cardPoints[index];
                if (i < n - 1) {
                    continue;
                }
                ans = Math.max(ans, s);
                int outIndex = (i - k + 1) % n;
                s -= cardPoints[outIndex];
            }
            return ans;
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

    
    public static void main(String[] args) {
        Solution solution = new MaximumPointsYouCanObtainFromCards().new Solution();
        // put your test code here
        
    }
}