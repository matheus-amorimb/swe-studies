package main

import "fmt"

func main() {
	arr := []int{2, 2}
	target := 2
	k := 3
	match := targetCountDivisibleByK(arr, target, k)
	fmt.Println(match)
}

// before < target
//	1	|	2	2	2	2	2	2	3

// before < target + 1
//	1	2	2	2	2	2	2	|	3

func targetCountDivisibleByK(arr []int, target, k int) bool {
	if len(arr) == 0 {
		return true
	}
	l, r := 0, len(arr)-1
	if arr[l] == target {
		r = l
	}
	for r-l > 1 {
		mid := (r + l) / 2
		if arr[mid] < target {
			l = mid
		} else {
			r = mid
		}
	}

	firstIdx := r

	if arr[firstIdx] != target {
		return true
	}

	l, r = 0, len(arr)-1
	if arr[r] == target {
		l = r
	}
	for r-l > 1 {
		mid := (r + l) / 2
		if arr[mid] <= target {
			l = mid
		} else {
			r = mid
		}
	}

	lastIdx := l

	occurs := lastIdx - firstIdx + 1
	fmt.Println("firstIdx:", firstIdx)
	fmt.Println("lastIdx:", lastIdx)
	fmt.Println("occurs:", occurs)
	fmt.Println("remainder:", occurs%k)
	return occurs%k == 0
}
