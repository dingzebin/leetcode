package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{0, 1, 2}, 3))
}

func threeSumClosest(nums []int, target int) int {
	r := nums[0] + nums[1] + nums[len(nums)-1]
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		start, end := i+1, len(nums)-1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if sum > target {
				end--
			} else if sum < target {
				start++
			} else {
				return sum
			}
			if abs(sum-target) < abs(r-target) {
				r = sum
			}
		}
	}
	return r
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func threeSumClosestBF(nums []int, target int) int {
	r := 0
	ab := -1
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				n := nums[i] + nums[j] + nums[k]
				tmp := getAb(n, target)
				if ab == -1 || tmp < ab {
					ab, r = tmp, n
				}
			}
		}
	}
	return r
}

func getAb(n, target int) int {
	if n > target {
		return n - target
	} else {
		return target - n
	}
}
