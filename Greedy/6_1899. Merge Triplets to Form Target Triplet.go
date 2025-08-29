package main

func mergeTriplets(triplets [][]int, target []int) bool {
	resultMap := map[int]bool{0: false, 1: false, 2: false}

	for _, triplet := range triplets {

		validTriplet := true
		for i := 0; i < 3; i++ {
			if triplet[i] == target[i] {
				triplet[i] = 1
			} else if triplet[i] < target[i] {
				triplet[i] = 0
			} else {
				validTriplet = false
				break
			}
		}

		if validTriplet {
			for i := 0; i < 3; i++ {
				if triplet[i] == 1 {
					resultMap[i] = true
				}
			}
		}
	}

	return resultMap[0] && resultMap[1] && resultMap[2]
}

// func main() {

// 	fmt.Println(mergeTriplets([][]int{{1, 3, 1}}, []int{1, 3, 2}))                       // true
// 	fmt.Println(mergeTriplets([][]int{{1, 2, 3}, {7, 1, 1}}, []int{7, 2, 3}))            // true
// 	fmt.Println(mergeTriplets([][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}, []int{2, 7, 5})) // true
// 	fmt.Println(mergeTriplets([][]int{{3, 4, 5}, {4, 5, 6}}, []int{3, 2, 5}))            // false
// 	fmt.Println(mergeTriplets([][]int{{2, 5, 6}, {1, 4, 4}, {5, 7, 5}}, []int{5, 4, 6})) // false
// }
