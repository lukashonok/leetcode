package main

import (
	"container/heap"
	"fmt"
	"math"
)

type IntHeap [][2]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.([2]int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func countRestrictedPaths(n int, edges [][]int) int {
	adj := map[int][][2]int{}
	for _, edge := range edges {
		if _, ok := adj[edge[0]]; !ok {
			adj[edge[0]] = [][2]int{}
		}
		if _, ok := adj[edge[1]]; !ok {
			adj[edge[1]] = [][2]int{}
		}
		adj[edge[0]] = append(adj[edge[0]], [2]int{edge[1], edge[2]})
		adj[edge[1]] = append(adj[edge[1]], [2]int{edge[0], edge[2]})
	}
	distances := make([]int, n+1)
	for i := 0; i < len(distances); i++ {
		distances[i] = math.MaxInt
	}
	distances[n] = 0

	h := &IntHeap{}
	heap.Init(h)
	heap.Push(h, [2]int{n, 0})

	for h.Len() > 0 {
		cur := heap.Pop(h).([2]int)

		if cur[1] != distances[cur[0]] {
			continue
		}

		for _, edge := range adj[cur[0]] {
			newDistance := distances[cur[0]] + edge[1]
			if distances[edge[0]] > newDistance {
				distances[edge[0]] = newDistance
				heap.Push(h, [2]int{edge[0], newDistance})
			}
		}
	}

	var dfs func(v int) int
	dfs = func(v int) int {
		if v == n {
			return 1
		}

		res := 0

		for _, u := range adj[v] {
			if distances[v] > distances[u[0]] {
				res = (res + dfs(u[0])) % int(math.Pow10(9)+7)
			}
		}

		return res
	}

	return dfs(1)
}

func main() {
	fmt.Println(countRestrictedPaths(5, [][]int{
		{1, 2, 3},
		{1, 3, 3},
		{2, 3, 1},
		{1, 4, 2},
		{5, 2, 2},
		{3, 5, 1},
		{5, 4, 10},
	}))
}
