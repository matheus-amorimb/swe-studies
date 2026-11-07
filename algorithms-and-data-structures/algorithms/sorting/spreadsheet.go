package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	spreadSheet := Spreadsheet{}
	spreadSheet.new(3, 3)
	spreadSheet.set(0, 0, 5)
	spreadSheet.set(0, 1, 3)
	spreadSheet.set(0, 2, 8)
	spreadSheet.set(1, 0, 6)
	spreadSheet.set(2, 1, 1)
	spreadSheet.print("original matrix")
	spreadSheet.sortColumnsByRow(0)
	spreadSheet.sortRowsByColumn(1)
}

type Spreadsheet struct {
	matrix [][]int
}

func (s *Spreadsheet) new(r, c int) {
	matrix := make([][]int, r)
	for idx := range matrix {
		matrix[idx] = make([]int, c)
	}
	s.matrix = matrix
}

func (s *Spreadsheet) set(r, c, value int) {
	s.matrix[r][c] = value
}

func (s *Spreadsheet) get(r, c int) int {
	return s.matrix[r][c]
}

func (s *Spreadsheet) sortColumnsByRow(r int) {
	rLen := len(s.matrix[0])
	newIdxs := make([]int, rLen)
	for idx := range newIdxs {
		newIdxs[idx] = idx
	}
	slices.SortStableFunc(newIdxs, func(a, b int) int {
		return cmp.Compare(s.matrix[r][a], s.matrix[r][b])
	})
	for idxR, r := range s.matrix {
		rBuf := slices.Clone(r)
		for idxC := range r {
			s.matrix[idxR][idxC] = rBuf[newIdxs[idxC]]
		}
	}
	s.print("sortedByRow")
}

func (s *Spreadsheet) sortRowsByColumn(c int) {
	slices.SortStableFunc(s.matrix, func(a, b []int) int {
		return cmp.Compare(a[c], b[c])
	})
	s.print("sortedByColumn")
}

func (s *Spreadsheet) print(t string) {
	fmt.Println(fmt.Sprintf("%v:", t), s.matrix)
}
