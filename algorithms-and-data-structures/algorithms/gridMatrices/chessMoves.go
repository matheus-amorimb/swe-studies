package main

import "fmt"

func main() {
	board := [][]int{
		{0, 0, 0, 1, 0, 0},
		{0, 1, 1, 1, 0, 0},
		{0, 1, 0, 1, 1, 0},
		{1, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
	}
	piece := "king"
	currPos := []int{3, 5}
	uncCells := getPieceUnoccupiedCells(piece, currPos, board)
	fmt.Println(uncCells)
}

func getPieceUnoccupiedCells(piece string, currPos []int, board [][]int) [][]int {
	directions := getPieceDirections(piece)
	r, c := currPos[0], currPos[1]
	moves := make([][]int, 0)
	for _, direc := range directions {
		newR := r + direc[0]
		newC := c + direc[1]
		if piece == "queen" {
			for isValidCell(board, newR, newC) {
				moves = append(moves, []int{newR, newC})
				newR += direc[0]
				newC += direc[1]
			}
		} else if isValidCell(board, newR, newC) {
			moves = append(moves, []int{newR, newC})
		}
	}

	return moves
}

func getPieceDirections(piece string) [][]int {
	switch piece {
	case "king":
		return [][]int{
			{-1, 0},
			{1, 0},
			{0, -1},
			{0, 1},
			{-1, -1},
			{-1, 1},
			{1, -1},
			{1, 1},
		}
	case "queen":
		return [][]int{
			{-1, 0},
			{1, 0},
			{0, -1},
			{0, 1},
			{-1, -1},
			{-1, 1},
			{1, -1},
			{1, 1},
		}
	case "knight":
		return [][]int{
			{-1, -2},
			{-1, 2},
			{1, -2},
			{1, 2},
			{-2, -1},
			{2, -1},
			{-2, 1},
			{2, 1},
		}
	}

	panic("invalid piece")
}

func isValidCell(board [][]int, r, c int) bool {
	if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) {
		return false
	}
	return board[r][c] != 1
}
