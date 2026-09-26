package main

import (
	"fmt"
	"math"
)

func main() {
	pictures := [][][]int{
		{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		{{1, 0, 0}, {0, 0, 0}, {1, 0, 0}},
		{{1, 1, 0}, {0, 0, 0}, {1, 0, 0}},
		{{1, 1, 0}, {1, 1, 1}, {1, 0, 0}},
		{{1, 1, 1}, {1, 1, 1}, {1, 1, 0}},
	}
	ans := tideAerialView(pictures)
	fmt.Println("idx:", ans)
}

func tideAerialView(pictures [][][]int) int {
	l, r := 0, len(pictures)-1
	for r-l > 1 {
		mid := (r + l) / 2
		if getAerialViewBalance(pictures[mid]) < 0.5 {
			l = mid
		} else {
			r = mid
		}
	}

	fmt.Println("r:", r, "|", "balance:", getAerialViewBalance(pictures[r]), "|", "dist:", dist(getAerialViewBalance(pictures[r])))
	fmt.Println("l:", l, "|", "balance:", getAerialViewBalance(pictures[l]), "|", "dist:", dist(getAerialViewBalance(pictures[l])))

	if dist(getAerialViewBalance(pictures[l])) <= dist(getAerialViewBalance(pictures[r])) {
		return l
	}

	return r
}

func getAerialViewBalance(picture [][]int) float32 {
	ones := 0
	for _, r := range picture {
		onesInRow := getOnesInRow(r)
		ones += onesInRow
	}
	total := math.Pow(float64(len(picture)), 2)
	return float32(ones) / float32(total)
}

func getOnesInRow(row []int) int {
	if row[0] == 0 {
		return 0
	}
	if row[len(row)-1] == 1 {
		return len(row)
	}

	l, r := 0, len(row)-1
	for r-l > 1 {
		mid := (l + r) / 2
		if row[mid] == 1 {
			l = mid
		} else {
			r = mid
		}
	}

	return r
}

func dist(f float32) float32 {
	if f < 0.5 {
		return 0.5 - f
	}
	return f - 0.5
}
