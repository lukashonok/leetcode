package main

func lastStoneWeight(stones []int) int {
	heapStones := []HeapData{}
	for _, s := range stones {
		heapStones = append(heapStones, HeapData{key: float32(s)})
	}
	MaxHeapBuild(heapStones)
	for len(heapStones) > 1 {
		x, y := MaxHeapPop(&heapStones), MaxHeapPop(&heapStones)
		if x.key > y.key {
			x, y = y, x
		}
		if x.key != y.key {
			MaxHeapInsert(&heapStones, HeapData{key: y.key - x.key})
		}
	}
	if len(heapStones) == 0 {
		return 0
	}
	return int(heapStones[0].key)
}

// func main() {
// 	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1})) // 1
// 	fmt.Println(lastStoneWeight([]int{1}))                // 1
// 	fmt.Println(lastStoneWeight([]int{}))                 // 0
// 	fmt.Println(lastStoneWeight([]int{1, 1}))             // 0
// }
