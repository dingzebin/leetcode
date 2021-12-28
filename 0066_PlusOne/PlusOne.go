package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{1, 2, 9}))
	fmt.Println(plusOne([]int{1, 9, 9}))
	fmt.Println(plusOne([]int{9, 9, 9}))
	fmt.Println(plusOne([]int{9}))
}

func plusOne(digits []int) []int {
	result := make([]int, len(digits)+1)

	plus := 1
	for i := len(digits) - 1; i >= 0; i-- {
		num := digits[i] + plus
		plus = num / 10
		result[i+1] = num % 10
	}
	if plus > 0 {
		result[0] = plus
		return result
	}
	return result[1:]
}
