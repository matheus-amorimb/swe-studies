package main

import "strings"

func main() {
	input := 7
	output := generateParenthesis(input)
	for _, item := range output {
		println(item)
	}
}

func generateParenthesis(n int) []string {
	var output, stack []string

	var backtrackFunc func(int, int)
	backtrackFunc = func(open, close int) {
		if open < n {
			stack = append(stack, "(")
			backtrackFunc(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ")")
			backtrackFunc(open, close+1)
			stack = stack[:len(stack)-1]
		}

		if open == n && close == n {
			paranthesis := strings.Join(stack, "")
			output = append(output, paranthesis)
			return
		}
	}

	backtrackFunc(0, 0)
	return output
}
