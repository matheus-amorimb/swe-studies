package main

import "fmt"

func main() {
	grid := [][]int{
		{-1, 2, 3},
		{4, 0, 0},
		{-2, 0, 9},
	}
	subgridSum := subgridSum(grid)
	fmt.Println(subgridSum)
}

func subgridSum(grid [][]int) [][]int {
	R, C := len(grid), len(grid[0])
	res := make([][]int, R)
	for idx, _ := range res {
		res[idx] = make([]int, C)
	}

	for r := R - 1; r >= 0; r-- {
		for c := C - 1; c >= 0; c-- {
			rightSum := getRightSum(r, c, res)
			leftSum := getBelowSum(r, c, res)
			diagonalSum := getDiagonalSum(r, c, res)
			res[r][c] = grid[r][c] + rightSum + leftSum - diagonalSum
		}
	}

	return res
}

func getRightSum(r, c int, grid [][]int) int {
	if c+1 >= len(grid[0]) {
		return 0
	}
	return grid[r][c+1]
}
func getBelowSum(r, c int, grid [][]int) int {
	if r+1 >= len(grid) {
		return 0
	}
	return grid[r+1][c]
}
func getDiagonalSum(r, c int, grid [][]int) int {
	if r+1 >= len(grid) || c+1 >= len(grid[0]) {
		return 0
	}
	return grid[r+1][c+1]
}
