package main

import "fmt"

func main() {
	fmt.Println(jump([]int{2, 3, 1, 4}))
	fmt.Println(jump([]int{2, 0, 1, 1, 4}))
	fmt.Println(jump([]int{5, 6, 4, 4, 6, 9, 4, 4, 7, 4, 4, 8, 2, 6, 8, 1, 5, 9, 6, 5, 2, 7, 9, 7, 9, 6, 9, 4, 1, 6, 8, 8, 4, 4, 2, 0, 3, 8, 5}))

}

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	step, curEnd, curFarthest := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > curFarthest {
			curFarthest = i + nums[i]
			if curFarthest >= len(nums)-1 {
				return step + 1
			}
		}
		if i == curEnd {
			curEnd = curFarthest
			step++
		}
	}
	return step
}
