package main

import (
	"fmt"
)

func main() {
	fmt.Println(strStrKMP("hello", "ll"))
	fmt.Println(strStrKMP("", ""))
	fmt.Println(strStrKMP("aaaaa", "bba"))
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

func strStrKMP(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := make([]int, len(needle))
	getNext(needle, next)
	fmt.Println(next)
	i, j := 0, 0
	for i < len(haystack) && j < len(needle) {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(needle) {
		return i - j
	}
	return -1
}

func getNext(needle string, next []int) {
	next[0] = -1
	for i, j := 0, -1; i < len(needle)-1; {
		if j == -1 || needle[j] == needle[i] {
			j++
			i++
			next[i] = j
		} else {
			j = next[j]
		}
	}
}
