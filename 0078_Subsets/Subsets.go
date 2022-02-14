package main

import "sort"

func main() {

}

func subsets(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	backtracking(nums, []int{}, &res, 0)
	return res
}
func backtracking(nums, sub []int, res *[][]int, start int) {
	*res = append(*res, sub)
	for i := start; i < len(nums); i++ {
		tmp := make([]int, len(sub))
		copy(tmp, sub)
		tmp = append(tmp, nums[i])
		backtracking(nums, tmp, res, i+1)
	}
}
