package main

import "fmt"

func main() {
	input := []int{2, 7, 3, 1, 3, 6, 8}
	output := nextGreaterElement(input)
	for _, elem := range output {
		fmt.Print(" ", elem, " ")
	}
}

func nextGreaterElement(input []int) []int {
	lenInput := len(input)
	output := make([]int, lenInput)
	stack := []int{}
	for i := lenInput - 1; i >= 0; i-- {
		currentElem := input[i]

		for len(stack) > 0 && currentElem >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			output[i] = stack[len(stack)-1]
		} else {
			output[i] = -1
		}

		stack = append(stack, currentElem)
	}

	return output
}
