package main

import "fmt"

func main() {
	a, b := 9, 2
	ans := waterRefilling(a, b)
	fmt.Println("ans:", ans)
}

// constraint: division operation is not allowed, but it is sill possible divide by power of 2 using right-shift operator.
// use guess-and-check technique

// range [0, max*]
// * use exponential search to find the max
func waterRefilling(a, b int) int {
	max := 1
	for b*max < a {
		max = max * 2
	}

	if max*b == a {
		return max
	}

	l, r := 0, max
	for r-l > 1 {
		// it could result in overflow
		// mid := (l + r) >> 1
		gap := r - l
		halfGap := gap >> 1
		mid := l + halfGap
		if mid*b <= a {
			l = mid
		} else {
			r = mid
		}
	}

	return l
}

//	0 	1 	2 	3 	4
//	l
//					r
//			mid

//	0 	1 	2 	3 	4
//			l
//					r
//				mid

//	0 	1 	2 	3 	4
//				l
//					r
//				mid
