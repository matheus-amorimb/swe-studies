package main

import "fmt"

func main() {
	arr := []int{-2, 0, 3, 4, 7, 9, 11}
	search := 1
	fmt.Println(binarySearch(search, arr))
}

//target: 7
//-2	0	3	4	7	9	11
//l
//							r
//				mid

//-2	0	3	4	7	9	11
//l
//				r
//		mid

//-2	0	3	4	7	9	11
//		l
//				r
//			mid

//-2	0	3	4	7	9	11
//		l
//			r
//			mid
//exit loop

func binarySearch(target int, arr []int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	l, r := 0, len(arr)-1
	if arr[l] >= target || arr[r] < target {
		if arr[l] == target {
			return 0
		}
		return -1
	}
	for l+1 < r {
		mid := int((l + r) / 2)
		if arr[mid] < target {
			l = mid
		} else {
			r = mid
		}
	}

	if arr[r] == target {
		return r
	}

	return -1
}
