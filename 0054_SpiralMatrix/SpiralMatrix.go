package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

func spiralOrder(matrix [][]int) []int {
	heigh, width := len(matrix)-1, len(matrix[0])-1
	sumOfMatrix := len(matrix) * len(matrix[0])
	res := make([]int, sumOfMatrix)

	// 0 right 1 down 2 left 3 up
	direction, n := 0, 0
	row, col := 0, 0
	for i := 0; i < sumOfMatrix; i++ {
		res[i] = matrix[row][col]
		if 0 == direction {
			if col == width-n {
				if row == heigh-n {
					break
				} else {
					direction = 1
					row++
				}
			} else {
				col++
			}
		} else if 1 == direction {
			if row == heigh-n {
				if col == n {
					break
				} else {
					direction = 2
					col--
				}
			} else {
				row++
			}
		} else if 2 == direction {
			if col == n {
				if row == n {
					break
				} else {
					direction = 3
					row--
					n++
				}
			} else {
				col--
			}
		} else if 3 == direction {
			if row == n {
				if col == width-n {
					break
				} else {
					direction = 0
					col++
				}
			} else {
				row--
			}
		}
	}
	return res
}
