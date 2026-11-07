package main

import "fmt"

func main() {
	arr := []int{8, 4, 2, 2}
	k := 2

	values := []int{}
	for _, elem := range compressArrayByK(arr, k) {
		values = append(values, elem.value)
	}
	fmt.Println(values)
}

type Elem struct {
	value int
	freq  int
}

func compressArrayByK(arr []int, k int) []Elem {
	var s Stack[Elem]
	for _, elem := range arr {
		freq := 1
		for s.Size() > 0 {
			top, _ := s.Peek()
			if top.value != elem {
				break
			}

			freq = top.freq + 1
			if freq == k {
				for range k - 1 {
					s.Pop()
				}
				elem *= k
				freq = 1
			}
		}
		s.Push(Elem{
			value: elem,
			freq:  freq,
		})
	}

	return s.arr
}
