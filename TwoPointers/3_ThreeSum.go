package main

func qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	pivotIndex := len(a) / 2
	a[pivotIndex], a[right] = a[right], a[pivotIndex]
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]
	qsort(a[:left])
	qsort(a[left+1:])
	return a
}

func threeSum(nums []int) [][]int {
	numMap := map[int]int{}
	result := [][]int{}
	numsSortedRaw := qsort(nums)
	nums = []int{}
	for _, v := range numsSortedRaw {
		_, exists := numMap[v]
		if !exists {
			nums = append(nums, v)
		}
		numMap[v]++
	}

	if numMap[0] >= 3 {
		result = append(result, []int{0, 0, 0})
	}

	length := len(nums)
	for i := 0; i < length; i++ {
		for k := length - 1; k > i; k-- {
			target := -nums[i] - nums[k]

			if target > nums[k] || target < nums[i] {
				continue
			}

			occurrence, exists := numMap[target]

			if target == nums[i] && occurrence == 1 {
				exists = false
			}

			if target == nums[k] && occurrence == 1 {
				exists = false
			}

			if exists {
				result = append(result, []int{nums[i], target, nums[k]})
			}
		}
	}

	return result
}

// func main() {
// 	start := time.Now()
// 	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))                                        // [[-1,-1,2],[-1,0,1]]
// 	fmt.Println(threeSum([]int{0, 1, 1}))                                                    // []
// 	fmt.Println(threeSum([]int{0, 0, 0}))                                                    // [[0, 0, 0]]
// 	fmt.Println(threeSum([]int{0, 0, 0, 0}))                                                 // [[0, 0, 0]]
// 	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))                                             // [[-2,0,2],[-2,1,1]]
// 	fmt.Println(threeSum([]int{-2, -1, -1, 0, 2}))                                           // [[-2,0,2],[2,-1,-1]]
// 	fmt.Println(threeSum([]int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10})) // [[-10,5,5],[-5,0,5],[-4,2,2],[-3,-2,5],[-3,1,2],[-2,0,2]]
// 	fmt.Println("Total ms: ", time.Since(start).Milliseconds())
// }
