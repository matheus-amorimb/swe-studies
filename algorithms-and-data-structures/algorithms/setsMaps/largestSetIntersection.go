package main

import (
	"fmt"
	"math"
)

func main() {
	sets := []map[int]struct{}{
		{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		{3: struct{}{}, 2: struct{}{}, 1: struct{}{}},
		{1: struct{}{}, 4: struct{}{}, 5: struct{}{}},
		{1: struct{}{}, 2: struct{}{}},
	}
	idx := largestSetIntersection(sets)
	fmt.Println("idx:", idx)
}

func largestSetIntersection(sets []map[int]struct{}) int {
	n := len(sets)
	if n == 1 {
		return 0
	}

	if n == 2 {
		if len(sets[1]) > len(sets[0]) {
			return 0
		}

		return 1
	}

	elemToFreq := map[int]int{}
	for _, set := range sets {
		for elem := range set {
			elemToFreq[elem]++
		}
	}

	fmt.Println(elemToFreq)

	excludedIdx := 0
	smallestCount := math.MaxInt64
	for idx, set := range sets {
		count := 0
		for elem := range set {
			if elemToFreq[elem] == n-1 {
				count++
			}
		}

		fmt.Println("setIdx:", idx, "|", "count:", count)

		if count < smallestCount {
			smallestCount = count
			excludedIdx = idx
		}
	}

	// inters := sets[0]
	// for _, set := range sets[1:] {
	// 	currInter := map[int]struct{}{}
	// 	for elem := range inters {
	// 		if _, ok := set[elem]; ok {
	// 			currInter[elem] = struct{}{}
	// 		}
	// 	}
	// 	inters = currInter
	// }

	// inters := maps.Clone(sets[0])
	// for _, set := range sets[1:] {
	// 	for elem := range inters {
	// 		if _, ok := set[elem]; !ok {
	// 			delete(inters, elem)
	// 		}
	// 	}
	// }

	// fmt.Println("interArr:", inters)

	return excludedIdx
}
