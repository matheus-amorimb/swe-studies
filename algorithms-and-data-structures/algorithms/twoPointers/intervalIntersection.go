package main

import (
	"fmt"
)

type Interval struct {
	start int
	end   int
}

func main() {
	arr1 := [][]int{
		[]int{2, 4},
		[]int{5, 8},
	}
	arr2 := [][]int{
		[]int{3, 3},
		[]int{4, 7},
	}
	intArr := intervalIntersection(arr1, arr2)
	fmt.Println(intArr)
}

func intervalIntersection(arr1, arr2 [][]int) [][]int {
	//parallel pointers

	//cases:
	//int arr1 before int arr2, move to next int arr1
	//int arr2 before int arr1, move to next int arr2
	//overlap, calculate the interval. if int arr1 before int arr2, move int arr1

	p1, p2 := 0, 0
	output := make([][]int, 0)
	for p1 < len(arr1) && p2 < len(arr2) {
		int1, int2 := arr1[p1], arr2[p2]

		if int1[1] < int2[0] {
			p1++
		} else if int2[1] < int1[0] {
			p2++
		} else {
			inter := calculateIntersectionInterval(int1, int2)
			output = append(output, inter)
			if int1[1] < int2[1] {
				p1++
			} else {
				p2++
			}
		}
	}

	return output
}

func calculateIntersectionInterval(arr1, arr2 []int) []int {
	start := max(arr1[0], arr2[0])
	end := min(arr1[1], arr2[1])

	return []int{
		start,
		end,
	}
}
