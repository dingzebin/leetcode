package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	dict := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	r := make([][]string, len(digits))
	r[0] = append(r[0], dict[digits[0:1]]...)

	for i := 1; i < len(digits); i++ {
		c := dict[digits[i:i+1]]
		for j := 0; j < len(c); j++ {
			for k := 0; k < len(r[i-1]); k++ {
				r[i] = append(r[i], r[i-1][k]+c[j])
			}
		}
	}
	return r[len(digits)-1]
}
