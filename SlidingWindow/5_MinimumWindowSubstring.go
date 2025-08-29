package main

func minWindow(s string, t string) string {
	minWindow := ""

	if len(t) > len(s) {
		return ""
	}

	sourceLength := len(s)
	target, current := map[byte]int{}, map[byte]int{}
	for i := range t {
		target[t[i]]++
	}

	eq := func() bool {
		if len(current) < len(target) {
			return false
		}
		for k, v := range target {
			if w, ok := current[k]; !ok || v > w {
				return false
			}
		}
		return true
	}

	i, j := 0, 0
	current[s[0]]++
	for i < sourceLength {
		if eq() {
			if (j-i+1) < len(minWindow) || minWindow == "" {
				minWindow = s[i : j+1]
			}
			current[s[i]]--
			i++
		} else if j == sourceLength-1 {
			i++
			if i == sourceLength {
				break
			}
			current[s[i]]--
		} else {
			j++
			current[s[j]]++
		}
	}

	return minWindow
}

// func main() {
// 	fmt.Println(minWindow("ADOBECODEBANC", "ABC")) // BANC
// 	fmt.Println(minWindow("a", "a"))               // a
// 	fmt.Println(minWindow("a", "aa"))              // ""
// 	fmt.Println(minWindow("OUZODYXAZV", "XYZ"))    // YXAZ
// 	fmt.Println(minWindow("xyz", "xyz"))           // xyzxyz
// 	fmt.Println(minWindow("x", "xy"))              // ""
// 	fmt.Println(minWindow("aaaa", "a"))            // ""
// }
