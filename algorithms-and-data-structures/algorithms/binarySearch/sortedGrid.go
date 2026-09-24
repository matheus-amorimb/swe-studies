package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 2, 4, 5},
		{6, 7, 8, 9},
	}
	target := 1
	r, c := sortedGrid(target, grid)
	fmt.Println("r:", r, "|", "c:", c)
}

func sortedGrid(target int, grid [][]int) (int, int) {
	R, C := len(grid), len(grid[0])
	if R == 0 {
		return -1, -1
	}

	l, r := 0, (R*C)-1
	if isBefore(target, r, grid) {
		return -1, -1
	}
	if !isBefore(target, l, grid) {
		r = l
	}
	for r-l > 1 {
		mid := (r + l) / 2
		if isBefore(target, mid, grid) {
			l = mid
		} else {
			r = mid
		}
	}

	idxR, idxC := getRowAndColumnIdxFromFlattenedGrid(r, grid)

	if grid[idxR][idxC] == target {
		return idxR, idxC
	}

	return -1, -1
}

func getRowAndColumnIdxFromFlattenedGrid(idx int, grid [][]int) (int, int) {
	C := len(grid[0])
	idxR, idxC := int(idx/C), idx%C
	return idxR, idxC
}

func isBefore(target, mid int, grid [][]int) bool {
	idxR, idxC := getRowAndColumnIdxFromFlattenedGrid(mid, grid)
	return grid[idxR][idxC] < target
}

// func binarySearch(target int, arr []int) int {
// 	l, r := 0, len(arr)-1
// 	if arr[l] >= target || arr[r] < target {
// 		if arr[l] == target {
// 			return 0
// 		}
// 		return -1
// 	}
// 	for r-l > 1 {
// 		mid := (r + l) / 2
// 		if arr[mid] < target {
// 			l = mid
// 		} else {
// 			r = mid
// 		}

// 	}

// 	if arr[r] == target {
// 		return r
// 	} else {
// 		return -1
// 	}
// }
