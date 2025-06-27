package leetcode.editor.cn;

import java.util.*;
import leetcode.editor.common.*;

public class MinimumRecolorsToGetKConsecutiveBlackBlocks {

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int minimumRecolors(String blocks, int k) {
            int ans = Integer.MAX_VALUE;
            char[] arr = blocks.toCharArray();
            int whiteNum = 0;
            for (int i = 0; i < arr.length; i ++) {
                if (arr[i] == 'W') {
                    whiteNum ++;
                }
                if (i < k - 1) {
                    continue;
                }
                ans = Math.min(whiteNum, ans);
                if (arr[i - k + 1] == 'W') {
                    whiteNum --;
                }
            }
            return ans;
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

    
    public static void main(String[] args) {
        Solution solution = new MinimumRecolorsToGetKConsecutiveBlackBlocks().new Solution();
        // put your test code here
        solution.minimumRecolors("WBBWWBBWBW", 7);
    }
}