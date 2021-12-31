package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
	fmt.Println(lengthOfLastWord("l"))
}

func lengthOfLastWord(s string) int {
	lenOfLastWord := 0
	for i := len(s); i > 0; i-- {
		if s[(i-1):i] == " " {
			if lenOfLastWord > 0 {
				break
			}
		} else {
			lenOfLastWord += 1
		}
	}
	return lenOfLastWord
}
