package main

func longestCommonPrefix(strs []string) string {
	res, curIndex := "", 0
	for {
		for i := 1; i < len(strs); i++ {
			if curIndex+1 > len(strs[i]) || curIndex+1 > len(strs[i-1]) || strs[i][curIndex] != strs[i-1][curIndex] {
				return res
			}
		}
		if curIndex+1 > len(strs[0]) {
			return res
		}
		res += string(strs[0][curIndex])
		curIndex++
	}
}

// func main() {
// 	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
// 	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
// 	fmt.Println(longestCommonPrefix([]string{""}))
// 	fmt.Println(longestCommonPrefix([]string{"", ""}))
// 	fmt.Println(longestCommonPrefix([]string{"", "", ""}))
// 	fmt.Println(longestCommonPrefix([]string{"", "b"}))
// }
