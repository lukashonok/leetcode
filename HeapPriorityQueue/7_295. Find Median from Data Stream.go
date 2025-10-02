package main

func Heapify(data []int, n int, i int, minHeap bool) []int {
	largest := i
	l, r := i*2+1, i*2+2

	if l < n && ((!minHeap && data[l] > data[largest]) || (minHeap && data[l] < data[largest])) {
		largest = l
	}

	if r < n && ((!minHeap && data[r] > data[largest]) || (minHeap && data[r] < data[largest])) {
		largest = r
	}

	if largest != i {
		data[largest], data[i] = data[i], data[largest]
		return Heapify(data, n, largest, minHeap)
	}
	return data
}

func HeapifyBackward(data []int, i int, minHeap bool) []int {
	parent := (i - 1) / 2
	if parent >= 0 && ((!minHeap && data[i] > data[parent]) || (minHeap && data[i] < data[parent])) {
		data[parent], data[i] = data[i], data[parent]
		return HeapifyBackward(data, parent, minHeap)
	}
	return data
}

func HeapPop(data []int, minHeap bool) ([]int, int) {
	valueToPop := data[0]
	data[0], data[len(data)-1] = data[len(data)-1], data[0]
	data = data[:len(data)-1]
	data = Heapify(data, len(data), 0, minHeap)
	return data, valueToPop
}

func HeapInsert(data []int, num int, minHeap bool) []int {
	data = append(data, num)
	return HeapifyBackward(data, len(data)-1, minHeap)
}

type MedianFinder struct {
	minHeap []int
	maxHeap []int
}

func Constructor() MedianFinder {
	return MedianFinder{minHeap: []int{}, maxHeap: []int{}}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.minHeap) == 0 || (len(this.minHeap) != 0 && num > this.minHeap[0]) {
		this.minHeap = HeapInsert(this.minHeap, num, true)
	} else {
		this.maxHeap = HeapInsert(this.maxHeap, num, false)
	}

	diff := len(this.maxHeap) - len(this.minHeap)

	if diff > 1 {
		maxHeap := this.maxHeap
		maxHeap, value := HeapPop(maxHeap, false)
		this.maxHeap = maxHeap
		this.minHeap = HeapInsert(this.minHeap, value, true)
	} else if diff < -1 {
		minHeap := this.minHeap
		minHeap, value := HeapPop(minHeap, true)
		this.minHeap = minHeap
		this.maxHeap = HeapInsert(this.maxHeap, value, false)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	diff := len(this.maxHeap) - len(this.minHeap)
	switch diff {
	case 1:
		return float64(this.maxHeap[0])
	case -1:
		return float64(this.minHeap[0])
	default:
		return (float64(this.maxHeap[0]) + float64(this.minHeap[0])) / 2
	}
}

func runCommandsMedian(commands []string, args [][]int) []any {
	var obj MedianFinder
	results := []any{}
	for i := range commands {
		args := args[i]
		switch commands[i] {
		case "MedianFinder":
			obj = Constructor()
			results = append(results, nil)
		case "addNum":
			obj.AddNum(args[0])
			results = append(results, nil)
		case "findMedian":
			results = append(results, obj.FindMedian())
		}
	}
	return results
}

// func main() {

// 	fmt.Println(
// 		runCommandsMedian(
// 			[]string{"MedianFinder", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"},
// 			[][]int{{}, {1}, {}, {2}, {}, {3}, {}, {4}, {}, {5}, {}, {6}, {}, {7}, {}, {8}, {}, {9}, {}, {10}, {}},
// 		),
// 	)

// 	// fmt.Println(
// 	// 	runCommandsMedian(
// 	// 		[]string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"},
// 	// 		[][]int{{}, {1}, {2}, {}, {3}, {}},
// 	// 	),
// 	// )

// 	// fmt.Println(
// 	// 	runCommandsMedian(
// 	// 		[]string{"MedianFinder", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"},
// 	// 		[][]int{{}, {-1}, {}, {-2}, {}, {-3}, {}, {-4}, {}, {-5}, {}},
// 	// 	),
// 	// )
// }
