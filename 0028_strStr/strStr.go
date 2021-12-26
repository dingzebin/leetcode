package main

import (
	"fmt"
)

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("", ""))
	fmt.Println(strStr("aaaaa", "bba"))
}

func strStr(haystack string, needle string) int {
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				return i
			}
			if i+j == len(haystack) {
				return -1
			}
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}
