package main

type Task struct {
	left     int
	value    byte
	cooldown int
}

func TaskHeapify(data []Task, i int) {
	largest, l, r := i, i*2+1, i*2+2
	if l < len(data) && data[l].left > data[largest].left {
		largest = l
	}
	if r < len(data) && data[r].left > data[largest].left {
		largest = r
	}
	if largest != i {
		data[largest], data[i] = data[i], data[largest]
		TaskHeapify(data, largest)
	}
}

func TaskHeapifyBackward(data *[]Task, i int) {
	parent := (i - 1) / 2
	if parent >= 0 && (*data)[i].left > (*data)[parent].left {
		(*data)[i], (*data)[parent] = (*data)[parent], (*data)[i]
		TaskHeapifyBackward(data, parent)
	}
}

func TaskHeapBuild(data []Task) {
	for i := len(data)/2 - 1; i >= 0; i-- {
		TaskHeapify(data, i)
	}
}

func TaskHeapInsert(data *[]Task, value Task) {
	*data = append(*data, value)
	TaskHeapifyBackward(data, len(*data)-1)
}

func TaskHeapPop(data *[]Task) Task {
	s := *data
	valueToPop := s[0]
	s[0], s[len(s)-1] = s[len(s)-1], s[0]
	*data = s[:len(s)-1]
	TaskHeapify(*data, 0)
	return valueToPop
}

func TaskHeapRemove(data *[]Task, i int) {
	s := *data
	s[i], s[len(s)-1] = s[len(s)-1], s[i]
	*data = s[:len(s)-1]
	TaskHeapify(*data, i)
}

func leastInterval(tasks []byte, n int) int {
	tasksMap := map[byte]*Task{}
	for _, taskChar := range tasks {
		task, exists := tasksMap[taskChar]
		if !exists {
			tasksMap[taskChar] = &Task{left: 0, value: taskChar, cooldown: 0}
			task = tasksMap[taskChar]
		}
		task.left++
	}

	tasksHeap := []Task{}
	for _, task := range tasksMap {
		tasksHeap = append(tasksHeap, *task)
	}
	TaskHeapBuild(tasksHeap)

	queue := []Task{}
	time := 0
	for len(tasksHeap) != 0 || len(queue) != 0 {
		time++

		if len(tasksHeap) != 0 {
			current := TaskHeapPop(&tasksHeap)
			current.left--
			current.cooldown = n
			if current.left != 0 {
				queue = append(queue, current)
			}
		}

		if len(queue) != 0 && queue[0].cooldown == 0 {
			TaskHeapInsert(&tasksHeap, queue[0])
			queue = queue[1:]
		}

		for i := 0; i < len(queue); i++ {
			if queue[i].cooldown > 0 {
				queue[i].cooldown--
			}
		}
	}

	return time
}

// func main() {
// 	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))                               // 8
// 	fmt.Println(leastInterval([]byte{'A', 'C', 'A', 'B', 'D', 'B'}, 1))                               // 6
// 	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 3))                               // 10
// 	fmt.Println(leastInterval([]byte{'X', 'X', 'Y', 'Y'}, 2))                                         // 5
// 	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'C'}, 3))                                    // 9
// 	fmt.Println(leastInterval([]byte{'A', 'B', 'C', 'D', 'E', 'A', 'B', 'C', 'D', 'E'}, 4))           // 10
// 	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C', 'D', 'D', 'E'}, 2)) // 12
// }
