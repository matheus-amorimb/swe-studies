package main

import (
	"fmt"
	"math"
)

func main() {
	arr1 := []int{1}
	arr2 := []int{2, 10}
	arr3 := []int{3, 9}
	arr := threeWayMerge(arr1, arr2, arr3)
	fmt.Printf("[")
	for _, elem := range arr {
		fmt.Printf(" %v ", elem)
	}
	fmt.Printf("]\n")
}

func threeWayMerge(arr1, arr2, arr3 []int) []int {
	//3 parallel pointers
	//loop invariant: the current arrX[idxX] in arrays XYX is always equal or greater than the last element of output
	//cases:
	// - arrX[idX] is the smallest number between XYZ, idxX++
	// - arrX[idX] == arrY[idxY], and they are the smallest numbers, idxX++ and idxY++
	// - output.append only if arrX[idX] != lastElemOutput
	//exit case: idxX < len(idxX) for x e {1, 2, 3}

	idx1, idx2, idx3 := 0, 0, 0
	output := make([]int, 0)
	for idx1 < len(arr1) || idx2 < len(arr2) || idx3 < len(arr3) {
		min := math.MaxInt

		if idx1 < len(arr1) && arr1[idx1] < min {
			min = arr1[idx1]
		}

		if idx2 < len(arr2) && arr2[idx2] < min {
			min = arr2[idx2]
		}

		if idx3 < len(arr3) && arr3[idx3] < min {
			min = arr3[idx3]
		}

		if idx1 < len(arr1) && min == arr1[idx1] {
			idx1++
		}

		if idx2 < len(arr2) && min == arr2[idx2] {
			idx2++
		}

		if idx3 < len(arr3) && min == arr3[idx3] {
			idx3++
		}

		output = appendToSortedArrayWithoutDuplicates(min, output)

	}
	return output
}

func appendToSortedArrayWithoutDuplicates(value int, arr []int) []int {
	if len(arr) == 0 {
		arr = append(arr, value)
		return arr
	}

	lastElem := arr[len(arr)-1]
	if value > lastElem {
		arr = append(arr, value)
		return arr
	}

	return arr
}
