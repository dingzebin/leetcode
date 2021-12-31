package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}

func maxProfit(prices []int) int {
	maxProfit := 0
	min := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		profit = prices[i] - min
		if profit > maxProfit {
			maxProfit = profit
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return maxProfit
}
