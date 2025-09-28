package main

import (
	"math"
)

func myPow(x float64, n int) float64 {
	if n == 0 || x == 1 {
		return 1
	}
	if x == -1 {
		if n&1 == 0 {
			return 1
		} else {
			return -1
		}
	}
	negativePower, n := n < 0, int(math.Abs(float64(n)))
	if x == 2 {
		x = float64(int(x) << (n - 1))
		if negativePower && x != 0 {
			x = 1 / x
		}
		return x
	}
	cur := x
	for n > 1 {
		cur *= x
		n--
	}
	if negativePower {
		cur = 1 / cur
	}
	return cur
}

// func main() {
// 	// fmt.Println(myPow(2.0, 10))
// 	// fmt.Println(myPow(2.1, 3))
// 	// fmt.Println(myPow(2, -2))
// 	// fmt.Println(myPow(2, -2147483648))
// 	fmt.Println(myPow(-1, 2147483648))
// 	fmt.Println(myPow(-1, 2147483647))
// }
