package main

import (
	"fmt"
)

func main() {
	a := 'a'
	d := 'B'
	c := '9'
	b := '!'
	fmt.Printf("toUppercase (%v): %v\n", string(a), string(toUppercase(a)))
	fmt.Printf("toUppercase (%v): %v\n", string(b), string(toUppercase(b)))
	fmt.Printf("toUppercase (%v): %v\n", string(c), string(toUppercase(c)))
	fmt.Printf("toUppercase (%v): %v\n", string(d), string(toUppercase(d)))
}

func toUppercase(c rune) rune {
	isUppCase := isUpperCase(c)
	if isUppCase {
		fmt.Println("doing nothing... character is already uppercase")
		return c
	}

	if isLowerCase(c) {
		c = c - 'a' + 'A'
		return c
	}

	return c
}
