package leetcode_solutions

import (
	"container/heap"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
type PriorityQueue378 [][]int

func (pq PriorityQueue378) Len() int {
	return len(pq)
}

func (pq PriorityQueue378) Less(i, j int) bool {
	return pq[i][0] < pq[j][0]
}

func (pq PriorityQueue378) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue378) Push(e interface{}) {
	*pq = append(*pq, e.([]int))
}

func (pq *PriorityQueue378) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	headIndex := make([]int, n)
	pq := &PriorityQueue378{}
	heap.Init(pq)
	for i := 0; i < n; i++ {
		headIndex[i] = 1
		heap.Push(pq, []int{matrix[i][0], i})
	}
	for i := k; i > 1; i-- {
		ele := heap.Pop(pq).([]int)
		row := ele[1]
		if headIndex[row] < n {
			heap.Push(pq, []int{matrix[row][headIndex[row]], row})
			headIndex[row]++
		}
	}
	return heap.Pop(pq).([]int)[0]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestKthSmallestElementInASortedMatrix(t *testing.T) {
	// your test code here
	matrix := [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}
	t.Log(kthSmallest(matrix, 8))
}
