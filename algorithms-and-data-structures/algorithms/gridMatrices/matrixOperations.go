package main

import "fmt"

func main() {
	grid := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	matrix := NewMatrix(grid)
	fmt.Println("transpose: ", matrix.Transpose().rows)
	fmt.Println("clockwise rotation: ", matrix.RotateClockwise().rows)
	fmt.Println("anticlockwise rotation: ", matrix.RotateCounterClockwise().rows)
	fmt.Println("horizontal reflection: ", matrix.ReflectHorizontally().rows)
	fmt.Println("vertical reflection: ", matrix.ReflectVertically().rows)
}

type Matrix struct {
	rows [][]float64
}

func NewMatrix(grid [][]float64) *Matrix {
	return &Matrix{
		rows: grid,
	}
}

func (m Matrix) Transpose() Matrix {
	for r := range len(m.rows) {
		for c := range r {
			m.rows[r][c], m.rows[c][r] = m.rows[c][r], m.rows[r][c]
		}
	}

	return m
}

func (m Matrix) RotateClockwise() Matrix {
	m.Transpose().ReflectHorizontally()
	return m
}

func (m Matrix) RotateCounterClockwise() Matrix {
	m.Transpose().ReflectVertically()
	return m
}

func (m Matrix) ReflectHorizontally() Matrix {
	C := len(m.rows[0])
	center := C / 2
	for r := range m.rows {
		for c := range center {
			m.rows[r][c], m.rows[r][C-c-1] = m.rows[r][C-c-1], m.rows[r][c]
		}
	}

	return m
}

func (m Matrix) ReflectVertically() Matrix {
	R, C := len(m.rows), len(m.rows[0])
	center := R / 2
	for r := range center {
		for c := range C {
			m.rows[r][c], m.rows[R-r-1][c] = m.rows[R-r-1][c], m.rows[r][c]
		}
	}

	return m
}
