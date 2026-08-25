package main

import "fmt"

func main() {
	arr1 := []int{1, 3, 4, 5, 6}
	arr2 := []int{}
	arrR := merge(arr1, arr2)
	fmt.Printf("[")
	for _, el := range arrR {
		fmt.Printf(" %v ", el)
	}
	fmt.Printf("]\n")
}

func merge(arr1, arr2 []int) []int {
	// parallel pointers
	// loop invariant: all elems from 0...idx have already been sorted in the final array.
	idx1, idx2 := 0, 0
	output := make([]int, 0)
	for idx1 < len(arr1) && idx2 < len(arr2) {
		if arr1[idx1] < arr2[idx2] {
			output = append(output, arr1[idx1])
			idx1++
		} else if arr2[idx2] < arr1[idx1] {
			output = append(output, arr2[idx2])
			idx2++
		} else {
			output = append(output, arr1[idx1])
			output = append(output, arr2[idx2])
			idx1++
			idx2++
		}
	}

	for idx1 < len(arr1) {
		output = append(output, arr1[idx1])
		idx1++
	}

	for idx2 < len(arr2) {
		output = append(output, arr2[idx2])
		idx2++
	}

	return output
}
