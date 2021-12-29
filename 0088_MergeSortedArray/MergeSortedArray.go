package main

import "fmt"

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge([]int{1}, 1, []int{}, 0)
	merge([]int{0}, 0, []int{1}, 1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i1, i2 := m-1, n-1
	for k := len(nums1) - 1; k >= 0; k-- {
		if i2 < 0 || (i1 >= 0 && nums1[i1] > nums2[i2]) {
			nums1[k] = nums1[i1]
			i1 = i1 - 1
		} else {
			nums1[k] = nums2[i2]
			i2 = i2 - 1
		}
	}
	fmt.Println(nums1)
}
