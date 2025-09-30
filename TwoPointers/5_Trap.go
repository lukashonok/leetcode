package main

func trap(height []int) int {
	maxTrapped := 0
	length := len(height)

	i, k := 0, 1

	currentTrapped := 0

	for k < length {
		if height[i] <= height[k] {
			i, k = k, k+1
			maxTrapped += currentTrapped
			currentTrapped = 0
		} else {
			currentTrapped += height[i] - height[k]
			k++
		}
	}

	if currentTrapped > 0 {
		currentTrapped = 0
		conflictIndex := i
		i, k = length-1, length-2
		for conflictIndex <= k {
			if height[i] <= height[k] {
				i, k = k, k-1
				maxTrapped += currentTrapped
				currentTrapped = 0
			} else {
				currentTrapped += height[i] - height[k]
				k--
			}
		}
	}

	return maxTrapped
}

// func main() {
// 	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6
// 	fmt.Println(trap([]int{0, 2, 0, 3, 1, 0, 1, 3, 2, 1}))       // 9
// 	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))                   // 9
// 	fmt.Println(trap([]int{4, 2, 3}))                            // 1
// }
