package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
	fmt.Println(fourSum([]int{5, 5, 3, 5, 1, -5, 1, -2}, 4))
	fmt.Println(fourSum([]int{-492, -479, -468, -447, -432, -428, -418, -406, -388, -369, -300, -275, -238, -231, -228, -225, -224, -221, -220, -219, -189, -186, -180, -144, -130, -73, -67, -66, -55, -54, -53, -19, -6, 13, 28, 36, 47, 57, 80, 82, 87, 97, 97, 120, 132, 142, 148, 174, 176, 176, 205, 225, 232, 238, 245, 247, 264, 268, 268, 275, 319, 334, 387, 392, 412, 413, 426, 434, 442, 451, 475, 478, 485, 490},
		4631))

}

func fourSum(nums []int, target int) (result [][]int) {
	if len(nums) < 4 {
		return result
	}
	sort.Ints(nums)
	l := len(nums)
	if nums[l-1]+nums[l-2]+nums[l-3]+nums[l-4] < target {
		return result
	}
	for i := 0; i < l-3; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[l-1]+nums[l-2]+nums[l-3] < target {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[l-1]+nums[l-2] < target {
				continue
			}
			left, right := j+1, len(nums)-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum > target {
					right--
				} else if sum < target {
					left++
				} else {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				}
			}
		}
	}
	return result
}
