package main

import "fmt"

func main() {
	arr := []rune{
		'R', 'W', 'B', 'B', 'W', 'R', 'W',
	}
	arr = dutchFlag(arr)
	fmt.Println(arr)
}

func dutchFlag(arr []rune) []rune {

	rCount, wCount, bCount := 0, 0, 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 'R' {
			rCount++
		} else if arr[i] == 'W' {
			wCount++
		} else {
			bCount++
		}
	}

	wIdx := 0
	for i := 0; i < rCount; i++ {
		arr[wIdx] = 'R'
		wIdx++
	}

	for i := 0; i < wCount; i++ {
		arr[wIdx] = 'W'
		wIdx++
	}

	for i := 0; i < bCount; i++ {
		arr[wIdx] = 'B'
		wIdx++
	}

	return arr
}
