package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(1))
}

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	rowBegin, rowEnd, colBegin, colEnd := 0, n-1, 0, n-1
	num := 1
	for rowBegin <= rowEnd && colBegin <= colEnd {
		for i := colBegin; i <= colEnd; i++ {
			matrix[rowBegin][i] = num
			num++
		}
		rowBegin++
		for i := rowBegin; i <= rowEnd; i++ {
			matrix[i][colEnd] = num
			num++
		}
		colEnd--
		for i := colEnd; i >= colBegin; i-- {
			matrix[rowEnd][i] = num
			num++
		}
		rowEnd--
		for i := rowEnd; i >= rowBegin; i-- {
			matrix[i][colBegin] = num
			num++
		}
		colBegin++
	}
	return matrix
}
