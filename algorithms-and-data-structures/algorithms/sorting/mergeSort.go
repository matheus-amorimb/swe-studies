package main

import "fmt"

func main() {
	arr := []int{7, -2, -1, 12, 1, 10, -2, 3, 9, 11}
	sorted := mergesort(arr, 0)
	fmt.Println()
	fmt.Println("sorted:", sorted)
}

func mergesort(arr []int, counter int) []int {
	counter++
	n := len(arr)
	if n <= 1 {
		return arr
	}
	left := mergesort(arr[:n/2], counter)
	right := mergesort(arr[n/2:], counter)
	return mergeTwoSortedArr(left, right)
}

func mergeTwoSortedArr(sortedArr1, sortedArr2 []int) []int {
	p1, p2 := 0, 0
	mergedArr := make([]int, 0, len(sortedArr1)+len(sortedArr2))
	for p1 < len(sortedArr1) && p2 < len(sortedArr2) {
		if sortedArr1[p1] < sortedArr2[p2] {
			mergedArr = append(mergedArr, sortedArr1[p1])
			p1++
		} else if sortedArr2[p2] < sortedArr1[p1] {
			mergedArr = append(mergedArr, sortedArr2[p2])
			p2++
		} else {
			mergedArr = append(mergedArr, sortedArr1[p1])
			mergedArr = append(mergedArr, sortedArr2[p2])
			p1++
			p2++
		}
	}

	mergedArr = append(mergedArr, sortedArr1[p1:]...)
	mergedArr = append(mergedArr, sortedArr2[p2:]...)

	return mergedArr
}
