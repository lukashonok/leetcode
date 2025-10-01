package main

func shipWithinDays(weights []int, days int) int {
	maxValue, minValue := 0, weights[0]
	for _, w := range weights {
		maxValue += w
		if w > minValue {
			minValue = w
		}
	}

	left, right := minValue, maxValue

	for left < right {
		mid := (left + right) / 2
		curDays := 1
		currentWeight := 0

		for _, w := range weights {
			if currentWeight+w > mid {
				curDays++
				currentWeight = 0
			}
			currentWeight += w
		}

		if curDays <= days {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// func main() {
// 	fmt.Println(shipWithinDays([]int{1, 2, 3, 1, 1}, 4))                 // 3
// 	fmt.Println(shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3))              // 6
// 	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)) // 15
// }
