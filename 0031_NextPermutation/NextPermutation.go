package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	nextPermutation(a)
	fmt.Println(a)

	a = []int{2, 1, 3}
	nextPermutation(a)
	fmt.Println(a)

	a = []int{3, 1, 2}
	nextPermutation(a)
	fmt.Println(a)

	a = []int{3, 2, 1}
	nextPermutation(a)
	fmt.Println(a)

	a = []int{1, 3, 2}
	nextPermutation(a)
	fmt.Println(a)

}

func nextPermutation(nums []int) {
	i := len(nums) - 1
	for ; i > 0; i-- {
		if nums[i] > nums[i-1] {
			break
		}
	}
	if i == 0 {
		revert(nums)
	} else {
		i = i - 1
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
				break
			}
		}
		revert(nums[i+1:])
	}
}

func revert(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
