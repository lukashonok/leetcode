package main

import (
	"math"
	"slices"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	prices := make([]int, n)
	for i := range prices {
		prices[i] = math.MaxInt
	}
	prices[src] = 0

	for range k + 1 {
		tmpPrices := slices.Clone(prices)

		for _, flight := range flights {
			s, d, p := flight[0], flight[1], flight[2]
			if prices[s] == math.MaxInt {
				continue
			} else if prices[s]+p < tmpPrices[d] {
				tmpPrices[d] = prices[s] + p
			}
		}

		prices = tmpPrices
	}

	if prices[dst] == math.MaxInt {
		return -1
	}

	return prices[dst]
}

// func main() {
// 	fmt.Println(findCheapestPrice(4, [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}, 0, 3, 1))
// }
