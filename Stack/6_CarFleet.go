package main

type CRStackElement struct {
	Value int
	Next  *CRStackElement
}

type CRStack struct {
	Head *CRStackElement
}

func (this *CRStack) Push(val int) {
	oldHead := this.Head
	this.Head = &CRStackElement{Value: val, Next: oldHead}
}

func (this *CRStack) Pop() {
	newHead := this.Head.Next
	this.Head = newHead
}

func (this *CRStack) Peek() int {
	return this.Head.Value
}

func isFleet(t float32, p1 float32, s1 float32, p2 float32, s2 float32) bool {
	if s1 == s2 {
		return false
	}
	x := (p2 - p1) / (s1 - s2)
	y := (s1*x + p1)
	return x > 0 && y > 0 && y <= t
}

// QuickSort sorts the first array and adjusts the second array accordingly
func quickSort(arr []int, values []int, low, high int) {
	if low < high {
		p := partition(arr, values, low, high)
		quickSort(arr, values, low, p-1)
		quickSort(arr, values, p+1, high)
	}
}

// Partition function for QuickSort
func partition(arr []int, values []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] > pivot { // Sorting in ascending order
			arr[i], arr[j] = arr[j], arr[i]
			values[i], values[j] = values[j], values[i] // Swap in second array
			i++
		}
	}
	// Swap pivot into correct position
	arr[i], arr[high] = arr[high], arr[i]
	values[i], values[high] = values[high], values[i]
	return i
}

func carFleet(target int, position []int, speed []int) int {
	fleets := 0
	t := float32(target)
	quickSort(position, speed, 0, len(position)-1)

	for len(position) > 0 {
		i := 0
		stackToDelete := CRStack{Head: &CRStackElement{Value: i}}

		p1 := float32(position[i])
		s1 := float32(speed[i])

		for j := 1; j < len(position); j++ {
			p2 := float32(position[j])
			s2 := float32(speed[j])

			if isFleet(t, p1, s1, p2, s2) {
				if p2 > p1 {
					p1 = p2
				}
				if s1 > s2 {
					s1 = s2
				}
				stackToDelete.Push(j)
			}

		}

		fleets++

		for stackToDelete.Head != nil {
			i := stackToDelete.Peek()
			currentLen := len(position)
			if i+1 >= currentLen {
				position = position[:i]
				speed = speed[:i]
			} else {
				position = append(position[:i], position[i+1:]...)
				speed = append(speed[:i], speed[i+1:]...)
			}
			stackToDelete.Pop()
		}
	}
	return fleets
}

// func main() {
// 	// fmt.Println(carFleet( //3
// 	// 	12,
// 	// 	[]int{10, 8, 0, 5, 3},
// 	// 	[]int{2, 4, 1, 1, 3},
// 	// ))
// 	// fmt.Println(carFleet( //1
// 	// 	10,
// 	// 	[]int{3},
// 	// 	[]int{3},
// 	// ))
// 	// fmt.Println(carFleet( //1
// 	// 	100,
// 	// 	[]int{0, 2, 4},
// 	// 	[]int{4, 2, 1},
// 	// ))
// 	// fmt.Println(carFleet( //1
// 	// 	10,
// 	// 	[]int{0, 4, 2},
// 	// 	[]int{2, 1, 3},
// 	// ))
// 	// fmt.Println(carFleet( //2
// 	// 	20,
// 	// 	[]int{6, 2, 17},
// 	// 	[]int{3, 9, 2},
// 	// ))
// 	// fmt.Println(carFleet( //4
// 	// 	12,
// 	// 	[]int{4, 0, 5, 3, 1, 2},
// 	// 	[]int{6, 10, 9, 6, 7, 2},
// 	// ))
// 	// fmt.Println(carFleet( //2
// 	// 	16,
// 	// 	[]int{11, 14, 13, 6},
// 	// 	[]int{2, 2, 6, 7},
// 	// ))
// 	// fmt.Println(carFleet( //2
// 	// 	13,
// 	// 	[]int{10, 2, 5, 7, 4, 6, 11},
// 	// 	[]int{7, 5, 10, 5, 9, 4, 1},
// 	// ))
// }
