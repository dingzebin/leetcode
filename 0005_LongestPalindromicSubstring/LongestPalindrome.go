package main

import "fmt"

/*
Given a string s, return the longest palindromic substring in s

babad

ba
a
*/

func main() {
	fmt.Println(longestPalindrome("baabad"))
}

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	longestPalindromeStr := s[:1]
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}

	for i := 1; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if isPalindrome(j, i, s, dp) {
				palindromStr := s[j : i+1]
				if len(palindromStr) > len(longestPalindromeStr) {
					longestPalindromeStr = palindromStr
				}
			}
		}
	}

	return longestPalindromeStr
}
func isPalindrome(start int, end int, s string, dp [][]int) bool {
	if start == end {
		return true
	}

	if dp[start][end] != 0 {
		fmt.Println("1")
		return dp[start][end] == 1
	}

	if s[start] != s[end] {
		dp[start][end] = -1
		return false
	}

	if (end - 1) == start {
		if s[start] == s[end] {
			dp[start][end] = 1
			return true
		} else {
			dp[start][end] = -1
			return false
		}
	}
	return isPalindrome(start+1, end-1, s, dp)
}
