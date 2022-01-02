package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}

func maxArea(height []int) int {
	start, end, maxArea, h, width := 0, len(height)-1, 0, 0, 0
	for start < end {
		width = end - start
		if height[start] < height[end] {
			h = height[start]
			start++
		} else {
			h = height[end]
			end--
		}
		area := h * width
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
