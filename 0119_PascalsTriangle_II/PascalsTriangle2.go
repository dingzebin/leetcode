package main

import "fmt"

func main() {
	fmt.Println(getRow(3))
	fmt.Println(getRow(0))
	fmt.Println(getRow(1))
}

func getRow(rowIndex int) []int {
	rows := make([][]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1
		for j := 1; j < i; j++ {
			row[j] = rows[i-1][j-1] + rows[i-1][j]
		}
		rows[i] = row
	}
	return rows[rowIndex]
}
