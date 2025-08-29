package main

type HeapData []int

func MinHeapify(data *[]HeapData, n int, i int) {
	smallest, l, r := i, i*2+1, i*2+2
	if l < n && (*data)[l][0] < (*data)[smallest][0] {
		smallest = l
	}
	if r < n && (*data)[r][0] < (*data)[smallest][0] {
		smallest = r
	}
	if smallest != i {
		(*data)[smallest], (*data)[i] = (*data)[i], (*data)[smallest]
		MinHeapify(data, n, smallest)
	}
}

func MinHeapifyBackward(data []HeapData, i int) {
	parent := (i - 1) / 2
	if parent >= 0 && data[i][0] < data[parent][0] {
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

func networkDelayTime(times [][]int, n int, k int) int {
	adj := map[int][][]int{}
	for _, time := range times {
		if adj[time[0]] == nil {
			adj[time[0]] = [][]int{}
		}
		adj[time[0]] = append(adj[time[0]], []int{time[2], time[1]})
	}

	heap, visited, result := []HeapData{{0, k}}, map[int]bool{}, 0
	for len(heap) != 0 {
		current := MinHeapPop(&heap)
		if visited[current[1]] {
			continue
		}
		visited[current[1]] = true
		result = max(current[0], result)
		for _, v := range adj[current[1]] {
			if visited[v[1]] {
				continue
			}

			MinHeapInsert(&heap, []int{v[0] + current[0], v[1]})
		}
	}

	if len(visited) != n {
		return -1
	}

	return result
}

// func main() {
// 	fmt.Println(networkDelayTime([][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 2}}, 3, 1))
// 	fmt.Println(networkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2))
// 	fmt.Println(networkDelayTime([][]int{{1, 2, 1}}, 2, 1))
// 	fmt.Println(networkDelayTime([][]int{{1, 2, 1}}, 2, 2))
// }
