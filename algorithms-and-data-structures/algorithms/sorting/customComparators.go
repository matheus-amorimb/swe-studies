package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println("caseInsensitiveSortDesc:", caseInsensitiveSort())
	fmt.Println("sortByElementAtIdx:", sortByElementAtIdx())
	fmt.Println("sortByField:", sortByField())
	fmt.Println("newDeckOrder:", newDeckOrder())
	fmt.Println("stableSorting:", stableSorting())
}

func caseInsensitiveSort() []string {
	arr := []string{"apple", "banana", "3", "Cherry", "42", "GRAPE", "10"}
	descOrd := func(a, b string) int {
		return cmp.Compare(strings.ToLower(b), strings.ToLower(a))
	}

	slices.SortFunc(arr, descOrd)
	return arr
}

func sortByElementAtIdx() [][]int {
	arr := [][]int{
		{3, 9}, {1, 4}, {4, 7}, {2, 3},
	}
	by := func(a, b []int) int {
		return cmp.Compare(a[1], b[1])
	}
	slices.SortFunc(arr, by)
	return arr
}

type card struct {
	value int
	suit  string
}

func sortByField() []card {
	arr := []card{
		{8, "hearts"},
		{8, "clubs"},
		{3, "clubs"},
		{3, "hearts"},
	}

	by := func(a, b card) int {
		suitToInt := map[string]int{
			"clubs":    0,
			"hearts":   1,
			"spades":   2,
			"diamonds": 3,
		}
		return cmp.Or(
			cmp.Compare(a.value, b.value),
			cmp.Compare(suitToInt[a.suit], suitToInt[b.suit]),
		)
	}

	slices.SortFunc(arr, by)
	return arr
}

func newDeckOrder() []card {
	arr := []card{
		{8, "hearts"},
		{8, "clubs"},
		{3, "clubs"},
		{3, "hearts"},
	}

	by := func(a, b card) int {
		suitToInt := map[string]int{
			"hearts":   0,
			"clubs":    1,
			"diamonds": 2,
			"spades":   3,
		}
		return cmp.Or(
			cmp.Compare(suitToInt[a.suit], suitToInt[b.suit]),
			cmp.Compare(a.value, b.value),
		)
	}

	slices.SortFunc(arr, by)
	return arr
}

func stableSorting() []card {
	arr := []card{
		{9, "clubs"},
		{4, "spades"},
		{9, "spades"},
		{4, "clubs"},
	}

	cardsToIdx := map[card]int{}
	for idx, card := range arr {
		cardsToIdx[card] = idx
	}

	// by := func(a, b card) int {
	// 	return cmp.Or(
	// 		cmp.Compare(a.value, b.value),
	// 		cmp.Compare(cardsToIdx[a], cardsToIdx[b]),
	// 	)
	// }

	byStable := func(a, b card) int {
		return cmp.Or(
			cmp.Compare(a.value, b.value),
		)
	}

	// slices.SortFunc(arr, by)
	slices.SortStableFunc(arr, byStable)
	return arr
}
