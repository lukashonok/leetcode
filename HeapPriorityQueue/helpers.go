package main

type HeapData struct {
	key   float32
	value any
}

func MaxHeapify(data []HeapData, i int) {
	largest, l, r := i, i*2+1, i*2+2
	if l < len(data) && data[l].key > data[largest].key {
		largest = l
	}
	if r < len(data) && data[r].key > data[largest].key {
		largest = r
	}
	if largest != i {
		data[largest], data[i] = data[i], data[largest]
		MaxHeapify(data, largest)
	}
}

func MaxHeapifyBackward(data []HeapData, i int) {
	parent := (i - 1) / 2
	if parent >= 0 && data[i].key > data[parent].key {
		data[i], data[parent] = data[parent], data[i]
		MaxHeapifyBackward(data, parent)
	}
}

func MaxHeapBuild(data []HeapData) {
	for i := len(data)/2 - 1; i >= 0; i-- {
		MaxHeapify(data, i)
	}
}

func MaxHeapInsert(data *[]HeapData, value HeapData) {
	*data = append(*data, value)
	MaxHeapifyBackward(*data, len(*data)-1)
}

func MaxHeapPop(data *[]HeapData) HeapData {
	s := *data
	valueToPop := s[0]
	s[0], s[len(s)-1] = s[len(s)-1], s[0]
	*data = s[:len(s)-1]
	MaxHeapify(*data, 0)
	return valueToPop
}

func MaxHeapSort(data *[]HeapData) {
	for i := len(*data) - 1; i > 0; i-- {
		(*data)[0], (*data)[i] = (*data)[i], (*data)[0]
		MaxHeapify(*data, i)
	}
}

//

func MinHeapify(data *[]HeapData, n int, i int) {
	smallest, l, r := i, i*2+1, i*2+2
	if l < n && (*data)[l].key < (*data)[smallest].key {
		smallest = l
	}
	if r < n && (*data)[r].key < (*data)[smallest].key {
		smallest = r
	}
	if smallest != i {
		(*data)[smallest], (*data)[i] = (*data)[i], (*data)[smallest]
		MinHeapify(data, n, smallest)
	}
}

func MinHeapifyBackward(data []HeapData, i int) {
	parent := (i - 1) / 2
	if parent >= 0 && data[i].key < data[parent].key {
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

func MinHeapPop(data *[]HeapData) HeapData {
	s := *data
	valueToPop := s[0]
	s[0], s[len(s)-1] = s[len(s)-1], s[0]
	*data = s[:len(s)-1]
	MinHeapify(data, len(s)-1, 0)
	return valueToPop
}

func MinHeapSort(data *[]HeapData) {
	for i := len(*data) - 1; i >= 0; i-- {
		(*data)[0], (*data)[i] = (*data)[i], (*data)[0]
		MinHeapify(data, i, 0)
	}
}
