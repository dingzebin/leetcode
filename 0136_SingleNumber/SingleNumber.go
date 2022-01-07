package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{1}))
}

/*
x ^ y ^ y = x
*/
func singleNumber(nums []int) int {
	p := 0
	for i := 0; i < len(nums); i++ {
		p ^= nums[i]
	}
	return p
}
