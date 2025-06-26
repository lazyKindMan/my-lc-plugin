package leetcode.editor.cn;

import java.util.*;
import leetcode.editor.common.*;

public class SlidingPuzzle {

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int slidingPuzzle(int[][] board) {
            String target = "123450";
            StringBuilder startStringBuilder = new StringBuilder();
            for (int[] ints : board) {
                for (int anInt : ints) {
                    startStringBuilder.append(anInt);
                }
            }
            String start = startStringBuilder.toString();
            Queue<String> q = new ArrayDeque<>();
            Set<String> visited = new HashSet<>();
            visited.add(start);
            int result = 0;
            q.add(start);
            while (!q.isEmpty()) {
                int sz = q.size();
                for (int i = 0; i < sz; i++) {
                    String cur = q.poll();
                    if (cur.equals(target)) {
                        return result;
                    }
                    for (String neighbor : getNeighbors(cur)) {
                        if (!visited.contains(neighbor)) {
                            visited.add(neighbor);
                            q.add(neighbor);
                        }
                    }
                }
                result ++;
            }
            return -1;
        }

        private List<String> getNeighbors(String board) {
            int[][] mapping = new int[][]{
                    {1, 3},
                    {0, 2, 4},
                    {1, 5},
                    {0, 4},
                    {1, 3, 5},
                    {2, 4}
            };
            int idx = board.indexOf("0");
            List<String> neighbors = new ArrayList<>();
            for (int adj : mapping[idx]) {
                neighbors.add(swap(board.toCharArray(), adj, idx));
            }
            return neighbors;
        }
        private String swap(char[] chars, int i, int j) {
            char temp = chars[i];
            chars[i] = chars[j];
            chars[j] = temp;
            return new String(chars);
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

    
    public static void main(String[] args) {
        Solution solution = new SlidingPuzzle().new Solution();
        // put your test code here
        int[][] input = new int[][]{
                {1, 2, 3},
                {4, 0, 5}
        };
        solution.slidingPuzzle(input);
    }
}