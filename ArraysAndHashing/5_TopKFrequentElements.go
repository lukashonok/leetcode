package main

func topKFrequent(nums []int, k int) []int {
	topMap := map[int]int{}

	for _, num := range nums {
		topMap[num]++
	}

	tops := make([]int, k)

	for i := range tops {
		topMapNumToDelete := 0
		maxNum := -100000
		for numKey, num := range topMap {
			if num > maxNum {
				topMapNumToDelete = numKey
				maxNum = num
			}
		}
		tops[i] = topMapNumToDelete
		delete(topMap, topMapNumToDelete)
	}

	return tops
}
