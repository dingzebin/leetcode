package main

func main() {

}

func combine(n int, k int) [][]int {
	res := [][]int{}
	backtracking(n, k, 1, []int{}, &res)
	return res
}

func backtracking(n, k, start int, nums []int, res *[][]int) {
	if k == 0 {
		*res = append(*res, nums)
	} else {
		for i := start; i <= n-k+1; i++ {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			tmp = append(tmp, i)
			backtracking(n, k-1, i+1, tmp, res)
		}
	}
}
