package main

import "fmt"

func main() {
	m := [][]int{{1, 2, 3, 4}, {5, 0, 7, 8}, {0, 10, 11, 12}, {13, 14, 15, 0}}
	setZeroes(m)
	fmt.Println(m)

}

func setZeroes(matrix [][]int) {
	col0, l, w := false, len(matrix), len(matrix[0])
	for i := 0; i < l; i++ {
		if matrix[i][0] == 0 {
			col0 = true
		}
		for j := 1; j < w; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := l - 1; i >= 0; i-- {
		for j := w - 1; j > 0; j-- {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}

func setZeroes2(matrix [][]int) {
	l, w := len(matrix), len(matrix[0])
	isFirstRowExistZero, isFirstColExistZero := false, false
	for i := 0; i < l; i++ {
		if matrix[i][0] == 0 {
			isFirstRowExistZero = true
		}
	}
	for i := 0; i < w; i++ {
		if matrix[0][i] == 0 {
			isFirstColExistZero = true
		}
	}

	for i := 0; i < l; i++ {
		for j := 0; j < w; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < l; i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < w; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := 1; i < w; i++ {
		if matrix[0][i] == 0 {
			for j := 0; j < l; j++ {
				matrix[j][i] = 0
			}
		}
	}
	if isFirstRowExistZero {
		for i := 0; i < l; i++ {
			matrix[i][0] = 0
		}
	}
	if isFirstColExistZero {
		for i := 0; i < w; i++ {
			matrix[0][i] = 0
		}
	}
}

func setZeroes1(matrix [][]int) {
	rows, cols := map[int]bool{}, map[int]bool{}
	l, w := len(matrix), len(matrix[0])
	for i := 0; i < l; i++ {
		for j := 0; j < w; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}
	for k := range rows {
		for i := 0; i < w; i++ {
			matrix[k][i] = 0
		}
	}
	for k := range cols {
		for i := 0; i < l; i++ {
			matrix[i][k] = 0
		}
	}
}
