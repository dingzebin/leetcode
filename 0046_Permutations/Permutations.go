package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	res := [][]int{}
	backtracking(nums, &res, []int{})
	return res
}

func backtracking(nums []int, res *[][]int, items []int) {
	if len(items) == len(nums) {
		*res = append(*res, items)
	}
	for i := 0; i < len(nums); i++ {
		if !checkIsExit(nums[i], items) {
			tmp := make([]int, len(items))
			copy(tmp, items)
			tmp = append(tmp, nums[i])
			backtracking(nums, res, tmp)
		}
	}
}

func checkIsExit(i int, items []int) bool {
	for _, n := range items {
		if n == i {
			return true
		}
	}
	return false
}
