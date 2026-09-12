package main

import (
	"fmt"
)

func main() {
	snowPrints := [][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0},
		{1, 1, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 1},
	}
	closest := closestToTheRiver(snowPrints)
	fmt.Println(closest)
}

func closestToTheRiver(snowPrints [][]int) int {
	//interate through all the matrix would take O(R*C)
	closest := len(snowPrints) - 1
	for idx, row := range snowPrints {
		if row[0] == 1 {
			closest = idx
		}
	}

	directionsRows := []int{-1, 0, 1}
	r, c := closest, 0
	for c < len(snowPrints[0])-1 {
		for _, dirR := range directionsRows {
			newR, newC := r+dirR, c+1
			if isPrint(snowPrints, newR, newC) {
				r, c = newR, newC
				closest = min(closest, r)
				break
			}
		}
	}

	return closest
}

func isPrint(matrix [][]int, r, c int) bool {
	lenRow := len(matrix)
	lenCol := len(matrix[0])
	if r < 0 || r >= lenRow || c < 0 || c >= lenCol {
		return false
	}

	return matrix[r][c] == 1
}
