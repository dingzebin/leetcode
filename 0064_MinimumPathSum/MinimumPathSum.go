package main

func main() {

}

func minPathSum(grid [][]int) int {
	h, w := len(grid), len(grid[0])
	for i := 1; i < h; i++ {
		grid[i][0] += grid[i-1][0]
	}

	for i := 1; i < w; i++ {
		grid[0][i] += grid[0][i-1]
	}

	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[h-1][w-1]
}

func minPathSum2(grid [][]int) int {
	h, w := len(grid), len(grid[0])
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[h-1][w-1]
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
