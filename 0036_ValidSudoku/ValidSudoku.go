package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValidSudoku([][]byte{
		{5, 3, 46, 46, 7, 46, 46, 46, 46},
		{6, 46, 46, 1, 9, 5, 46, 46, 46},
		{46, 9, 8, 46, 46, 46, 46, 6, 46},
		{8, 46, 46, 46, 6, 46, 46, 46, 3},
		{4, 46, 46, 8, 46, 3, 46, 46, 1},
		{7, 46, 46, 46, 2, 46, 46, 46, 6},
		{46, 6, 46, 46, 46, 46, 2, 8, 46},
		{46, 46, 46, 4, 1, 9, 46, 46, 5},
		{46, 46, 46, 46, 8, 46, 46, 7, 9}}))
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		cols := map[byte]bool{}
		rows := map[byte]bool{}
		cube := map[byte]bool{}
		rowIdx := 3 * (i / 3)
		colIdx := 3 * (i % 3)
		for j := 0; j < 9; j++ {
			if board[i][j] != 46 && rows[board[i][j]] {
				return false
			} else if board[j][i] != 46 && cols[board[j][i]] {
				return false
			} else if board[rowIdx+j%3][colIdx+j/3] != 46 && cube[board[rowIdx+j%3][colIdx+j/3]] {
				return false
			} else {
				rows[board[i][j]] = true
				cols[board[j][i]] = true
				cube[board[rowIdx+j%3][colIdx+j/3]] = true
			}

		}
	}
	return true
}

func isValidSudok2(board [][]byte) bool {
	cols := [9]map[byte]bool{}
	rows := [9]map[byte]bool{}
	tmp := [3][3]map[byte]bool{}
	for i := 0; i < 9; i++ {
		cols[i] = map[byte]bool{}
		rows[i] = map[byte]bool{}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tmp[i][j] = map[byte]bool{}
		}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			v := board[i][j]
			if v == 46 {
				continue
			}
			if _, ok := cols[j][v]; ok {
				return false
			} else if _, ok := rows[i][v]; ok {
				return false
			} else if _, ok = tmp[i/3][j/3][v]; ok {
				return false
			} else {
				cols[j][v] = true
				rows[i][v] = true
				tmp[i/3][j/3][v] = true
			}
		}
	}
	return true
}
