package main

func main() {

}

func rotate(matrix [][]int) {
	for start, end := 0, len(matrix)-1; start < end; {
		matrix[start], matrix[end] = matrix[end], matrix[start]
		start++
		end--
	}
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
