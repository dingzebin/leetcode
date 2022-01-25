package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	backtracking(candidates, []int{}, &res, target)
	return res
}

func backtracking(candidates []int, nums []int, res *[][]int, target int) {
	if target < 0 {
		return
	} else if target == 0 {
		*res = append(*res, nums)
	} else {
		for i := 0; i < len(candidates); i++ {
			n := make([]int, len(nums))
			copy(n, nums)
			n = append(n, candidates[i])
			backtracking(candidates[i:], n, res, target-candidates[i])
		}
	}
}
