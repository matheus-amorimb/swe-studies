package main

import (
	"fmt"
)

func main() {
	grid := [][]int{
		{1, 5, 3},
		{4, -1, 0},
		{2, 0, 2},
	}
	subgrid := subgridMax(grid)
	fmt.Println(subgrid)
}

func subgridMax(grid [][]int) [][]int {
	R, C := len(grid), len(grid[0])
	for idxR := R - 1; idxR >= 0; idxR-- {
		for idxC := C - 1; idxC >= 0; idxC-- {
			currValue := grid[idxR][idxC]
			rightValue, belowValue := getAdjacentValues(grid, idxR, idxC)
			grid[idxR][idxC] = greatestValue(currValue, rightValue, belowValue)
		}
	}

	return grid
}

func getAdjacentValues(grid [][]int, r, c int) (int, int) {
	var rightValue = 0
	var belowValue = 0
	if c+1 < len(grid[0]) {
		rightValue = grid[r][c+1]
	}

	if r+1 < len(grid) {
		belowValue = grid[r+1][c]
	}

	return rightValue, belowValue
}

func greatestValue(a, b, c int) int {
	return max(a, max(b, c))
}
