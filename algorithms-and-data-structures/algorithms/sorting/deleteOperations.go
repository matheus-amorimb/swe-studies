package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	nums := []int{50, 30, 70, 20, 80, 80}
	ops := []int{-1, -1}
	res := deleteOperations(nums, ops)
	fmt.Println(res)
}

func deleteOperations(arr, ops []int) []int {
	sortedIdxs := make([]int, len(arr))
	for i := range len(arr) {
		sortedIdxs[i] = i
	}

	slices.SortStableFunc(sortedIdxs, func(a, b int) int {
		return cmp.Compare(arr[a], arr[b])
	})

	deleted := make([]bool, len(arr))
	for _, op := range ops {
		if op >= 0 {
			deleted[op] = true
		} else {
			for pos, i := range sortedIdxs {
				if !deleted[i] {
					deleted[i] = true
					sortedIdxs = sortedIdxs[pos+1:]
					break
				}
			}
		}
	}

	res := make([]int, 0, len(arr))
	for idx, elem := range arr {
		if !deleted[idx] {
			res = append(res, elem)
		}
	}

	return res
}
