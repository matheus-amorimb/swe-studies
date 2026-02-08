package main

import "strconv"

func main() {
	input := [][]byte{
		[]byte{'1', '2', '.', '.', '3', '.', '.', '.', '.'},
		[]byte{'4', '.', '.', '5', '.', '.', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '.', '3'},
		[]byte{'5', '.', '.', '.', '6', '.', '.', '.', '4'},
		[]byte{'.', '.', '.', '8', '.', '3', '.', '.', '5'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '.', '.', '.', '.', '.', '2', '.', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '8'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	println("isValidSudoku: ", isValidSudoku(input))
}

// valid rules
// - each row must contain the digits 1-9 without duplicates
// - each column must contain the digits 1-9 without duplicates
// - each of the nine 3x3 sub-boxes of the grid must contain the digits 1-9 without duplicates

func isValidSudoku(board [][]byte) bool {
	nRows := len(board)

	rows := make([]map[byte]bool, nRows)
	columns := make([]map[byte]bool, nRows)
	subBoxes := make([]map[byte]bool, nRows)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		columns[i] = make(map[byte]bool)
		subBoxes[i] = make(map[byte]bool)
	}

	for idxRow, row := range board {
		for idxCol, value := range row {
			_, err := strconv.Atoi(string(value))
			if err != nil {
				continue
			}

			subBoxIdx := getSubBoxIdx(idxRow, idxCol)

			if rows[idxRow][value] || columns[idxCol][value] || subBoxes[subBoxIdx][value] {
				return false
			}

			rows[idxRow][value] = true
			columns[idxCol][value] = true
			subBoxes[subBoxIdx][value] = true
		}
	}

	return true
}

func getSubBoxIdx(row int, col int) int {
	subGridSize := 3
	boxRow := row / subGridSize
	bowColumn := col / subGridSize

	index := (boxRow * 3) + bowColumn

	return index
}
