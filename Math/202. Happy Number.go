package main

func isHappy(n int) bool {
	m := map[int]bool{}
	for n != 1 {
		if _, ok := m[n]; ok {
			return false
		}
		m[n] = true

		newN := 0
		for n != 0 {
			d := n % 10
			newN += d * d
			n /= 10
		}
		n = newN
	}
	return true
}

// func main() {
// 	fmt.Println(isHappy(2))
// 	fmt.Println(isHappy(19))
// }
