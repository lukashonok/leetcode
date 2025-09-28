package main

import (
	"slices"
	"strings"
)

func findItinerary(tickets [][]string) []string {
	adj := map[string][]string{}
	slices.SortFunc(tickets, func(a, b []string) int {
		return strings.Compare(b[1], a[1])
	})

	for _, t := range tickets {
		if _, ok := adj[t[0]]; !ok {
			adj[t[0]] = []string{}
		}
		adj[t[0]] = append(adj[t[0]], t[1])
	}

	path := []string{}
	var dfs func(from string)
	dfs = func(from string) {
		for len(adj[from]) > 0 {
			nextDest := adj[from][len(adj[from])-1]
			adj[from] = adj[from][:len(adj[from])-1]
			dfs(nextDest)
		}
		path = append(path, from)
	}
	dfs("JFK")
	slices.Reverse(path)
	return path
}

// func main() {
// 	fmt.Println(findItinerary([][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}))
// 	fmt.Println(findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}))
// 	fmt.Println(findItinerary([][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "JFK"}, {"ATL", "AAA"}, {"AAA", "ATL"}, {"ATL", "BBB"}, {"BBB", "ATL"}, {"ATL", "CCC"}, {"CCC", "ATL"}, {"ATL", "DDD"}, {"DDD", "ATL"}, {"ATL", "EEE"}, {"EEE", "ATL"}, {"ATL", "FFF"}, {"FFF", "ATL"}, {"ATL", "GGG"}, {"GGG", "ATL"}, {"ATL", "HHH"}, {"HHH", "ATL"}, {"ATL", "III"}, {"III", "ATL"}, {"ATL", "JJJ"}, {"JJJ", "ATL"}, {"ATL", "KKK"}, {"KKK", "ATL"}, {"ATL", "LLL"}, {"LLL", "ATL"}, {"ATL", "MMM"}, {"MMM", "ATL"}, {"ATL", "NNN"}, {"NNN", "ATL"}}))
// }
