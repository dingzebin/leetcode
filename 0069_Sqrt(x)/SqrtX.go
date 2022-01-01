package main

import (
	"fmt"
)

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
}

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	mid, low, high := 0, 2, x
	for low <= high {
		mid = low + (high-low)/2
		p := mid * mid
		if p == x {
			return mid
		} else if p > x || p < 0 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low - 1
}
