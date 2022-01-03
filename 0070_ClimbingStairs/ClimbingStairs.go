package main

import "fmt"

func main() {
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
	fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {
	dp := []int{1, 1}
	for i := 2; i <= n; i++ {
		dp[i%2] = dp[0] + dp[1]
	}
	return dp[n%2]
}
