package leetcode_solutions

import (
	"testing"
)

func openLock(deadends []string, target string) int {
	start := "0000"
	deads := make(map[string]struct{})
	for _, s := range deadends {
		deads[s] = struct{}{}
	}
	if _, found := deads["0000"]; found {
		return -1
	}
	q1 := make(map[string]struct{})
	q2 := make(map[string]struct{})
	visited := make(map[string]struct{})
	visited[start] = struct{}{}
	visited[target] = struct{}{}
	q1["0000"] = struct{}{}
	q2[target] = struct{}{}
	res := 0
	for len(q1) > 0 && len(q2) > 0 {
		res++
		newQ1 := make(map[string]struct{})
		for cur := range q1 {
			for _, neighbor := range getNeighbors1(cur) {
				if _, found := q2[neighbor]; found {
					return res
				}
				if _, found := deads[neighbor]; !found {
					if _, visit := visited[neighbor]; !visit {
						visited[neighbor] = struct{}{}
						newQ1[neighbor] = struct{}{}
					}
				}
			}
		}
		q1 = newQ1
		if len(q1) > len(q2) {
			q1, q2 = q2, q1
		}
	}
	return -1
}

// leetcode submit region begin(Prohibit modification and deletion)
func openLock1(deadends []string, target string) int {
	start := "0000"
	deads := make(map[string]struct{})
	for _, s := range deadends {
		deads[s] = struct{}{}
	}
	if _, found := deads["0000"]; found {
		return -1
	}
	visited := map[string]struct{}{}
	visited[start] = struct{}{}
	var q []string
	q = append(q, start)
	result := 0
	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			if cur == target {
				return result
			}
			for _, adj := range getNeighbors1(cur) {
				if _, found := visited[adj]; !found {
					if _, dead := deads[adj]; !dead {
						q = append(q, adj)
						visited[adj] = struct{}{}
					}
				}
			}
		}
		result++
	}
	return -1
}

func getNeighbors1(cur string) []string {
	var res []string
	for i := 0; i < len(cur); i++ {
		res = append(res, plusOne(cur, i))
		res = append(res, minusOne(cur, i))
	}
	return res
}

func plusOne(s string, j int) string {
	ch := []byte(s)
	if ch[j] == '9' {
		ch[j] = '0'
	} else {
		ch[j] += 1
	}
	return string(ch)
}

func minusOne(s string, j int) string {
	ch := []byte(s)
	if ch[j] == '0' {
		ch[j] = '9'
	} else {
		ch[j] -= 1
	}
	return string(ch)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestOpenTheLock(t *testing.T) {
	// your test code here
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	println(openLock(deadends, target))
}
