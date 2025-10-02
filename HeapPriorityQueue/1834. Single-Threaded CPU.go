package main

import (
	"container/heap"
	"sort"
)

type ProcTask struct {
	EnqueueTime int
	ProcessTime int
	Index       int
}

type IntHeap []ProcTask

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool {
	if h[i].ProcessTime == h[j].ProcessTime {
		return h[i].Index < h[j].Index
	}
	return h[i].ProcessTime < h[j].ProcessTime
}
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(ProcTask))
}

func (h *IntHeap) Pop() interface{} {
	l := len(*h)
	res := (*h)[l-1]
	*h = (*h)[:l-1]
	return res
}

func getOrder(tasks [][]int) []int {
	heapData := &IntHeap{}
	heap.Init(heapData)
	procTasks := make([]ProcTask, len(tasks))

	for i := 0; i < len(tasks); i++ {
		procTasks[i] = ProcTask{
			EnqueueTime: tasks[i][0],
			ProcessTime: tasks[i][1],
			Index:       i,
		}
	}

	sort.Slice(procTasks, func(i int, j int) bool {
		return procTasks[i].EnqueueTime < procTasks[j].EnqueueTime
	})

	time, i, res := procTasks[0].EnqueueTime, 0, []int{}
	for heapData.Len() > 0 || i < len(procTasks) {
		for i < len(procTasks) && procTasks[i].EnqueueTime <= time {
			heap.Push(heapData, procTasks[i])
			i++
		}

		if heapData.Len() == 0 && i < len(procTasks) {
			time = procTasks[i].EnqueueTime
			continue
		}

		lastTask := heap.Pop(heapData).(ProcTask)
		res = append(res, lastTask.Index)
		time += lastTask.ProcessTime
	}
	return res
}

// func main() {
// 	fmt.Println(getOrder([][]int{{19, 13}, {16, 9}, {21, 10}, {32, 25}, {37, 4}, {49, 24}, {2, 15}, {38, 41}, {37, 34}, {33, 6}, {45, 4}, {18, 18}, {46, 39}, {12, 24}}))
// 	fmt.Println(getOrder([][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}}))
// 	fmt.Println(getOrder([][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}}))
// }
