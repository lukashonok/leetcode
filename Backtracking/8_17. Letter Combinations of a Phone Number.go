package main

func letterCombinations(digits string) []string {
	keyboard := map[string]string{"2": "abc", "3": "def", "4": "ghi", "5": "jkl", "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz"}
	result, solution, length := []string{}, "", len(digits)

	if length == 0 {
		return result
	}

	var dfs func(i int)
	dfs = func(i int) {
		if i == length {
			result = append(result, solution)
			return
		}

		for _, v := range keyboard[string(digits[i])] {
			solution += string(v)
			dfs(i + 1)
			solution = solution[:len(solution)-1]
		}
	}
	dfs(0)

	return result
}

// func main() {
// 	fmt.Println(letterCombinations("23"))
// 	fmt.Println(letterCombinations("22"))
// }
