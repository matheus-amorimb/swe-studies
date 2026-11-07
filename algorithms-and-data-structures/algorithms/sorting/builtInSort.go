package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	// descAscOrder()
	// specificIdx()
	// nestedArr()
	// fieldObject()
	// fieldObjectWithTiebreak()
	fieldObjectWithTiebreakOriginalIdx()
}

func descAscOrder() {
	arr := []string{"obama", "arthur", "vincent", "bryan", "nora"}
	desc := func(a, b string) int {
		return cmp.Compare(b, a)
	}
	slices.SortFunc(arr, desc)
	fmt.Println("desc:", arr)

	slices.Sort(arr)
	fmt.Println("asc:", arr)
}

func specificIdx() {
	arr := []string{"o", "acthur", "vdncent", "beyan", "nara", "a"}
	specIdx := func(a, b string) int {
		idx := 1
		if len(a) <= idx {
			return 1
		}
		if len(b) <= idx {
			return -1
		}
		return cmp.Compare(a[idx], b[idx])
	}
	slices.SortFunc(arr, specIdx)
	fmt.Println("idx1:", arr)
}

func nestedArr() {
	arr := [][]string{
		{"matheus", "batista"},
		{"joão", "pedro"},
		{"maria", "clara"},
		{"ana", "maria"},
		{"carlos", "alberto"},
	}
	bySecond := func(a, b []string) int {
		idx := 1
		if len(a) <= idx && len(b) <= idx {
			return 0
		}
		if len(a) <= idx {
			return 1
		}

		if len(b) <= idx {
			return -1
		}

		return cmp.Compare(a[idx], b[idx])
	}
	slices.SortFunc(arr, bySecond)
	fmt.Println(arr)
}

type User struct {
	name string
	age  int
}

func fieldObject() {
	arr := []User{
		{name: "matheus", age: 28},
		{name: "brenda", age: 25},
		{name: "ana", age: 48},
		{name: "nora", age: 0},
	}
	byAge := func(a, b User) int {
		return a.age - b.age
	}
	slices.SortFunc(arr, byAge)
	fmt.Println(arr)
}

func fieldObjectWithTiebreak() {
	arr := []User{
		{name: "matheus", age: 28},
		{name: "brenda", age: 25},
		{name: "ana", age: 48},
		{name: "nora", age: 0},
		{name: "arthur", age: 28},
	}
	byAge := func(a, b User) int {
		diff := a.age - b.age
		if diff == 0 {
			return cmp.Compare(a.name, b.name)
		}
		return diff
	}
	// byAgeIdiomatic := func(a, b User) int {
	// 	return cmp.Or(
	// 		cmp.Compare(a.age, b.age),
	// 		cmp.Compare(a.name, b.name),
	// 	)
	// }
	slices.SortFunc(arr, byAge)
	fmt.Println(arr)
}

type indexed struct {
	u   User
	idx int
}

func fieldObjectWithTiebreakOriginalIdx() {
	arr := []User{
		{name: "arthur", age: 28},
		{name: "matheus", age: 28},
		{name: "brenda", age: 25},
		{name: "ana", age: 48},
		{name: "nora", age: 0},
	}

	arrIndexed := make([]indexed, len(arr))
	for idx, user := range arr {
		arrIndexed[idx] = indexed{u: user, idx: idx}
	}

	byAgeTiebreakToIdx := func(a, b indexed) int {
		return cmp.Or(
			cmp.Compare(a.u.age, b.u.age),
			cmp.Compare(a.idx, b.idx),
		)
	}

	slices.SortFunc(arrIndexed, byAgeTiebreakToIdx)
	fmt.Println(arrIndexed)
}
