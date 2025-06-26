package leetcode.editor.cn;

import java.util.*;
import leetcode.editor.common.*;

public class UglyNumberIi {

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int nthUglyNumber(int n) {
            int p2 = 1, p3 = 1, p5 = 1;
            int product2 = 1, product3 = 1, product5 = 1;
            int[] arr = new int[n + 1];
            int p = 1;
            while (p <= n) {
                int min = Math.min(Math.min(product2, product3), product5);
                arr[p] = min;
                p++;
                if (min == product2) {
                    product2 = arr[p2] * 2;
                    p2++;
                }
                if (min == product3) {
                    product3 = arr[p3] * 3;
                    p3++;
                }
                if (min == product5) {
                    product5 = arr[p5] * 5;
                    p5++;
                }
            }
            return arr[n];
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

    
    public static void main(String[] args) {
        Solution solution = new UglyNumberIi().new Solution();
        // put your test code here
        
    }
}