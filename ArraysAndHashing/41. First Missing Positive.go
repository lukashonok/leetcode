package main

func firstMissingPositive(nums []int) int {
	size := len(nums)
	for i := size - 1; i >= 0; i-- {
		if nums[i] > 0 {
			break
		}
		size--
	}
	for i := size - 1; i >= 0; i-- {
		if nums[i] <= 0 {
			nums[i], nums[size-1] = nums[size-1], nums[i]
			size--
		}
	}
	arrayMap := make([]bool, size)
	for i := 0; i < size; i++ {
		if nums[i] > size {
			continue
		}
		arrayMap[nums[i]-1] = true
	}
	for i, exist := range arrayMap {
		if !exist {
			return i + 1
		}
	}
	return size + 1
}

// func main() {
// 	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1, -4}))
// 	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
// 	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
// }
