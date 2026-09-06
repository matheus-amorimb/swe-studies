package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr = reverseArray(arr)
	fmt.Println(arr)
}

func reverseArray(arr []int) []int {
	lP, rP := 0, len(arr)-1
	for lP < rP {
		arr[lP], arr[rP] = arr[rP], arr[lP]
		lP++
		rP--
	}

	return arr
}
