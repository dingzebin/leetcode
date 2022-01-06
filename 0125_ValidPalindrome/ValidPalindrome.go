package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))
	fmt.Println(isPalindrome("0P"))
}

func isPalindrome(s string) bool {
	start, end := 0, len(s)-1
	for start < end {
		if !isLetterOrDigit(s[start]) {
			start++
		} else if !isLetterOrDigit(s[end]) {
			end--
		} else {
			if s[start] == s[end] || (s[start] > 64 && s[end] >= 64 && diff(s[start], s[end]) == 32) {
				start++
				end--
			} else {
				return false
			}
		}
	}
	return true
}
func diff(x, y byte) byte {
	if x > y {
		return x - y
	}
	return y - x
}
func isLetterOrDigit(b byte) bool {
	return (b >= 48 && b <= 57) || (b >= 65 && b <= 90) || (b >= 97 && b <= 122)
}
