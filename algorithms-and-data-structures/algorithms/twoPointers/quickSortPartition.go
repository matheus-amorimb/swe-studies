package main

import "fmt"

func main() {
	arr := []int{1, 4, 7, 2, 3, 3, 5, 4, 3}
	pivot := 4
	arr = quickSortPartition(arr, pivot)
	fmt.Println(arr)
}

func quickSortPartition(arr []int, pivot int) []int {
	// loop invariant: every element on the left of l is smaller than or equal the pivot
	// every element on the right of r is greather the pivot
	l, r := 0, len(arr)-1
	for l < r {
		if arr[l] <= pivot {
			l++
		} else if arr[r] > pivot {
			r--
		} else {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}

	r2start := r
	if l == r {
		if arr[r] > pivot {
			r2start = r - 1
		}
	}
	l2, r2 := 0, r2start
	for l2 < r2 {
		if arr[l2] < pivot {
			l2++
		} else if arr[r2] == pivot {
			r2--
		} else {
			arr[l2], arr[r2] = arr[r2], arr[l2]
			l2++
			r2--
		}
	}
	return arr
}
