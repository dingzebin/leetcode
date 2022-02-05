package main

import "fmt"

func main() {
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	i := 0
	for ; i < len(intervals) && newInterval[0] > intervals[i][1]; i++ {
		res = append(res, intervals[i])
	}
	for ; i < len(intervals) && intervals[i][0] <= newInterval[1]; i++ {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
	}
	res = append(res, newInterval)

	for ; i < len(intervals); i++ {
		res = append(res, intervals[i])
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
