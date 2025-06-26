package leetcode_solutions

import (
	"strings"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func slidingPuzzle(board [][]int) int {
	target := "123450"
	start := ""
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			start += string(rune(board[i][j] + '0'))
		}
	}
	queue := []string{start}
	visited := make(map[string]bool)
	visited[start] = true
	res := 0
	for len(queue) > 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur == target {
				return res
			}
			for _, neighbor := range getNeighbors(cur) {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
					visited[neighbor] = true
				}
			}
		}
		res++
	}
	return -1
}

func getNeighbors(board string) []string {
	mapping := [][]int{
		{1, 3},
		{0, 4, 2},
		{1, 5},
		{0, 4},
		{3, 1, 5},
		{4, 2},
	}
	idx := strings.Index(board, "0")
	var neighbors []string
	for _, adj := range mapping[idx] {
		newBoard := swap(board, idx, adj)
		neighbors = append(neighbors, newBoard)
	}
	return neighbors
}

func swap(board string, i, j int) string {
	chars := []rune(board)
	chars[i], chars[j] = chars[j], chars[i]
	return string(chars)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSlidingPuzzle(t *testing.T) {
	// your test code here
	board := [][]int{{1, 2, 3}, {5, 4, 0}}
	println(slidingPuzzle(board))
}
