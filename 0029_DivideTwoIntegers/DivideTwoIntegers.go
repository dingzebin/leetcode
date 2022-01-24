package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(divide(10, 3))
	// fmt.Println(divide(7, -3))
	fmt.Println(divide(22, 3))
	// fmt.Println(divide(-2147483648, -1))
	// fmt.Println(divide(-2147483648, 1))
}

func divide(dividend int, divisor int) int {
	s := 1
	if dividend^divisor < 0 {
		s = -1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	r := 0
	for i := 31; i >= 0; i-- {
		if (dividend >> i) >= divisor {
			r += 1 << i
			dividend -= divisor << i
		}
	}
	r = r * s
	if r > math.MaxInt32 {
		return math.MaxInt32
	}
	return r
}
