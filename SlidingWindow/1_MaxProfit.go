package main

func maxProfit(prices []int) int {
	profit := 0

	i, j := 0, 1
	for j < len(prices) {
		if prices[j] < prices[i] {
			i = j
			j++
			continue
		}

		profit = max(prices[j]-prices[i], profit)

		j++
	}

	return profit
}

// func main() {
// 	fmt.Println(maxProfit([]int{1}))                 // 0
// 	fmt.Println(maxProfit([]int{4, 1}))              // 0
// 	fmt.Println(maxProfit([]int{1, 4}))              // 3
// 	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))  // 5
// 	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))     // 0
// 	fmt.Println(maxProfit([]int{10, 1, 5, 6, 7, 1})) // 6
// 	fmt.Println(maxProfit([]int{10, 8, 7, 5, 2}))    // 0
// }
