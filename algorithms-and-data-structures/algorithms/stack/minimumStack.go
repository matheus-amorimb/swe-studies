package main

func main() {
	minStack := Constructor()
	minStack.Push(-1)
	minStack.Push(5)
	minStack.Push(0)
	minStack.Push(-5)
	minStack.Pop()
	minStack.Pop()
	minStack.Pop()
	minStack.Pop()
	minStack.Push(4)
	minStack.Push(-4)
	minStack.Push(2)
	minStack.Pop()
	minStack.Pop()
	println()
	println("Get Min: ", minStack.GetMin()) // return 1
}

type MinStack struct {
	stack    []int
	minElems []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	if len(this.minElems) == 0 || val <= this.GetMin() {
		this.minElems = append(this.minElems, val)
	}
}

func (this *MinStack) Pop() {
	top := this.Top()
	this.stack = this.stack[:len(this.stack)-1]

	if top == this.GetMin() {
		this.minElems = this.minElems[:len(this.minElems)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minElems[len(this.minElems)-1]
}

func PrintArray(label string, myArray []int) {
	println()
	print(label, ": ")
	for _, item := range myArray {
		print(item, ", ")
	}
}
