package main

func majorityElement(nums []int) []int {
	candidates := map[int]int{}
	for _, n := range nums {
		candidates[n]++
		if len(candidates) <= 2 {
			continue
		}
		newCandidates := map[int]int{}
		for n, c := range candidates {
			if c > 1 {
				newCandidates[n] = c - 1
			}
		}
		candidates = newCandidates
	}

	res, target := []int{}, len(nums)/3
	for candidate := range candidates {
		currentCount := 0
		for _, v := range nums {
			if v == candidate {
				currentCount++
			}
		}
		if currentCount > target {
			res = append(res, candidate)
		}
	}
	return res
}

// func main() {
// 	// fmt.Println(majorityElement([]int{1, 1, 2, 2, 2, 4, 4, 1, 1}))
// 	// fmt.Println(majorityElement([]int{1}))
// 	// fmt.Println(majorityElement([]int{1, 2}))
// 	fmt.Println(majorityElement([]int{2, 2}))
// }
