package main

import "fmt"

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		return addBinary(b, a)
	}
	r := ""
	i, j := len(a)-1, len(b)-1
	tmp := 0
	for i >= 0 {
		if j >= 0 && b[j] == 49 {
			tmp += 1
		}
		if a[i] == 49 {
			tmp += 1
		}
		if tmp%2 == 0 {
			r = "0" + r
		} else {
			r = "1" + r
		}
		tmp = tmp / 2
		i--
		j--
	}
	if tmp > 0 {
		r = "1" + r
	}
	return r
}
