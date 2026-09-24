package main

import "fmt"

func main() {
	sortedArr := []int{-5, -4, -1, 4, 6, 6, 7}
	unsortedArr := []int{-3, 7, 18, 4, 6}
	idx1, idx2 := twoArrayTwoSum(sortedArr, unsortedArr)
	fmt.Println(idx1, idx2)
}

// O(len(unsortedArr * log len(sortedArr)))
func twoArrayTwoSum(sortedArr []int, unsortedArr []int) (int, int) {
	for idxUn, elemUn := range unsortedArr {
		target := -elemUn
		idxSort := binarySearch(target, sortedArr)
		if idxSort != -1 {
			return idxSort, idxUn
		}
	}

	return -1, -1
}

func binarySearch(elem int, arr []int) int {
	l, r := 0, len(arr)-1
	for r-l > 1 {
		mid := (r + l) / 2
		if arr[mid] < elem {
			l = mid
		} else {
			r = mid
		}
	}

	if arr[r] == elem {
		return r
	}

	return -1
}
