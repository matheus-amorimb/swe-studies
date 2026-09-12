package main

import "fmt"

func main() {
	board := [][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0},
	}
	board = queenReach(board)
	fmt.Println(board)
}

func queenReach(board [][]int) [][]int {
	boardSafety := make([][]int, len(board))
	for row := range boardSafety {
		boardSafety[row] = make([]int, len(board[0]))
	}

	for idxRow, row := range board {
		for idxCol, col := range row {
			if col == 1 {
				boardSafety[idxRow][idxCol] = 1
				markReachableCells(board, boardSafety, idxRow, idxCol)
			}
		}
	}

	return boardSafety
}

var queenDirections = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
	{-1, -1},
	{-1, 1},
	{1, -1},
	{1, 1},
}

func markReachableCells(board, boardSafety [][]int, r, c int) {
	for _, dir := range queenDirections {
		newR, newC := r+dir[0], c+dir[1]
		for isValidCell(board, newR, newC) {
			boardSafety[newR][newC] = 1
			newR += dir[0]
			newC += dir[1]
		}
	}
}

func isValidCell(board [][]int, r, c int) bool {
	if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) {
		return false
	}
	return board[r][c] != 1
}
