package leetcode.editor.cn;

import java.util.*;
import leetcode.editor.common.*;

public class CountPrimes {

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int countPrimes(int n) {
            int[] isPrimes = new int[n];
            Arrays.fill(isPrimes, 1);
            int ans = 0;
            for (int i = 2; i < n; i++) {
                if (isPrimes[i] == 1) {
                    ans += 1;
                    if ((long) i * i < n) {
                        for (int j = i * i; j < n; j+=i) {
                            isPrimes[j] = 0;
                        }
                    }
                }
            }
            return ans;
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

    
    public static void main(String[] args) {
        Solution solution = new CountPrimes().new Solution();
        // put your test code here
        
    }
}