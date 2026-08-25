package main

import "fmt"

func main() {
	arr1 := []int{0, 1, 3, 4}
	arr2 := []int{3, 3, 3, 4}
	arrInt := arrayIntersection(arr1, arr2)
	fmt.Printf("[")
	for _, elem := range arrInt {
		fmt.Printf(" %v ", elem)
	}
	fmt.Printf("]")
	fmt.Println()
}

func arrayIntersection(arr1, arr2 []int) []int {
	arrayInt := []int{}

	lenArr1 := len(arr1)
	lenArr2 := len(arr2)

	idxArr1 := 0
	idxArr2 := 0
	for idxArr1 < lenArr1 && idxArr2 < lenArr2 {
		elem1 := arr1[idxArr1]
		elem2 := arr2[idxArr2]

		if elem1 < elem2 {
			idxArr1++
		} else if elem1 > elem2 {
			idxArr2++
		} else {
			arrayInt = append(arrayInt, elem1)
			idxArr1++
			idxArr2++
		}
	}

	return arrayInt
}
