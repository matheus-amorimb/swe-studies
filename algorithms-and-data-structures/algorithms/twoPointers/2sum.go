package main

import "fmt"

func main() {
	//sorted array
	arr1 := []int{-5, -2, -1, 1, 1, 10}
	arr2 := []int{-3, 0, 0, 1, 2}
	arr3 := []int{-5, -3, -1, 0, 2, 4, 6}
	isTwoSum1 := twoSum(arr1)
	isTwoSum2 := twoSum(arr2)
	isTwoSum3 := twoSum(arr3)
	fmt.Printf("isTwoSum1: %v | expected: true\n", isTwoSum1)
	fmt.Printf("isTwoSum2: %v | expected: true\n", isTwoSum2)
	fmt.Printf("isTwoSum2: %v | expected: false\n", isTwoSum3)
}

func twoSum(arr []int) bool {
	// inward pointer
	// loop invariant: arr[lP] + arr[rP] are different from zero for any value of rP...len(arr)-1 and lP...len(arr)
	// use cases:
	// arr[lp] + arr[rP] > 0 => we must decrease the sum, therefore rP--
	// arr[lp] + arr[rP] < 0 => we must increase the sum, therefore lP++
	// arr[lp] + arr[rP] == 0 => return true
	// lp < rP

	lP, rP := 0, len(arr)-1
	for lP < rP {
		lElem := arr[lP]
		rElem := arr[rP]
		twoSum := lElem + rElem

		if twoSum > 0 {
			rP--
		} else if twoSum < 0 {
			lP++
		} else {
			return true
		}

	}

	return false
}
