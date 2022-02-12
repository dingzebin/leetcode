package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
}
func searchMatrix(matrix [][]int, target int) bool {
	low, high, w := 0, len(matrix)*len(matrix[0])-1, len(matrix[0])
	for low <= high {
		mid := low + (high-low)/2
		if matrix[mid/w][mid%w] == target {
			return true
		} else if matrix[mid/w][mid%w] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
