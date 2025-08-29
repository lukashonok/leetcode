package main

func maxArea(height []int) int {
	length := len(height)
	maxArea, i, k := 0, 0, length-1

	for i < k {
		maxArea = max(maxArea, (k-i)*min(height[i], height[k]))

		if height[i] < height[k] {
			i++
		} else {
			k--
		}
	}

	return maxArea
}

// func main() {
// 	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
// }
