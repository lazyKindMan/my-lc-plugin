package leetcode.editor.cn;

import java.util.*;
import leetcode.editor.common.*;

public class NumberOfSubstringsContainingAllThreeCharacters {

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int numberOfSubstrings(String s) {
            char[] arr = s.toCharArray();
            int left = 0, ans = 0;
            int[] cnt = new int[] {0, 0, 0};
            for (int i = 0; i < arr.length; i++) {
                cnt[arr[i] - 'a'] ++;
                while (cnt[0] > 0 && cnt[1] > 0 && cnt[2] > 0) {
                    cnt[arr[left] - 'a'] --;
                    left ++;
                }
                ans += left;
            }
            return ans;
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

    
    public static void main(String[] args) {
        Solution solution = new NumberOfSubstringsContainingAllThreeCharacters().new Solution();
        // put your test code here
        
    }
}