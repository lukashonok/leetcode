package main

func findKthLargest(nums []int, k int) int {
	data := []HeapData{}
	for _, num := range nums {
		MinHeapInsert(&data, HeapData{key: float32(num)})
		if len(data) > k {
			MinHeapPop(&data)
		}
	}
	return int(MinHeapPop(&data).key)
}

// func main() {
// 	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
// 	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
// }
