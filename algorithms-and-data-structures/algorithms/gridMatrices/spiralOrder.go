package main

import (
	"fmt"
)

func main() {
	n := 5
	for _, row := range spiralOrder(n) {
		for _, val := range row {
			fmt.Printf("%4d", val)
		}
		fmt.Println()
	}
}

func spiralOrder(len int) [][]int {
	matrix := make([][]int, len)
	for idx := range matrix {
		matrix[idx] = make([]int, len)
	}
	r, c := len-1, len-1
	lastElem := len*len - 1
	offSetR, offSetC := -1, 0
	for lastElem >= 0 {
		for isValid(matrix, r, c) {
			matrix[r][c] = lastElem
			lastElem--
			r, c = r+offSetR, c+offSetC
		}
		r, c = r-offSetR, c-offSetC
		offSetR, offSetC = newMov(offSetR, offSetC)
		r, c = r+offSetR, c+offSetC
	}
	return matrix
}

func newMov(offsetR, offsetC int) (int, int) {
	return -offsetC, offsetR
}

func isValid(matrix [][]int, r, c int) bool {
	lenMatrix := len(matrix)
	if r < 0 || r >= lenMatrix || c < 0 || c >= lenMatrix {
		return false
	}

	if matrix[r][c] != 0 {
		return false
	}

	return true
}
