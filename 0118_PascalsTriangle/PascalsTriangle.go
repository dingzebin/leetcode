package main

import "fmt"

func main() {
	fmt.Println(generate(5))
	fmt.Println(generate(1))
	fmt.Println(generate(2))
}

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1
		for j := 1; j < i; j++ {
			row[j] = result[i-1][j-1] + result[i-1][j]
		}
		result[i] = row
	}
	return result
}
