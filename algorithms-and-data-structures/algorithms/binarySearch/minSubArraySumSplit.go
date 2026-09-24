package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := []int{10, 5, 8, 9, 11}
	k := 3
	minSum := minSubArraySumSplit(arr, k)
	fmt.Println("minSum:", minSum)
}

func minSubArraySumSplit(arr []int, k int) int {
	l, r := slices.Max(arr), sum(arr)
	for r-l > 1 {
		mid := (l + r) / 2
		fmt.Print("l: ", l, " r: ", r, " | ")
		if isSumBefore(arr, k, mid) {
			l = mid
		} else {
			r = mid
		}

	}
	return r
}

func isSumBefore(arr []int, k, maxSum int) bool {
	splitsRequired := getSplitsRequired(arr, maxSum)
	return splitsRequired > k
}

func getSplitsRequired(arr []int, maxSum int) int {
	splitsRequired := 1
	currSum := 0
	fmt.Print("maxSum: ", maxSum, " |")
	for _, elem := range arr {
		if currSum+elem > maxSum {
			fmt.Print("|")
			splitsRequired++
			currSum = elem
			fmt.Print(" ", elem, " ")
		} else {
			fmt.Print(" ", elem, " ")
			currSum += elem
		}
	}

	fmt.Print("| splitsRequired: ", splitsRequired)
	fmt.Println()
	return splitsRequired
}

func sum(arr []int) int {
	sum := 0
	for _, elem := range arr {
		sum += elem
	}
	return sum
}
