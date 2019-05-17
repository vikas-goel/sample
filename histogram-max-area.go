package main

import "fmt"

func main() {
	//histogram := Histogram{2, 1, 2, 3, 1}
	histogram := Histogram{1, 4, 2, 5, 6, 3, 2, 6, 6, 5, 2, 1, 3}
	fmt.Println("Area =", maxArea(histogram))
}

func maxArea(bars Histogram) (maxArea int) {
	curBar := 0
	myStack := NewStack(bars.Len())

	evalMaxBarArea := func(index int) {
		thisArea := 0
		lastBar := bars.Height(myStack.Pop())
		if myStack.Empty() {
			thisArea = lastBar * index
		} else {
			thisArea = lastBar * (index - myStack.Peek() - 1)
		}

		if thisArea > maxArea {
			maxArea = thisArea
		}
	}

	for curBar < bars.Len() {
		if myStack.Empty() || bars.Height(myStack.Peek()) < bars.Height(curBar) {
			myStack.Push(curBar)
			curBar += 1
		} else {
			evalMaxBarArea(curBar)
		}
	}

	for !myStack.Empty() {
		evalMaxBarArea(curBar)
	}

	return
}

type Histogram []int

func (this *Histogram) Len() int {
	return len(*this)
}

func (this *Histogram) Height(index int) int {
	return (*this)[index]
}

type Stack struct {
	size int
	elems []int
}

func NewStack(capacity int) *Stack {
	stack := new(Stack)
	stack.elems = make([]int, capacity)
	return stack
}

func (this *Stack) Push(value int) {
	this.elems[this.size] = value
	this.size++
}

func (this *Stack) Pop() int {
	if this.size > 0  {
		this.size--
		value := this.elems[this.size]
		this.elems[this.size] = 0
		return value
	}

	return 0
}

func (this *Stack) Peek() int {
	if this.size > 0  {
		return this.elems[this.size-1]
	}

	return 0
}

func (this *Stack) Len() int {
	return this.size
}

func (this *Stack) Empty() bool {
	if this.size == 0 {
		return true
	}
	return false
}
