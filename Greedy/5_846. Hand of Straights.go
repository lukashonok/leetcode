package main

import (
	"math"
	"slices"
)

func isNStraightHand(hand []int, groupSize int) bool {
	length := len(hand)
	if length%groupSize != 0 {
		return false
	}

	frequences := map[int]int{}

	for _, num := range hand {
		frequences[num]++
	}

	from := slices.Min(hand)
	for {
		currentFrequence := frequences[from]
		to := from + groupSize - 1
		_, exists := frequences[to]
		if !exists {
			return false
		}

		for i := from; i <= to; i++ {
			frequence, exists := frequences[i]
			if !exists || frequence-currentFrequence < 0 {
				return false
			}
			frequences[i] -= currentFrequence
			if frequences[i] == 0 {
				delete(frequences, i)
			}
		}

		if len(frequences) == 0 {
			break
		}

		from = math.MaxInt
		for frequence, _ := range frequences {
			if frequence < from {
				from = frequence
			}
		}
	}

	return true
}

// func main() {
// 	fmt.Println(isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3)) // true
// 	fmt.Println(isNStraightHand([]int{1, 2, 3, 4, 5}, 4))             // false
// 	fmt.Println(isNStraightHand([]int{8, 10, 12}, 3))                 // false
// }
