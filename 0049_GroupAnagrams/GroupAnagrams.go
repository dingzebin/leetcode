package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func groupAnagrams(strs []string) [][]string {
	dict := map[string][]string{}
	for _, str := range strs {
		b := []rune(str)
		sort.Sort(sortRunes(b))
		key := string(b)
		dict[key] = append(dict[key], str)
	}
	res := [][]string{}
	for _, val := range dict {
		res = append(res, val)
	}
	return res
}
