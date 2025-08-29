package main

func getStringCharactersMap(str string) map[rune]int {
	charactersMap := map[rune]int{}
	for _, character := range str {
		charactersMap[character]++
	}
	return charactersMap
}

func isMapStrAnagram(characterMap map[rune]int, s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	targetWordMap := map[rune]int{}
	for _, character := range t {
		_, exists := characterMap[character]
		if !exists {
			return false
		}

		targetWordMap[character]++

		if targetWordMap[character] > characterMap[character] {
			return false
		}
	}
	return true
}

func groupAnagrams(strs []string) [][]string {
	var result [][]string = [][]string{}

	for len(strs) > 0 {
		currentStr := strs[0]
		indexesToRemove := []int{}
		currentStrMap := getStringCharactersMap(currentStr)
		strs = strs[1:]
		strsToResult := []string{currentStr}
		for i, str := range strs {
			if isMapStrAnagram(currentStrMap, currentStr, str) {
				indexesToRemove = append(indexesToRemove, i)
				strsToResult = append(strsToResult, str)
			}
		}

		for i := len(indexesToRemove) - 1; i >= 0; i-- {
			strs = append(strs[:indexesToRemove[i]], strs[indexesToRemove[i]+1:]...)
		}

		result = append(result, strsToResult)
	}

	return result
}
