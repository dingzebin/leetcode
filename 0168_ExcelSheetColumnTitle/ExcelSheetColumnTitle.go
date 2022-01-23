package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(701))
	fmt.Println(convertToTitle(52))
}

func convertToTitle(columnNumber int) string {
	if columnNumber <= 26 {
		return string([]byte{byte(columnNumber + 64)})
	} else {
		return convertToTitle((columnNumber-1)/26) + convertToTitle((columnNumber-1)%26+1)
	}
}
