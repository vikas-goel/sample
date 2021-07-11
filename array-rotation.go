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
	rotateRight(array, -1*by)
}

func rotateRight(array []int, by int) {
	length := len(array)

	if length == 0 {
		return
	}

	// Normalize rotation by
	by = by % length
	if by == 0 {
		return
	}

	// Left rotation
	if by < 0 {
		by += length
	}

	cIndex, cValue := 0, array[0]
	for count := length; count > 0; count-- {
		nIndex := (cIndex + by) % length
		cValue, array[nIndex] = array[nIndex], cValue
		cIndex = nIndex
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	by := 4

	fmt.Printf("%v <- %v = ", a, by)
	rotateLeft(a, by)
	fmt.Println(a)

	by = 4
	fmt.Printf("%v -> %v = ", a, by)
	rotateRight(a, by)
	fmt.Println(a)

	fmt.Printf("%v <--> = ", a)
	reverse(a)
	fmt.Println(a)
}
