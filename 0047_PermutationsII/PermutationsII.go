package main

import (
	"fmt"
)

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	backtracking(nums, &res, []int{})
	return res
}

func backtracking(nums []int, res *[][]int, items []int) {
	if len(items) == len(nums) {
		t := make([]int, len(items))
		for i, item := range items {
			t[i] = nums[item]
		}
		*res = append(*res, t)
	}
	filter := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if !checkIsExit(i, items) && !filter[nums[i]] {
			tmp := make([]int, len(items))
			copy(tmp, items)
			tmp = append(tmp, i)
			filter[nums[i]] = true
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
