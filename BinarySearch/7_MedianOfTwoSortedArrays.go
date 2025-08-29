package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums2) < len(nums1) {
		nums1, nums2 = nums2, nums1
	}

	total := len(nums1) + len(nums2)
	half := (total + 1) / 2
	left, right := 0, len(nums1)
	odd := total%2 == 1

	for {
		i := (left + right) >> 1 // nums1
		j := half - i            // nums2

		nums1Left, nums1Right := math.MinInt, math.MaxInt
		nums2Left, nums2Right := math.MinInt, math.MaxInt

		if i < len(nums1) {
			nums1Right = nums1[i]
		}
		if j < len(nums2) {
			nums2Right = nums2[j]
		}
		if i-1 >= 0 {
			nums1Left = nums1[i-1]
		}
		if j-1 >= 0 {
			nums2Left = nums2[j-1]
		}

		numMax, numMin := max(nums1Left, nums2Left), min(nums1Right, nums2Right)

		if nums1Left <= nums2Right && nums2Left <= nums1Right {
			if odd {
				return float64(numMax)
			}
			return (float64(numMax) + float64(numMin)) / 2
		} else if nums1Left > nums2Right {
			right = i - 1
		} else {
			left = i + 1
		}
	}
}
