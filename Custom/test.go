package main

import "fmt"

func quicksort(a []int) []int {
	length := len(a)
	if length < 2 {
		return a
	}

	left, right, pivot := 0, length-1, length/2

	a[right], a[pivot] = a[pivot], a[right]
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

func main() {
	a := []int{6, 2, 1, 0, 20}
	b := quicksort(a)
	fmt.Println(a, b)
}
