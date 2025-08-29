package main

import "fmt"

func singleNumber(nums []int) int {
	m := 0
	for _, n := range nums {
		m ^= n
	}
	return m
}

func hammingWeight(n int) int {
	i := 0
	for n > 0 {
		if n&0x01 == 1 {
			i++
		}
		n >>= 1
	}
	return i
}

func countBits(n int) []int {
	total := []int{0}
	for i := 1; i <= n; i++ {
		total = append(total, hammingWeight(i))
	}
	return total
}

func main() {

	fmt.Println(countBits(5))
	fmt.Println(hammingWeight(2147483645))
	fmt.Println(hammingWeight(3))

	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2, -1, -1}))
}
