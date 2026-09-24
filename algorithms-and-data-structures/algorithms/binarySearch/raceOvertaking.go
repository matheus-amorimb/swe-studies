package main

import "fmt"

func main() {
	p1 := []int{2, 3}
	p2 := []int{1, 4}
	idx := raceOvertaking(p1, p2)
	fmt.Println("idx:", idx)
}

func raceOvertaking(p1, p2 []int) int {
	l, r := 0, len(p1)-1
	for r-l > 1 {
		mid := (r + l) / 2
		if isBeforeOvertaking(mid, p1, p2) {
			l = mid
		} else {
			r = mid
		}

	}

	return r
}

func isBeforeOvertaking(idx int, p1, p2 []int) bool {
	return p1[idx] > p2[idx]
}
