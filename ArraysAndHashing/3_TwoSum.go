package main

func twoSum2(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, n := range nums {
		m[n] = i
	}
	for i, n := range nums {
		if j, ok := m[target-n]; ok && i != j {
			return []int{i, j}
		}
	}
	return []int{}
}

// func main() {
// 	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
// }
