package main

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	l := len(s1)
	target, current := map[byte]int{}, map[byte]int{}
	for i := range l {
		target[s1[i]]++
		current[s2[i]]++
	}

	eq := func() bool {
		if len(current) != len(target) {
			return false
		}
		for k, v := range current {
			if w, ok := target[k]; !ok || v != w {
				return false
			}
		}
		return true
	}

	i, j := 0, l-1
	current[s2[j]]--
	if current[s2[j]] == 0 {
		delete(current, s2[j])
	}
	for j < len(s2) {
		current[s2[j]]++

		_, ok := target[s2[j]]
		if ok && eq() {
			return true
		}

		current[s2[i]]--
		if current[s2[i]] == 0 {
			delete(current, s2[i])
		}
		i++
		j++
	}

	return false
}

// func main() {
// 	fmt.Println(checkInclusion("ab", "eidbaooo"))  // true
// 	fmt.Println(checkInclusion("ab", "eidboaoo"))  // false
// 	fmt.Println(checkInclusion("abc", "lecabee"))  // true
// 	fmt.Println(checkInclusion("abc", "lecaabee")) // false
// 	fmt.Println(checkInclusion("abc", "dca"))      // false
// 	fmt.Println(checkInclusion("abc", "dc"))       // false
// }
