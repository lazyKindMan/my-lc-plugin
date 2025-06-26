package leetcode.editor.cn;

import java.util.*;
import leetcode.editor.common.*;

public class MaximumNumberOfVowelsInASubstringOfGivenLength {

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
       static Set<Character> vowels = new HashSet<>();
       static {
           vowels.add('a');
           vowels.add('e');
           vowels.add('i');
           vowels.add('o');
           vowels.add('u');
       }

        public int maxVowels(String s, int k) {
            char[] arr = s.toCharArray();
            int ans = -1;
            int vowelNum = 0;
            int windowSize = 0;
            int windowLeftIndex = 0;
            for (int i = 0; i < arr.length; i++) {
                if (vowels.contains(arr[i])) {
                    vowelNum++;
                }
                if (windowSize < k) {
                    windowSize ++;
                } else {
                    if (vowels.contains(arr[windowLeftIndex])) {
                        vowelNum--;
                    }
                    windowLeftIndex++;
                }
                ans = Math.max(ans, vowelNum);
            }
            return ans;
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

    
    public static void main(String[] args) {
        Solution solution = new MaximumNumberOfVowelsInASubstringOfGivenLength().new Solution();
        // put your test code here
        
    }
}