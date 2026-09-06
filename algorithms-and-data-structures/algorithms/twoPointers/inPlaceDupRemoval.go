package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 3, 3, 3, 5}
	arr = inReplaceDupRemoval(arr)
	fmt.Println(arr)
}

func inReplaceDupRemoval(arr []int) []int {
	//loop invariant: every value on the left of p1 is unique
	if len(arr) == 0 {
		return arr
	}

	p1, p2 := 0, 1
	for p2 < len(arr) {
		if arr[p2] != arr[p1] {
			p1++
			arr[p1] = arr[p2]
		}
		p2++
	}

	return arr
}
