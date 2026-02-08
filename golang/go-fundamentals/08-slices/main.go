package main

import (
	"errors"
	"fmt"
)

func main() {
	array := [4]int{0, 1, 2, 3}
	originalSlice := array[:]
	copySlice := originalSlice
	underlyingArray(originalSlice)

	fmt.Println("array: ", array)
	fmt.Println("copySlice: ", copySlice)
}

// Arrays
func arrays() {
	var myInts [10]int

	primes := [6]int{2, 3, 5, 7, 11, 13}
}

// Slices
func slices() {
	myFirstSlice := []int{}
	primes := [6]int{2, 3, 5, 7, 11, 13}
	mySlice := primes[1:4]
}

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	// ?
	if plan == planPro {
		return messages[:], nil
	}

	if plan == planFree {
		return messages[:2], nil
	}

	return nil, errors.New("unsupported plan")
}

func underlyingArray(slice []int) {
	slice = slice[1:]
	fmt.Println("underlyingSlice: ", slice)
}

// Make
func getMessageCosts(messages []string) []float64 {
	var messagesLen = len(messages)
	costs := make([]float64, messagesLen)
	for i := 0; i < messagesLen; i++ {
		costs[i] = float64(len(messages[i])) * 0.01
	}

	return costs
}

// Variadic
func sum(nums ...int) int {
	acc := 0
	for i := 0; i < len(nums); i++ {
		acc += nums[i]
	}

	return acc
}

// Append
type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	dayCosts := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		if cost.day == day {
			dayCosts = append(dayCosts, cost.value)
		}
	}

	return dayCosts
}

// Range
func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, message := range msg {
		for _, badWord := range badWords {
			if message == badWord {
				return i
			}
		}
	}

	return -1
}

// Slice of Slices
func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = i * j
		}
	}
	return matrix
}
