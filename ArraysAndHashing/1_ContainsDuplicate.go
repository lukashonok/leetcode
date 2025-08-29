package main

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
