package main

import "fmt"

func main() {
	arr := []int{8, 4, 2, 2, 2, 4, 5}
	otp := compressArray(arr)
	fmt.Println(otp.arr)
}

func compressArray(arr []int) Stack[int] {
	var s Stack[int]
	for _, elem := range arr {
		for s.Size() > 0 {
			top, _ := s.Peek()
			if top != elem {
				break
			}

			s.Pop()
			elem += top
		}
		s.Push(elem)
	}

	return s
}
