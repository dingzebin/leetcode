package main

import (
	"fmt"
)

func main() {
	fmt.Println(myPow(0.00001, 2147483647))
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / pow(x, -n)
	} else {
		return pow(x, n)
	}
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	tmp := pow(x, n/2)
	if n%2 == 0 {
		return tmp * tmp
	} else {
		return x * tmp * tmp
	}
}
