package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}
	if n == 0 {
		return
	}
	l, r, i := m-1, n-1, m+n-1
	for ; i >= 0; i-- {
		if l == -1 || r == -1 {
			break
		}
		if nums1[l] > nums2[r] {
			nums1[i] = nums1[l]
			l--
		} else {
			nums1[i] = nums2[r]
			r--
		}
	}
	for l >= 0 && i >= 0 {
		nums1[i] = nums1[l]
		l--
		i--
	}
	for r >= 0 && i >= 0 {
		nums1[i] = nums2[r]
		l--
		i--
	}
}

// func main() {
// 	// merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
// 	// merge([]int{1}, 1, []int{}, 0)
// 	// merge([]int{}, 0, []int{1}, 1)
// 	merge([]int{2, 0}, 1, []int{1}, 1)
// 	fmt.Println("Good")
// }
