package main

type KthLargest struct {
	data []HeapData
	k    int
}

func Constructor2(k int, nums []int) KthLargest {
	this := KthLargest{data: []HeapData{}, k: k}
	for _, num := range nums {
		this.Add(num)
	}
	return this
}

func (this *KthLargest) getKth() int {
	return 0
}

func (this *KthLargest) Add(val int) int {
	MinHeapInsert(&this.data, HeapData{key: float32(val)})
	if len(this.data) > this.k {
		MinHeapPop(&this.data)
	}
	return int(this.data[0].key)
}

// func main() {
// 	k, nums := 3, []int{4, 5, 8, 2}

// 	kthLargest := Constructor(k, nums)

// 	fmt.Println(kthLargest.Add(3))  // return 4
// 	fmt.Println(kthLargest.Add(5))  // return 5
// 	fmt.Println(kthLargest.Add(10)) // return 5
// 	fmt.Println(kthLargest.Add(9))  // return 8
// 	fmt.Println(kthLargest.Add(4))  // return 8

// }
