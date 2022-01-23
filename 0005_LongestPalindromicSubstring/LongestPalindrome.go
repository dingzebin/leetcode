package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("baabad"))
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res = maxPalindrome(i, i, s, res)
		res = maxPalindrome(i, i+1, s, res)
	}
	return res
}

func maxPalindrome(i, j int, s, res string) string {
	r := ""
	for i >= 0 && j < len(s) && s[i] == s[j] {
		r = s[i : j+1]
		i--
		j++
	}
	if len(r) > len(res) {
		return r
	}
	return res
}

func longestPalindrome2(s string) string {
	if len(s) == 0 {
		return s
	}
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	r := s[:1]
	for i := 1; i < len(s); i++ {
		for j := 0; j < i; j++ {
			dp[j][i] = (s[j] == s[i]) && ((i-j < 3) || dp[j+1][i-1])
			if dp[j][i] && (i-j+1) > len(r) {
				r = s[j : i+1]
			}
		}
	}
	return r
}
