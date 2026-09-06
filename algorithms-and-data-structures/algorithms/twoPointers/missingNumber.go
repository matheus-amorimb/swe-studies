package main

import "fmt"

func main() {
	arr := []int{6, 9, 12, 15, 18}
	low, high := 9, 13
	missingArr := missingNumber(arr, low, high)
	fmt.Printf("[")
	for _, elem := range missingArr {
		fmt.Printf(" %v ", elem)
	}
	fmt.Printf("]\n")
}

func missingNumber(arr []int, low, high int) []int {
	// pre fill the missing array with the next number from the sequence. this values might or not be in the array.
	// cases:
	// while arr[i] < high or i < len(arr)
	// arr[i] < low, i++
	// arr[i] == low, low++, i++
	// arr[i] > low, append to array low..arr[i]-1, low = arr[i]+1, i++
	i := 0
	output := make([]int, 0)
	for i < len(arr) && arr[i] <= high {
		if arr[i] == low {
			low++
		} else if arr[i] > low {
			for missNum := low; missNum < arr[i]; missNum++ {
				output = append(output, missNum)
			}
			low = arr[i] + 1
		}
		i++
	}

	for low <= high {
		output = append(output, low)
		low++
	}

	return output
}

func missingNumberCleverSolution(arr []int, low, high int) []int {
	output := make([]int, 0)
	i, num := 0, low

	for num <= high {
		if i < len(arr) && arr[i] < num {
			i++
		} else if i < len(arr) && arr[i] == num {
			i++
			num++
		} else {
			output = append(output, num)
			num++
		}

	}

	return output
}
