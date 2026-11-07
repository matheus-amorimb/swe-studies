package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Circle struct {
	x int
	y int
	r int
}

func main() {
	circles := []Circle{
		{x: 5, y: 3, r: 3},
		{x: 5, y: 3, r: 2},
		{x: 4, y: 4, r: 5},
	}
	// circles := []Circle{
	// 	{x: 5, y: 3, r: 3},
	// }

	fmt.Println(nestedCircles(circles))
}

func nestedCircles(circles []Circle) bool {
	slices.SortFunc(circles, func(a, b Circle) int {
		return cmp.Compare(b.r, a.r)
	})
	return nestedSorted(circles)
}

func nestedSorted(circles []Circle) bool {
	if len(circles) <= 1 {
		return true
	}
	return isNested(circles[1], circles[0]) && nestedSorted(circles[1:])
}

func isNested(smaller, larger Circle) bool {
	dx := smaller.x - larger.x
	dy := smaller.y - larger.y
	gap := larger.r - smaller.r

	return gap > 0 && dx*dx+dy*dy < gap*gap
}
