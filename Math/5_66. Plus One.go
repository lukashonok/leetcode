package main

func plusOne(digits []int) []int {
	i := len(digits) - 1
	for i >= 0 {
		if digits[i]+1 != 10 {
			digits[i]++
			break
		}
		digits[i] = 0
		i--
	}
	if i == -1 {
		return append([]int{1}, digits...)
	}
	return digits
}

// func main() {
// 	fmt.Println(plusOne([]int{1, 2, 3}))
// 	fmt.Println(plusOne([]int{1, 2, 9}))
// 	fmt.Println(plusOne([]int{9, 9, 9}))
// }
