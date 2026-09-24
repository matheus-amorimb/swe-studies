package main

import (
	"fmt"
	"math/rand"
)

func main() {
	target := 0
	l, r := 0, 1

	for fetch(r) != -1 {
		r = 2 * r
	}

	lowerBound := fetch(l)
	upperBound := fetch(r)
	if lowerBound >= target || (upperBound < target && upperBound != -1) {
		if lowerBound == target {
			fmt.Println("idx:", 0)
			return
		}
		fmt.Println("idx:", -1)
		return
	}
	for r-l > 1 {
		mid := (l + r) / 2
		value := fetch(mid)
		if value < target && value != -1 {
			l = mid
		} else {
			r = mid
		}
	}

	fmt.Println("idx:", r)
	return
}

var huge = generateSortedArray(10_000_000, 3)

func fetch(i int) int {
	arr := []int{0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 3, 4, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 8, 8, 8, 8, 8}
	if i >= len(arr) {
		return -1
	}
	return arr[i]
}

func generateSortedArray(n int, maxDelta int) []int {
	arr := make([]int, n)
	current := 1
	for i := 0; i < n; i++ {
		arr[i] = current
		current += rand.Intn(maxDelta + 1)
	}
	return arr
}
