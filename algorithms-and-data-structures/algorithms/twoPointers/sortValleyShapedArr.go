package main

import "fmt"

func main() {
	arr := []int{9, 7, 5, 3, 1, 2}
	sortArr := sortValleySharpedArray(arr)
	fmt.Printf("[")
	for _, elem := range sortArr {
		fmt.Printf(" %v ", elem)
	}
	fmt.Printf("]\n")
}

func sortValleySharpedArray(arr []int) []int {
	//inward pointes
	//arr[l] and arr[r] will always be equal or greater than arr[l+n] and arr[r-m], while l <= r
	//cases:
	//start fill in output array from back to to front, optIdx = len(arr) - 1
	//arr[l] >= arr[r] => opt[optIdx] = arr[l], l++, optIdx--
	//arr[r] > arr[l] => opt[optIdx] = arr[r], r--, optIdx--
	lP, rP := 0, len(arr)-1
	output := make([]int, len(arr))
	currOtpIdx := len(arr) - 1
	for lP < rP {
		if arr[lP] >= arr[rP] {
			output[currOtpIdx] = arr[lP]
			lP++
			currOtpIdx--
		} else {
			output[currOtpIdx] = arr[rP]
			rP--
			currOtpIdx--
		}
	}

	return output
}
