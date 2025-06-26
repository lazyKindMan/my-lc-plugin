package leetcode.editor.cn;

import java.util.*;
import leetcode.editor.common.*;

public class OpenTheLock {

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int openLock(String[] deadends, String target) {
            Set<String> deads = new HashSet<>(Arrays.asList(deadends));
            if (deads.contains("0000")) {
                return -1;
            }
            if ("0000".equals(target)) {
                return 0;
            }
            Set<String> q1 = new HashSet<>();
            q1.add("0000");
            Set<String> q2 = new HashSet<>();
            q2.add(target);
            Set<String> visited = new HashSet<>();
            visited.add("0000");
            int step = 0;
            while (!q1.isEmpty()) {
                Set<String> newQ = new HashSet<>();
                step ++;
                for (String cur : q1) {
                    for (String neighbor : getNeighbors(cur)) {
                        if (q2.contains(neighbor)) {
                            return step;
                        }
                        if (deads.contains(neighbor)) {
                            continue;
                        }
                        if (!visited.contains(neighbor)) {
                            newQ.add(neighbor);
                            visited.add(neighbor);
                        }
                    }
                }
                q1 = newQ;
                if (q1.size() > q2.size()) {
                    Set<String> tmp = q1;
                    q1 = q2;
                    q2 = tmp;
                }
            }
            return -1;
        }
        private Set<String> getNeighbors(String cur) {
            Set<String> result = new HashSet<>();
            for (int i = 0; i < cur.length(); i++) {
                result.add(plusOne(cur, i));
                result.add(minusOne(cur, i));
            }
            return result;
        }

        private String plusOne(String cur, int j) {
            char[] chars = cur.toCharArray();
            if (chars[j] == '9') {
                chars[j] = '0';
            } else {
                chars[j]++;
            }
            return String.valueOf(chars);
        }

        private String minusOne(String cur, int j) {
            char[] chars = cur.toCharArray();
            if (chars[j] == '0') {
                chars[j] = '9';
            } else {
                chars[j]--;
            }
            return String.valueOf(chars);
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

    
    public static void main(String[] args) {
        Solution solution = new OpenTheLock().new Solution();
        // put your test code here
        
    }
}