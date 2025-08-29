package main

import "fmt"

func partitionLabels(s string) []int {
	result, characterBounds := []int{}, map[rune][]int{}

	for i, c := range s {
		_, exists := characterBounds[c]
		if !exists {
			characterBounds[c] = []int{i, i}
		} else {
			characterBounds[c][1] = i
		}
	}

	currentOccurence, lastIndex := 0, 0
	for i, c := range s {
		bounds := characterBounds[c]
		fmt.Println(string(c))
		if bounds[0] == bounds[1] {

		} else if bounds[0] == i {
			currentOccurence++
		} else if bounds[1] == i {
			currentOccurence--
		}
		if currentOccurence == 0 {
			result = append(result, i+1-lastIndex)
			lastIndex = i + 1
		}
	}

	return result
}

// func main() {
// 	fmt.Println(partitionLabels("ababcbacadefegdehijhklij")) // 9,7,8
// }
