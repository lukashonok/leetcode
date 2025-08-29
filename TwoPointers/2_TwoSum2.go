package main

func twoSum(numbers []int, target int) []int {
	length := len(numbers)
	i := 0
	j := length - 1

	for i < j {
		currentSum := numbers[i] + numbers[j]
		if currentSum == target {
			return []int{i + 1, j + 1}
		}

		if currentSum > target {
			j--
		} else {
			i++
		}
	}

	return []int{i, j}
}

// func main() {
// 	start := time.Now()
// 	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [1, 2]
// 	fmt.Println(twoSum([]int{2, 3, 4}, 6))      // [1, 3]
// 	fmt.Println(twoSum([]int{-1, 0}, -1))       // [1, 2]
// 	fmt.Println("Total ms: ", time.Since(start).Milliseconds())
// }
