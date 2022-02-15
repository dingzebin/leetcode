package main

import (
	"fmt"
)

func main() {
	fmt.Println(exist([][]byte{{"A"[0], "B"[0], "C"[0], "E"[0]}, {"S"[0], "F"[0], "C"[0], "S"[0]}, {"A"[0], "D"[0], "E"[0], "E"[0]}}, "ABCB"))
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if search(board, word, [2]int{i, j}, 0, [][2]int{}) {
				return true
			}
		}
	}
	return false
}
func search(board [][]byte, word string, cell [2]int, idx int, matched [][2]int) bool {
	if board[cell[0]][cell[1]] == word[idx] {
		if idx == len(word)-1 {
			return true
		}
		matched = append(matched, cell)
		if cell[0] < len(board)-1 && !hasMatched(matched, [2]int{cell[0] + 1, cell[1]}) {
			if search(board, word, [2]int{cell[0] + 1, cell[1]}, idx+1, matched) {
				return true
			}
		}
		if cell[0] > 0 && !hasMatched(matched, [2]int{cell[0] - 1, cell[1]}) {
			if search(board, word, [2]int{cell[0] - 1, cell[1]}, idx+1, matched) {
				return true
			}
		}
		if cell[1] > 0 && !hasMatched(matched, [2]int{cell[0], cell[1] - 1}) {
			if search(board, word, [2]int{cell[0], cell[1] - 1}, idx+1, matched) {
				return true
			}
		}
		if cell[1] < len(board[0])-1 && !hasMatched(matched, [2]int{cell[0], cell[1] + 1}) {
			if search(board, word, [2]int{cell[0], cell[1] + 1}, idx+1, matched) {
				return true
			}
		}
	}
	return false
}

func hasMatched(matched [][2]int, n [2]int) bool {
	for _, m := range matched {
		if m == n {
			return true
		}
	}
	return false
}
