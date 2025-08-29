package main

func modifyMap(numsMap map[int]int, num int, sign int) int {
	consecutiveness := 1
	nextNum := num + sign
	_, exists := numsMap[nextNum]

	if exists {
		consecutiveness += modifyMap(numsMap, nextNum, sign)

		numsMap[num] = consecutiveness
		numsMap[nextNum] = consecutiveness
	} else {
		numsMap[num] = consecutiveness
	}

	return consecutiveness
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxConsecutive := 1

	numsMap := map[int]int{}

	for _, num := range nums {
		numsMap[num] = -1
	}

	for _, num := range nums {
		numInMap := numsMap[num]
		if numInMap != -1 {
			continue
		}

		positive := modifyMap(numsMap, num, -1)
		negative := modifyMap(numsMap, num, +1)

		consecutie := positive + negative - 1
		if consecutie > maxConsecutive {
			maxConsecutive = consecutie
		}
	}

	return maxConsecutive
}
