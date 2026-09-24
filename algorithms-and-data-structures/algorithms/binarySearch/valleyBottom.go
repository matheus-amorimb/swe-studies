package main

import "fmt"

func main() {
	arr := []int{5, 6, 7}
	asn := valleyBottom(arr)
	fmt.Println("smallest value:", asn)
}

func valleyBottom(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	l, r := 0, len(arr)-1
	if isBefore(r, arr) {
		return arr[r]
	}
	for r-l > 1 {
		m := (r + l) / 2
		if isBefore(m, arr) {
			l = m
		} else {
			r = m
		}
	}

	return arr[l]
}

// descending
func isBefore(i int, arr []int) bool {
	return i == 0 || arr[i] < arr[i-1]
}
