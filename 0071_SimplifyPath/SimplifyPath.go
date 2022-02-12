package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
}

func simplifyPath(path string) string {
	pathArr := []string{}
	for _, s := range strings.Split(path, "/") {
		if s == ".." {
			if len(pathArr) > 0 {
				pathArr = pathArr[:len(pathArr)-1]
			}
		} else if s != "" && s != "." {
			pathArr = append(pathArr, s)
		}
	}
	return "/" + strings.Join(pathArr, "/")
}
