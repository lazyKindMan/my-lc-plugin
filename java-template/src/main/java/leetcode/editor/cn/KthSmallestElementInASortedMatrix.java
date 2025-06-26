package leetcode.editor.cn;

import java.util.*;
import leetcode.editor.common.*;

public class KthSmallestElementInASortedMatrix {

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int kthSmallest(int[][] matrix, int k) {
            int n = matrix.length;
            PriorityQueue<List<Integer>> pq = new PriorityQueue<>(n,
                    Comparator.comparingInt(o -> o.get(0)));
            for (int i = 0; i < n; i++) {
                pq.add(Arrays.asList(matrix[i][0], i, 0));
            }
            int res = 0;
            while (k > 0 && !pq.isEmpty()) {
                List<Integer> ele = pq.poll();
                res = ele.get(0);
                int rowNum = ele.get(1);
                int column = ele.get(2);
                if (column < n - 1) {
                    pq.add(Arrays.asList(matrix[rowNum][column + 1], rowNum, column + 1));
                }
                k --;
            }
            return res;
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

    
    public static void main(String[] args) {
        Solution solution = new KthSmallestElementInASortedMatrix().new Solution();
        // put your test code here
        
    }
}