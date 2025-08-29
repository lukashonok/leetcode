package main

func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	postfix := make([]int, len(nums))

	// Initialising
	for i := range prefix {
		prefix[i] = 1
		postfix[i] = 1
	}

	// Prefix
	for i := 1; i < len(prefix); i++ {
		prefix[i] *= nums[i-1] * prefix[i-1]
	}

	// Postfix
	for i := len(postfix) - 2; i >= 0; i-- {
		postfix[i] *= nums[i+1] * postfix[i+1]
	}

	// Result
	for i := range prefix {
		postfix[i] *= prefix[i]
	}

	return postfix
}
