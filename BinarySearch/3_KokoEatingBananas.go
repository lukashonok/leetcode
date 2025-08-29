package main

import "slices"

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, slices.Max(piles)
	mid := (left + right) / 2
	for left < right {
		actualHours := 0
		for _, v := range piles {
			actualHours += (v + mid - 1) / mid
		}
		if actualHours > h {
			left = mid + 1
		} else {
			right = mid
		}
		mid = (left + right) / 2
	}
	return mid
}
