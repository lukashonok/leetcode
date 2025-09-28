package main

func multiply(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	res := make([]int, l1+l2)
	resStr := ""
	shift := len(res) - 1
	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			product := int((num1[i] - '0') * (num2[j] - '0'))
			res[shift-1] += product / 10
			current := res[shift] + (product % 10)
			res[shift] = current % 10
			res[shift-1] += current / 10

			shift--
		}
		shift--
		shift += l2
	}
	i := 0
	for i < len(res) && res[i] == 0 {
		i++
	}
	for ; i < len(res); i++ {
		resStr += string(res[i] + '0')
	}
	if resStr == "" {
		return "0"
	}
	return string(resStr)
}

// func main() {
// 	fmt.Println(multiply("123", "456"))
// 	fmt.Println(multiply("9", "9"))
// 	fmt.Println(multiply("9", "99"))
// 	fmt.Println(multiply("0", "0"))
// }
