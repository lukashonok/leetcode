package main

func sortColors(nums []int) {
	counter := [3]int{0, 0, 0}
	for _, n := range nums {
		counter[n]++
	}
	currentLayer, i := 0, 0
	for {
		for counter[currentLayer] > 0 {
			nums[i] = currentLayer
			counter[currentLayer]--
			i++
		}
		currentLayer++
		if currentLayer == len(counter) {
			return
		}
	}
}

// func main() {
// 	a := []int{2, 0, 2, 1, 1, 0}
// 	sortColors(a)
// 	fmt.Println(a)
// 	a = []int{2, 0, 1}
// 	sortColors(a)
// 	fmt.Println(a)
// 	a = []int{1, 0, 1}
// 	sortColors(a)
// 	fmt.Println(a)
// 	a = []int{1, 2, 1}
// 	sortColors(a)
// 	fmt.Println(a)
// 	a = []int{2, 0}
// 	sortColors(a)
// 	fmt.Println(a)
// }
