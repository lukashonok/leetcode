package main

import "fmt"

func check(got, want any, desc string) {
	if got != want {
		panic(fmt.Sprintf("FAIL: %s → got %v, want %v", desc, got, want))
	}
}

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	m := map[int]bool{}
	for i := range len(nums) {
		_, exist := m[nums[i]]
		if exist {
			return true
		}
		m[nums[i]] = false
	}

	return false
}
