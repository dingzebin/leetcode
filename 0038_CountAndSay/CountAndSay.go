package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(5))
}

func countAndSay(n int) string {
	s := "1"
	for i := 1; i < n; i++ {
		cur, count := s[:1], 1
		tmp := ""
		for j := 1; j < len(s); j++ {
			if s[j:j+1] != cur {
				tmp += strconv.Itoa(count) + cur
				cur, count = s[j:j+1], 1
			} else {
				count++
			}
		}
		tmp += strconv.Itoa(count) + cur
		s = tmp
	}
	return s
}
