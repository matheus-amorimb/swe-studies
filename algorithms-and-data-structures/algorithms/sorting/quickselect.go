package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func main() {
	arr := []int{15, 4, 4, 13, 8, 10, 5, 2, 20, 3, 9, 11, 27}
	k := 5
	fmt.Println(quickselect(arr, k))
}

func ksmallest(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	small := slices.Min(arr)
	arrMap := map[int]int{}
	for _, elem := range arr {
		arrMap[elem]++
	}

	res := make([]int, 0, k)
	nextSmall := small
	for len(res) < k {
		for i := 0; i < arrMap[nextSmall] && len(res) < k; i++ {
			res = append(res, nextSmall)
		}
		nextSmall++
	}

	return res
}

func quickselect(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	small, equal, larg := partitionQS(arr)

	switch {
	case k <= len(small):
		return quickselect(small, k)
	case k <= len(small)+len(equal):
		fromEqual := k - len(small)
		return slices.Concat(small, equal[:fromEqual])
	default:
		largElems := k - len(small) - len(equal)
		return slices.Concat(small, equal, quickselect(larg, largElems))
	}
}

func partitionQS(arr []int) ([]int, []int, []int) {
	randIdx := rand.Intn(len(arr))
	pivot := arr[randIdx]
	small, equal, larg := []int{}, []int{}, []int{}
	for _, elem := range arr {
		switch {
		case elem < pivot:
			small = append(small, elem)
		case elem > pivot:
			larg = append(larg, elem)
		default:
			equal = append(equal, elem)
		}
	}
	return small, equal, larg
}
