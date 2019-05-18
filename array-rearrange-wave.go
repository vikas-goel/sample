package main

import "fmt"

func rearrange(array []int) {
	n := len(array)
	for i := 1; i < n; i++ {
		if (i % 2 == 0 && array[i-1] < array[i]) ||
			(i % 2 != 0 && array[i-1] > array[i]) {
				array[i-1], array[i] = array[i], array[i-1]
		}
	}
}

func main() {
	a := []int{2, 5, 8, 3, 10, 15, 12, 11, 9, 4, 5, 6, 7}
	fmt.Print(a)
	rearrange(a)
	fmt.Print(a)
}
