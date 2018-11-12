// https://www.geeksforgeeks.org/array-rotation/

package main

import "fmt"

func reverse(array []int) {
	n := len(array)
	for i, m := 0, (n-1)/2; i < m; i++ {
		array[i], array[n-1-i] = array[n-1-i], array[i]
	}
}

func rotateLeft(array []int, by int) {
	length := len(array)
	if by % length == 0 {
		return
	}

	rotate(array, length, length-by, 0)
}

func rotateRight(array []int, by int) {
	length := len(array)
	if by % length == 0 {
		return
	}

	rotate(array, length, by, 0)
}

func rotate(array []int, length, by, pos int) {
	if length == 0 || pos == length {
		return
	}

	elem := array[pos]
	rotate(array, length, by, pos+1)
	array[(pos + by) % length] = elem
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Print(a)
	rotateLeft(a, 4)
	fmt.Print(" -> ", a)
	rotateRight(a, 4)
	fmt.Print(" -> ", a)
	reverse(a)
	fmt.Println(" ->", a)
}
