package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
	fmt.Println(merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}))
}

type Interval [][]int

func (s Interval) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func (s Interval) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Interval) Len() int {
	return len(s)
}

func merge(intervals [][]int) [][]int {
	sort.Sort(Interval(intervals))
	interval, res := intervals[0], [][]int{}
	for _, tmp := range intervals {
		if tmp[0] <= interval[1] {
			if tmp[1] > interval[1] {
				interval[1] = tmp[1]
			}
		} else {
			res = append(res, interval)
			interval = tmp
		}
	}
	res = append(res, interval)
	return res
}
