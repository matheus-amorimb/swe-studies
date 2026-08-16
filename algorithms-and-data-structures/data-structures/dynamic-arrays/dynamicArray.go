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
	d.append(3)
	d.append(4)
	d.append(5)
	d.insert(2, 5)
	fmt.Printf("[")
	for i := 0; i < d.Size; i++ {
		fmt.Printf(" %v ", d.FixedSizeArray[i])
	}
	fmt.Printf("]\n")
}

type DynamicArray struct {
	FixedSizeArray []int
	Size           int
	Capacity       int
}

var defaultCapacity = 10

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
	err := da.checkIndex(i)
	if err != nil {
		return 0, err
	}
	return da.FixedSizeArray[i], nil
}

func (da *DynamicArray) set(i, x int) error {
	err := da.checkIndex(i)
	if err != nil {
		return err
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

func (da *DynamicArray) pop(i int) (int, error) {
	err := da.checkIndex(i)
	if err != nil {
		return 0, err
	}
	elem := da.FixedSizeArray[i]
	for idx := i; idx < da.Size-1; idx++ {
		da.FixedSizeArray[idx] = da.FixedSizeArray[idx+1]
	}
	da.pop_back()

	return elem, nil
}

func (da *DynamicArray) contains(x int) bool {
	for _, elem := range da.FixedSizeArray {
		if elem == x {
			return true
		}
	}
	return false
}

func (da *DynamicArray) insert(i int, x int) error {
	err := da.checkIndex(i)
	if err != nil {
		return err
	}
	da.append(0)
	for idx := da.Size - 1; idx > i; idx-- {
		da.FixedSizeArray[idx] = da.FixedSizeArray[idx-1]
	}
	da.FixedSizeArray[i] = x
	return nil
}

func (da *DynamicArray) checkIndex(i int) error {
	if i < 0 || i >= da.Size {
		return fmt.Errorf("index `%v` out of range for array with lenght: %v", i, da.Size)
	}
	return nil
}
