package main

import (
	"fmt"
	"slices"
)

func main() {
	ans := []rune{'a', 'b', 'c', 'c'}
	m := 5
	students := []Student{
		{
			id:      4,
			desk:    10,
			answers: []rune{'a', 'b', 'c', 'd'},
		},
		{
			id:      1,
			desk:    6,
			answers: []rune{'a', 'b', 'c', 'd'},
		},
		{
			id:      3,
			desk:    8,
			answers: []rune{'a', 'b', 'd', 'd'},
		},
		{
			id:      5,
			desk:    11,
			answers: []rune{'a', 'b', 'c', 'd'},
		},
		{
			id:      9,
			desk:    7,
			answers: []rune{'a', 'b', 'c', 'd'},
		},
		{
			id:      6,
			desk:    16,
			answers: []rune{'a', 'b', 'd', 'd'},
		},
	}
	cheaters := cheaterDetection(ans, m, students)
	fmt.Println("cheaters:", cheaters)
}

type Student struct {
	id      int
	desk    int
	answers []rune
}

func cheaterDetection(answers []rune, desksPerRow int, students []Student) [][]int {
	//Two students are considered suspect if they have made identical mistakes (matching correct scores are not suspicious)
	// and sit next to each other in the same row
	deskToIdx := map[int]int{}
	for idx, student := range students {
		deskToIdx[student.desk] = idx
	}

	cheaters := [][]int{}
	for _, s := range students {
		if !isNextDeskSameRow(desksPerRow, s.desk) {
			continue
		}
		rIdx, ok := deskToIdx[s.desk+1]
		if !ok {
			continue
		}
		rStd := students[rIdx]
		if slices.Equal(s.answers, rStd.answers) && !slices.Equal(s.answers, answers) {
			cheaters = append(cheaters, []int{s.id, rStd.id})
		}
	}

	return cheaters
}

func isNextDeskSameRow(desksPerRow, deskA int) bool {
	deskACol := (deskA - 1) % desksPerRow
	if deskACol == desksPerRow-1 {
		return false
	}
	return true
}
