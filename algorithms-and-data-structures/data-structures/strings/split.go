package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "ae"
	c := 'e'
	output := split(s, c)
	for _, char := range output {
		fmt.Println(char)
	}
}

func split(s string, c rune) []string {
	curStr := make([]string, 0)
	opt := make([]string, 0)
	for _, char := range s {
		if char == c {
			str := strings.Join(curStr, "")
			opt = append(opt, str)
			curStr = make([]string, 0)
		} else {
			//strings are immutable in go, so string concatenation is not suitable. use dynamic arrays instead.
			curStr = append(curStr, string(char))
		}
	}

	str := strings.Join(curStr, "")
	opt = append(opt, str)

	return opt
}
