package main

import (
	"fmt"
	"slices"
)

func main() {
	invalidSudoku := [][]int{
		{5, 0, 0, 0, 0, 0, 0, 0, 6},
		{0, 0, 9, 0, 5, 0, 3, 0, 0},
		{0, 3, 0, 0, 0, 2, 0, 0, 0},
		{8, 0, 0, 7, 0, 0, 0, 0, 9},
		{0, 0, 2, 0, 0, 0, 8, 0, 0},
		{4, 0, 0, 0, 0, 6, 0, 0, 3},
		{0, 0, 0, 3, 0, 0, 0, 4, 0},
		{0, 0, 3, 0, 8, 0, 2, 0, 0},
		{9, 0, 0, 0, 0, 0, 0, 0, 7},
	}
	isValid := isValidSudoku(invalidSudoku)
	fmt.Println("isValid: ", isValid)
}

func isValidSudoku(sudoku [][]int) bool {
	rows := map[int][]int{}
	columns := map[int][]int{}
	subgrids := map[int][]int{}

	for idxR, row := range sudoku {
		for idxC, col := range row {
			currValue := col
			subgrid := getSubgrid(idxR, idxC)

			if !isValueValid(currValue, subgrid, subgrids) {
				return false
			} else {
				addValue(currValue, subgrid, subgrids)
			}

			if !isValueValid(currValue, idxR, rows) {
				return false
			} else {
				addValue(currValue, idxR, rows)
			}

			if !isValueValid(currValue, idxC, columns) {
				return false
			} else {
				addValue(currValue, idxC, columns)
			}
		}
	}

	return true
}

func getSubgrid(r, c int) int {
	skippedCol := int(r/3) * 3
	return skippedCol + c/3
}

func isValueValid(value, position int, set map[int][]int) bool {
	if value == 0 {
		return true
	}
	currValues, exist := set[position]
	if !exist {
		return true
	}

	contains := slices.Contains(currValues, value)

	return !contains
}

func addValue(value, position int, set map[int][]int) {
	if value == 0 {
		return
	}
	set[position] = append(set[position], value)
}
