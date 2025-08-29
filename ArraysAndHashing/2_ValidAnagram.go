package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sourceWordMap := map[rune]int{}
	targetWordMap := map[rune]int{}

	for _, character := range s {
		sourceWordMap[character]++
	}

	for _, character := range t {
		_, exists := sourceWordMap[character]
		if !exists {
			return false
		}

		targetWordMap[character]++

		if targetWordMap[character] > sourceWordMap[character] {
			return false
		}
	}

	return true
}
