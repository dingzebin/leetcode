package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0}))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	resultMap, num := make(map[[3]int][]int), 0
	for i := 0; i < len(nums)-2; i++ {
		set := make(map[int]interface{})
		target := 0 - nums[i]
		for j := i + 1; j < len(nums); j++ {
			num = target - nums[j]
			if _, ok := set[num]; ok {
				resultMap[[3]int{nums[i], num, nums[j]}] = nil
			} else {
				set[nums[j]] = nil
			}
		}
	}
	r := make([][]int, len(resultMap))
	i := 0
	for key := range resultMap {
		r[i] = []int{key[0], key[1], key[2]}
		i++
	}
	return r
}
