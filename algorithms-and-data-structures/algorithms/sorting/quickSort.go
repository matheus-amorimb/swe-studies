package main

import (
	"cmp"
	"fmt"
	"math/rand"
	"slices"
)

func main() {
	arr := []int{1, 25, 6, 8, 17, 7, 9, 2, 3}
	desc := func(a, b int) int {
		//descending
		//b > a => must return positive
		//b < a => must return negative
		return cmp.Compare(b, a)
	}
	sorted := quicksort(arr, desc)
	fmt.Println(sorted)
}

func quicksort(arr []int, cmp func(a, b int) int) []int {
	if len(arr) <= 1 {
		return arr
	}
	before, equal, after := partition(arr, cmp)
	sSorted := quicksort(before, cmp)
	eSorted := equal
	lSorted := quicksort(after, cmp)

	return slices.Concat(sSorted, eSorted, lSorted)
}

func partition(arr []int, cmp func(a, b int) int) ([]int, []int, []int) {
	rand := rand.Intn(len(arr))
	pivot := arr[rand]
	before, equal, after := []int{}, []int{}, []int{}
	for _, elem := range arr {
		c := cmp(elem, pivot)
		switch {
		case c < 0:
			before = append(before, elem)
		case c > 0:
			after = append(after, elem)
		default:
			equal = append(equal, elem)
		}
	}

	return before, equal, after
}
