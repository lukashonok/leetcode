package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type CarTrip struct {
	People int
	From   int
	To     int
}

type TripHeap []CarTrip

func (h TripHeap) Len() int { return len(h) }
func (h TripHeap) Less(i, j int) bool {
	return h[i].To < h[j].To
}
func (h TripHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *TripHeap) Push(x interface{}) {
	*h = append(*h, x.(CarTrip))
}

func (h *TripHeap) Pop() interface{} {
	l := len(*h)
	res := (*h)[l-1]
	*h = (*h)[:l-1]
	return res
}

func carPooling(trips [][]int, capacity int) bool {
	carTrips := make([]CarTrip, len(trips))
	sort.Slice(trips, func(i, j int) bool {
		return trips[i][1] < trips[j][1]
	})
	for i, t := range trips {
		carTrips[i] = CarTrip{t[0], t[1], t[2]}
	}
	tripHeap := &TripHeap{}
	heap.Init(tripHeap)

	currentPeople := 0
	for _, trip := range carTrips {
		for tripHeap.Len() > 0 && (*tripHeap)[0].To <= trip.From {
			lastTrip := heap.Pop(tripHeap).(CarTrip)
			currentPeople -= lastTrip.People
		}
		if currentPeople+trip.People > capacity {
			return false
		}
		currentPeople += trip.People
		heap.Push(tripHeap, trip)
	}

	return true
}

func main() {
	fmt.Println(carPooling([][]int{
		{2, 1, 5},
		{3, 3, 7},
	}, 4)) // false
	fmt.Println(carPooling([][]int{
		{2, 1, 5},
		{3, 3, 7},
	}, 5)) // true
	fmt.Println(carPooling([][]int{
		{2, 7, 11},
		{3, 3, 7},
		{3, 5, 9},
		{3, 1, 5},
	}, 6)) // true
	fmt.Println(carPooling([][]int{
		{9, 3, 4},
		{9, 1, 7},
		{4, 2, 4},
		{7, 4, 5},
	}, 23)) // true
}
