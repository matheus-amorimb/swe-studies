package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	arr = paritySorting(arr)
	fmt.Println(arr)
}

func paritySorting(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	p1, p2 := 0, 0
	// cases:
	// p1 odd, p2 even, switch, p1++, p2++
	// p1 even, p2 even, p2++
	// p1 even, p2 odd, p1=p2, p2++
	// p1 odd, p2 odd, p2++
	for p2 < len(arr) {
		if !isEven(arr[p1]) && isEven(arr[p2]) {
			arr[p1], arr[p2] = arr[p2], arr[p1]
			p1++
			p2++
		} else if isEven(arr[p1]) && !isEven(arr[p2]) {
			p1 = p2
			p2++
		} else {
			p2++
		}
	}

	return arr
}

func isEven(value int) bool {
	return value%2 == 0
}
