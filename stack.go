package main

import (
	"fmt"
	"os"
	"strconv"
)

const StackCapacity = 10

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: <prog> <values>...")
		os.Exit(1)
	}

	myStack := new(Stack)
	for i := 1; i < len(os.Args); i++ {
		value, _ := strconv.Atoi(os.Args[i])
		myStack.Push(value)
	}

	fmt.Print(myStack)
	myStack.Sort()
	fmt.Println(myStack)
}

type Stack struct {
	size int
	elems [StackCapacity]int
}

func (this *Stack) Seed(values... int) {
	for _, v := range values {
		this.Push(v)
	}
}

func (this *Stack) Push(value int) {
	if this.size != StackCapacity {
		this.elems[this.size] = value
		this.size++
	}
}

func (this *Stack) Pop() int {
	if this.size > 0  {
		this.size--
		return this.elems[this.size]
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

func (this *Stack) Sort() {
	if this.Len() <= 1 {
		return
	}

	temp := Stack{}
	temp.Push(this.Pop())

	for !this.Empty() {
		for !this.Empty() && this.Peek() >= temp.Peek() {
			temp.Push(this.Pop())
		}

		if this.Empty() {
			break
		}

		topThis := this.Pop()

		for !temp.Empty() && temp.Peek() > topThis {
			this.Push(temp.Pop())
		}

		temp.Push(topThis)
	}

	for !temp.Empty() {
		this.Push(temp.Pop())
	}
}
