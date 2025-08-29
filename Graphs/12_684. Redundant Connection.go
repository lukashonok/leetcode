package main

func findRedundantConnection(edges [][]int) []int {
	parents, ranks := []int{}, []int{}
	for i := 0; i <= len(edges); i++ {
		parents = append(parents, i)
		ranks = append(ranks, 1)
	}

	var find func(n int) int
	find = func(n int) int {
		if n != parents[n] {
			parents[n] = find(parents[n])
		}
		return parents[n]
	}

	union := func(n1, n2 int) bool {
		p1, p2 := find(n1), find(n2)
		if p1 == p2 {
			return false
		}

		if ranks[p1] > ranks[p2] {
			parents[p2] = p1
			ranks[p1] += ranks[p2]
		} else {
			parents[p1] = p2
			ranks[p2] += ranks[p1]
		}

		return true
	}

	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return edge
		}
	}

	return []int{}
}

// func main() {
// 	fmt.Println(findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))
// }
