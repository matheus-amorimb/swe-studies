package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fizzbuzz()
	printPrimes(20)
}

// Loops
func bulkSend(numMessages int) float64 {
	totalPrice := 0.0
	for i := 0; i < numMessages; i++ {
		fee := float64(i) * 0.01
		totalPrice += 1.0 + fee
	}

	return totalPrice
}

// Ommitting conditions
func maxMessages(thresh int) int {
	var totalCost int
	for i := 0; ; i++ {
		messageCost := 100 + i
		totalCost += messageCost
		if totalCost > thresh {
			return i
		}
	}
}

// No while loop
func whileLoop() {
	plantHeight := 1
	for plantHeight < 5 {
		fmt.Println("still growing! current height:", plantHeight)
		plantHeight++
	}
}

// Exercise
func fizzbuzz() {
	for i := 1; i < 101; i++ {
		var textToPrint string
		isThreeMultiple := i%3 == 0
		isFiveMultiple := i%5 == 0
		if isThreeMultiple {
			textToPrint += "fizz"
		}
		if isFiveMultiple {
			textToPrint += "buzz"
		}
		if textToPrint == "" {
			textToPrint = strconv.Itoa(i)
		}

		fmt.Println(textToPrint)
	}
}

// Exercise 2
func printPrimes(max int) {
	for n := 2; n < max; n++ {
		isTwo := n == 2
		if isTwo {
			fmt.Println(n)
			continue
		}

		isEven := n%2 == 0
		if isEven {
			continue
		}

		isPrime := true
		for i := 5; i*i < n; i += 6 {
			if n%i == 0 || n%(i+2) == 0 {
				isPrime = false
			}
		}

		if isPrime {
			fmt.Println(n)
		}
	}
}

// Exercise 3
func countConnections(groupSize int) int {
	// ?
	return 0
}
