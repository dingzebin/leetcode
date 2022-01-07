package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{1}))
}

func singleNumber(nums []int) int {
	p := 0
	for i := 0; i < len(nums); i++ {
		p ^= nums[i]
	}
	return p
}

func singleNumber2(nums []int) int {
	set := map[int]int{}
	for i := 0; i < len(nums); i++ {
		set[nums[i]] = set[nums[i]] + 1
	}
	for key, val := range set {
		if val == 1 {
			return key
		}
	}
	return 0
}
