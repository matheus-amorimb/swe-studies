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
	fmt.Println(d.get(0))
	fmt.Println(d.get(1))
	fmt.Println(d.size())
}

type DynamicArray struct {
	FixedSizeArray []int
	Size           int
	Capacity       int
}

var defaultCapacity = 100

func NewDynamicArray() *DynamicArray {
	a := make([]int, defaultCapacity)
	return &DynamicArray{
		FixedSizeArray: a,
		Size:           0,
		Capacity:       defaultCapacity,
	}
}

func (da *DynamicArray) append(x int) {
	if da.Size == da.Capacity {
		da.resize(2 * da.Capacity)
	}

	da.FixedSizeArray[da.Size] = x
	da.Size++
}

func (da *DynamicArray) get(i int) (int, error) {
	if i < 0 || i >= da.Size {
		return 0, fmt.Errorf("index `%v` out of range for array with lenght: %v", i, da.Size)
	}
	return da.FixedSizeArray[i], nil
}

func (da *DynamicArray) set(i, x int) error {
	if i < 0 || i >= da.Size {
		return fmt.Errorf("index `%v` out of range for array with lenght: %v", i, da.Size)
	}
	da.FixedSizeArray[i] = x
	return nil
}

func (da *DynamicArray) size() int {
	return da.Size
}

func (da *DynamicArray) pop_back() error {
	if da.Size == 0 {
		return fmt.Errorf("array is already empty.")
	}
	da.Size--
	if float64(da.Size/da.Capacity) < 0.25 && da.Capacity > defaultCapacity {
		da.resize(da.Capacity / 2)
	}
	return nil
}

func (da *DynamicArray) resize(newSize int) {
	newA := make([]int, newSize)
	for i, el := range da.FixedSizeArray {
		newA[i] = el
	}
	da.FixedSizeArray = newA
	da.Capacity = newSize
}
