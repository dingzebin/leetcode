package main

import "fmt"

/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
The overall run time complexity should be O(log (m+n)).
*/
func main() {
	fmt.Print(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)/2, 0, 0
	for low <= high {
		nums1Mid = low + (high-low)/2
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] {
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] {
			low = nums1Mid + 1
		} else {
			break
		}
	}
	maxLeft := 0
	if nums1Mid == 0 {
		maxLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		maxLeft = nums1[nums1Mid-1]
	} else {
		maxLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}
	if (len(nums1)+len(nums2))%2 == 1 {
		return float64(maxLeft)
	}

	minRight := 0
	if nums1Mid == len(nums1) {
		minRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		minRight = nums1[nums1Mid]
	} else {
		minRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}
	return (float64(maxLeft) + float64(minRight)) / 2
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
