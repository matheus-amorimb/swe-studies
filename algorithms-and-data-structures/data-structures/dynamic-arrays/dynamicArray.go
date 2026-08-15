package main

import "fmt"

// # Implement Dynamic Array

// Assume your programming language only supports fixed-size arrays. Implement a dynamic array data structure that supports the following:

// `Dynamic Array API:`

// - `append(x)`: adds element `x` to the end of the array
// - `get(i)`: returns the element at index `i`
// - `set(i, x)`: updates the preexisting element at index `i` to be `x`
// - `size()`: returns the number of elements in the array
// - `pop_back()`: removes the last element

// You should only declare arrays of a fixed size and not use built-in `append()` methods or equivalent.

// In Go:
// slices: Slices are pointers to arrays.
// arrays: Arrays are fixed-length at *compile time*.

func main() {
	d := NewDynamicArray()
	d.append(1)
	d.append(2)
	// d.pop_back()
	// d.pop_back()
	fmt.Println(d.get(0))
	fmt.Println(d.get(1))
	fmt.Println(d.size())
}

type DynamicArray struct {
	a []int
}

func NewDynamicArray() *DynamicArray {
	a := make([]int, 0, 100)
	return &DynamicArray{
		a: a,
	}
}

func (da *DynamicArray) append(x int) {
	curLen := len(da.a)
	curCap := cap(da.a)

	newLen := curLen + 1
	if curLen == curCap {
		newA := make([]int, newLen, 2*newLen)
		for i, el := range da.a {
			newA[i] = el
		}
		newA[curLen] = x
		da.a = newA
		return
	}

	da.a = da.a[:newLen]
	da.a[curLen] = x
}

func (da *DynamicArray) get(i int) (int, error) {
	if i < 0 || i >= len(da.a) {
		return 0, fmt.Errorf("index `%v` out of range for array with lenght: %v", i, len(da.a))
	}
	return da.a[i], nil
}

func (da *DynamicArray) set(i, x int) error {
	if i < 0 || i >= len(da.a) {
		return fmt.Errorf("index `%v` out of range for array with lenght: %v", i, len(da.a))
	}
	da.a[i] = x
	return nil
}

func (da *DynamicArray) size() int {
	return len(da.a)
}

func (da *DynamicArray) pop_back() {
	newLen := len(da.a) - 1
	if newLen >= 0 {
		// array with lenght 4, idx goes from 0 to 3
		// to shrink by 1, i should slice from 0 to 3
		// since go slice bounds are half-open: a[low:high] includes low up to high-1
		da.a = da.a[0:newLen]
	}
}
