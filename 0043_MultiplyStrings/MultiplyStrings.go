package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println(multiply("2", "3"))
	fmt.Println(multiply("123", "456"))
	// fmt.Println(multiply("9133", "0"))
	//fmt.Println(multiply("498828660196",	"840477629533"))
}

func multiply(num1 string, num2 string) string {
	pos := make([]byte, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			mul := (num1[i] - 48) * (num2[j] - 48)
			p1, p2 := i+j, i+j+1
			sum := mul + pos[p2]
			pos[p1] += sum / 10
			pos[p2] = sum % 10
		}
	}
	res := ""
	for i := 0; i < len(pos); i++ {
		if !(len(res) == 0 && pos[i] == 0) {
			res += strconv.Itoa(int(pos[i]))
		}
	}
	if len(res) == 0 {
		return "0"
	} else {
		return res
	}
}

func multiply2(num1 string, num2 string) string {
	r := ""
	for i, j := len(num1)-1, 0; i >= 0; i-- {
		r = add(mul(num2, num1[i:], j), r)
		j++
	}
	return r
}

func add(n1, n2 string) string {
	if len(n1) < len(n2) {
		return add(n2, n1)
	}
	r := ""
	tmp := byte(0)
	for l1, l2 := len(n1)-1, len(n2)-1; l1 >= 0; {
		n := n1[l1] - 48 + tmp
		if l2 >= 0 {
			n += n2[l2] - 48
		}
		r = strconv.Itoa(int(n%10)) + r
		tmp = n / 10
		l1--
		l2--
	}
	if tmp > 0 {
		r = strconv.Itoa(int(tmp)) + r
	}
	return r
}

func mul(n1, n2 string, numOfZero int) string {
	if n1 == "0" || n2 == "0" {
		return "0"
	}
	r := ""
	num2 := n2[0] - 48
	for i := 0; i < numOfZero; i++ {
		r += "0"
	}
	tmp := byte(0)

	for i := len(n1) - 1; i >= 0; i-- {
		n := (n1[i]-48)*num2 + tmp
		r = strconv.Itoa(int(n%10)) + r
		tmp = n / 10
	}
	if tmp > 0 {
		r = strconv.Itoa(int(tmp)) + r
	}
	return r
}
