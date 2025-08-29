package main

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	m := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		m[i] = amount + 1
		for _, c := range coins {
			if i-c < 0 {
				continue
			}
			m[i] = min(m[i], 1+m[i-c])
		}
	}
	if m[len(m)-1] == amount+1 {
		return -1
	}
	return m[len(m)-1]
}

// func main() {
// 	fmt.Println(coinChange([]int{1, 2, 5}, 11))
// }
