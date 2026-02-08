package main

import "strconv"

func main() {
	input := []string{"4", "13", "5", "/", "+"}
	output := evalRPN(input)
	println(output)
}

func evalRPN(tokens []string) int {
	tokenOperatorToOperation := map[string]func(int, int) int{
		"+": Add,
		"-": Subtract,
		"*": Multiply,
		"/": Divide,
	}

	stack := []int{}
	for _, token := range tokens {
		number, err := strconv.Atoi(string(token))
		if err != nil {
			operation := tokenOperatorToOperation[token]
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, operation(a, b))
		} else {
			stack = append(stack, number)
		}
	}

	return stack[0]
}

func Add(val1 int, val2 int) int {
	return val1 + val2
}

func Subtract(val1 int, val2 int) int {
	return val1 - val2
}

func Divide(val1 int, val2 int) int {
	return val1 / val2
}

func Multiply(val1 int, val2 int) int {
	return val1 * val2
}
