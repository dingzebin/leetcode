package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{0}))
	fmt.Println(canJump([]int{2, 5, 0, 0}))
	fmt.Println(canJump([]int{0, 1}))
	fmt.Println(canJump([]int{0, 2, 3}))
}

func canJump(nums []int) bool {
	i := 0
	for reach := 0; i < len(nums) && i <= reach; i++ {
		if i+nums[i] > reach {
			reach = i + nums[i]
		}
	}
	return i == len(nums)
}
