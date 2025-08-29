package main

import (
	"fmt"
	"slices"
)

type HeapData struct {
	size int
	end  int
}

func MinHeapify(data *[]HeapData, n int, i int) {
	smallest, l, r := i, i*2+1, i*2+2
	if l < n && (*data)[l].size < (*data)[smallest].size {
		smallest = l
	}
	if r < n && (*data)[r].size < (*data)[smallest].size {
		smallest = r
	}
	if smallest != i {
		(*data)[smallest], (*data)[i] = (*data)[i], (*data)[smallest]
		MinHeapify(data, n, smallest)
	}
}

func MinHeapPop(data *[]HeapData) HeapData {
	s := *data
	valueToPop := s[0]
	s[0], s[len(s)-1] = s[len(s)-1], s[0]
	*data = s[:len(s)-1]
	MinHeapify(data, len(s)-1, 0)
	return valueToPop
}

func MinHeapifyBackward(data []HeapData, i int) {
	parent := (i - 1) / 2
	if parent >= 0 && data[i].size < data[parent].size {
		data[i], data[parent] = data[parent], data[i]
		MinHeapifyBackward(data, parent)
	}
}

func MinHeapBuild(data []HeapData) {
	for i := len(data)/2 - 1; i >= 0; i-- {
		MinHeapify(&data, len(data), i)
	}
}

func MinHeapInsert(data *[]HeapData, value HeapData) {
	*data = append(*data, value)
	MinHeapifyBackward(*data, len(*data)-1)
}

func compareInterval(a, b []int) int {
	if a[0] > b[0] {
		return 1
	} else if a[0] == b[0] {
		return 0
	}
	return -1
}

type Pair struct {
	value int
	index int
}

func compareQuery(a, b Pair) int {
	if a.value > b.value {
		return 1
	} else if a.value == b.value {
		return 0
	}
	return -1
}

func minInterval(intervals [][]int, queries []int) []int {
	queryPairs := []Pair{}
	for i, query := range queries {
		queryPairs = append(queryPairs, Pair{value: query, index: i})
	}
	queriesSorted := slices.SortedFunc(slices.Values(queryPairs), compareQuery)
	intervalsSorted := slices.SortedFunc(slices.Values(intervals), compareInterval)

	heap, i := []HeapData{}, 0
	for _, query := range queriesSorted {
		for i < len(intervalsSorted) && intervalsSorted[i][0] <= query.value {
			MinHeapInsert(&heap, HeapData{size: intervalsSorted[i][1] - intervalsSorted[i][0] + 1, end: intervalsSorted[i][1]})
			i++
		}

		for len(heap) > 0 && heap[0].end < query.value {
			MinHeapPop(&heap)
		}

		queries[query.index] = -1
		if len(heap) > 0 {
			queries[query.index] = heap[0].size
		}
	}

	return queries
}

func main() {
	// fmt.Println(minInterval([][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}}, []int{2, 3, 4, 5}))
	fmt.Println(minInterval([][]int{{4, 5}, {5, 8}, {1, 9}, {8, 10}, {1, 6}}, []int{7, 9, 3, 9, 3}))
}
