package leetcode_solutions

import (
	"container/heap"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	ans := make([][]int, 0)
	pq := &PriorityQueue373{}
	heap.Init(pq)
	for i := 0; i < len(nums1); i++ {
		heap.Push(pq, []int{nums1[i], nums2[0], 0})
	}
	for pq.Len() > 0 && k > 0 {
		minArr := heap.Pop(pq).([]int)
		ans = append(ans, []int{minArr[0], minArr[1]})
		nums2Index := minArr[2]
		if nums2Index < len(nums2)-1 {
			heap.Push(pq, []int{minArr[0], nums2[nums2Index+1], nums2Index + 1})
		}
		k--
	}
	return ans
}

type PriorityQueue373 [][]int

func (pq PriorityQueue373) Len() int { return len(pq) }
func (pq PriorityQueue373) Less(i, j int) bool {
	return (pq[i][0] + pq[i][1]) < (pq[j][0] + pq[j][1])
}

func (pq PriorityQueue373) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue373) Push(e interface{}) {
	*pq = append(*pq, e.([]int))
}

func (pq *PriorityQueue373) Pop() interface{} {
	old := *pq
	n := len(*pq)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindKPairsWithSmallestSums(t *testing.T) {
	// your test code here
	kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3)
}
