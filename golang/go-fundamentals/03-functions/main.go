package main

import (
	"errors"
	"fmt"
)

// Function
func concat(s1 string, s2 string) string {
	return s1 + s2
}

// Multiple Parameters
func addToDatabase(hp, damage int, name string) string {
	// When multiple arguments are of the same type, and are next to each other in the function signature,
	// the type only needs to be declared after the last argument.
	return ""
}

func exercise1(tier string) int {
	var price int
	switch tier {
	case "basic":
		price = 100
	case "premium":
		price = 150
	case "enterprise":
		price = 500
	default:
		price = 0
	}

	return price * 100
}

//Passing Variables by Value

// Ignoring Return Values
// The Go compiler will throw an error if you have any unused variable declarations in your code,
// so you need to ignore anything you don't intend to use.
func getPoint() (x int, y int) {
	return 3, 4
}

//Named Return Values
/*
	func getCoords() (x, y int) {
		// x and y are initialized with zero values

		return // automatically returns x and y
	}
*/
func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return
}

// Explicit Returns
// Using this explicit pattern we can even overwrite the return values
func getCoords() (x, y int) {
	return 5, 6 // this is explicit, x and y are NOT returned
}

// Early Returns
// Go supports the ability to return early from a function. This is a powerful feature that can clean up code, especially when used as guard clauses.
// Guard Clauses leverage the ability to return early from a function (or continue through a loop) to make nested conditionals one-dimensional.
// Instead of using if/else chains, we just return early from the function at the end of each conditional block.
func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return dividend / divisor, nil
}

// Functions As Values
// Go supports first-class and higher-order functions, which are just fancy ways of saying "functions as values".
// Functions are just another type -- like ints and strings and bools.
func reformat(message string, formatter func(string) string) string {
	firstResult := formatter(message)
	secondResult := formatter(firstResult)
	thirdResult := formatter(secondResult)

	return "TEXTIO: " + thirdResult
}

// Anonymous Functions
func conversions(converter func(int) int, x, y, z int) (int, int, int) {
	convertedX := converter(x)
	convertedY := converter(y)
	convertedZ := converter(z)
	return convertedX, convertedY, convertedZ
}

func double(a int) int {
	return a + a
}

func anonymous() (int, int, int) {
	// using a named function
	// newX, newY, newZ := conversions(double, 1, 2, 3)

	// using an anonymous function
	newX, newY, newZ := conversions(func(a int) int {
		return a + a
	}, 1, 2, 3)

	return newX, newY, newZ
}

// Defer
// The defer keyword is a fairly unique feature of Go. It allows a function to be executed automatically just before its enclosing function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
//
// Deferred functions are typically used to clean up resources that are no longer being used.
//
//	Often to close database connections, file handlers and the like.
func GetUsername(dstName, srcName string) (username string, err error) {
	// Open a connection to a database
	// conn, _ := db.Open(srcName)

	// Close the connection *anywhere* the GetUsername function returns
	// defer conn.Close()

	// username, err = db.FetchUser()
	if err != nil {
		// The defer statement is auto-executed if we return here
		return "", err
	}

	// The defer statement is auto-executed if we return here
	return username, nil
}

// Block Scope
// Go is block-scoped. Variables declared inside a block are only accessible within that block (and its nested blocks).
// Blocks are defined by curly braces {}. New blocks are created for:
//
// Functions
// Loops
// If statements
// Switch statements
// Select statements
// Explicit blocks
func lockscope() {
	{
		age := 19
		// this is okay
		fmt.Println(age)
	}

	// this is not okay
	// the age variable is out of scope
	// fmt.Println(age)
}

// Closure
// A closure is a function that references variables from outside its own function body
// The function may access and assign to the referenced variables.
func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

func callconcatter() {
	harryPotterAggregator := concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")
	harryPotterAggregator("of")
	harryPotterAggregator("number")
	harryPotterAggregator("four,")
	harryPotterAggregator("Privet")

	fmt.Println(harryPotterAggregator("Drive"))
	// Mr. and Mrs. Dursley of number four, Privet Drive
}

// Currying
// Function currying is a concept from functional programming and involves partial application of functions.
//
// It allows a function with multiple arguments to be transformed into a sequence of functions, each taking a single argument.
func currying() {
	squareFunc := selfMath(multiply)
	doubleFunc := selfMath(add)

	fmt.Println(squareFunc(5))
	// prints 25

	fmt.Println(doubleFunc(5))
	// prints 10
}

func multiply(x, y int) int {
	return x * y
}

func add(x, y int) int {
	return x + y
}

func selfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}

func main() {}
