package main

import (
	"fmt"
	"math"
)

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

func reverseBits(n int) int {
	r := 0
	for range 31 {
		r |= 1 & n
		n >>= 1
		r <<= 1
	}
	return r
}

func missingNumber(nums []int) int {
	n := 1
	for i := 2; i <= len(nums); i++ {
		n ^= i
	}
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		m ^= nums[i]
	}
	return n ^ m
}

func getSum(a int, b int) int {
	sum, carry := a^b, (a&b)<<1
	for carry != 0 {
		sum = a ^ b
		carry = (a & b) << 1
		a = sum
		b = carry
	}
	return sum
}

func reverse(x int) int {

	minInt := math.MinInt32 // -2147483648
	maxInt := math.MaxInt32 // 2147483647

	var n int32 = 0
	for x != 0 {
		reminder := x % 10
		x /= 10

		if n > int32(maxInt/10) || n == int32(maxInt) && reminder >= 7 {
			return 0
		}
		if n < int32(minInt/10) || n == int32(minInt) && reminder <= 8 {
			return 0
		}

		n *= 10
		n += int32(reminder)
	}
	return int(n)
}

func main() {

	fmt.Println(reverse(123))
	fmt.Println(reverse(-123456789))
	fmt.Println(reverse(0))

	// fmt.Println(getSum(1, 2)) // 15
	// fmt.Println(getSum(2, 3)) // 15

	// fmt.Println(missingNumber([]int{3, 0, 1}))    // 2
	// fmt.Println(missingNumber([]int{0, 1}))       // 2
	// fmt.Println(missingNumber([]int{0, 1, 4, 3})) // 1

	// fmt.Println(reverseBits(43261596))   // 964176192
	// fmt.Println(reverseBits(2147483644)) // 1073741822
	// fmt.Println(countBits(5))
	// fmt.Println(hammingWeight(2147483645))
	// fmt.Println(hammingWeight(3))

	// fmt.Println(singleNumber([]int{4, 1, 2, 1, 2, -1, -1}))
}
