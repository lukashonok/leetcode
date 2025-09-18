package main

func removeDuplicates(nums []int) int {
	uniqueIndicies := []int{0}

	for i := 1; i < len(nums); i++ {
		if nums[uniqueIndicies[len(uniqueIndicies)-1]] != nums[i] {
			uniqueIndicies = append(uniqueIndicies, i)
		}
	}

	for i := 0; i < len(uniqueIndicies); i++ {
		nums[i] = nums[uniqueIndicies[i]]
	}

	return len(uniqueIndicies)
}

// func main() {
// 	fmt.Println(removeDuplicates([]int{1, 1, 2}))
// 	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
// }
