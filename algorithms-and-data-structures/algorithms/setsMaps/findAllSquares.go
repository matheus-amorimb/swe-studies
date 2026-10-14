package main

import (
	"fmt"
	"math"
)

func main() {
	//unique, positive integers
	arr := []int{1}
	pars := findAllSquares(arr)
	fmt.Println(pars)
}

func findAllSquares(arr []int) [][]int {
	squares := map[int]int{}
	pars := [][]int{}

	for idx, elem := range arr {
		sqr := int(math.Pow(float64(elem), 2))
		squares[sqr] = idx
	}

	for idx, elem := range arr {
		if idxPair, ok := squares[elem]; ok {
			pars = append(pars, []int{idxPair, idx})
		}
	}

	return pars
}
