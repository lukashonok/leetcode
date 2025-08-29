package main

import "fmt"

func majorityElement(nums []int) int {
	frequency, length, threshold := map[int]int{}, len(nums), len(nums)/2
	for i := 0; i < length; i++ {
		frequency[nums[i]]++
		if frequency[nums[i]] > threshold {
			return nums[i]
		}
	}

	return 0
}

func quicksort(a []int) []int {
	length := len(a)
	if length < 2 {
		return a
	}

	left, right, pivot := 0, length-1, length/2
	a[right], a[pivot] = a[pivot], a[right]
	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[right], a[left] = a[left], a[right]

	quicksort(a[:left])
	quicksort(a[left+1:])
	return a
}

func main() { fmt.Println(quicksort([]int{4, 2, 45, 56, 7, 23, 21})) }
