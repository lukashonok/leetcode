package main

func partition(s string) [][]string {
	result, part, length := [][]string{}, []string{}, len(s)

	isPali := func(i, j int) bool {
		for j > i {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}

	var dfs func(i int)
	dfs = func(i int) {
		if i >= length {
			currentSolution := make([]string, len(part))
			copy(currentSolution, part)
			result = append(result, currentSolution)
			return
		}

		for c := i; c < length; c++ {
			if isPali(i, c) {
				part = append(part, s[i:c+1])
				dfs(c + 1)
				part = part[:len(part)-1]
			}
		}
	}

	dfs(0)
	return result
}

// func main() {
// 	fmt.Println(partition("aab"))
// }
