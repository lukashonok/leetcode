package main

func climbStairs(n int) int {
	a, b := 1, 1
	for range n - 1 {
		t := a
		a, b = a+b, t
	}
	return a
}

// func main() {
// 	fmt.Println(climbStairs(100))
// }
