package main

import (
	"fmt"
)

func main() {
	fmt.Println(exist([][]byte{{"A"[0], "B"[0], "C"[0], "E"[0]}, {"S"[0], "F"[0], "C"[0], "S"[0]}, {"A"[0], "D"[0], "E"[0], "E"[0]}}, "ABCB"))
	fmt.Println(exist([][]byte{{"A"[0], "B"[0], "C"[0], "E"[0]}, {"S"[0], "F"[0], "C"[0], "S"[0]}, {"A"[0], "D"[0], "E"[0], "E"[0]}}, "ABCCED"))
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if search(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}
func search(board [][]byte, word string, i, j, idx int) bool {
	if idx == len(word) {
		return true
	}
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || board[i][j] != word[idx] {
		return false
	}
	board[i][j] ^= 255
	exist := search(board, word, i-1, j, idx+1) || search(board, word, i+1, j, idx+1) || search(board, word, i, j-1, idx+1) || search(board, word, i, j+1, idx+1)

	board[i][j] ^= 255

	return exist
}
