package main

import "fmt"

func kthPermutation(array []int, k int) {
	length := len(array)
	if length == 0 || length == 1 || k == 0 || k == 1 {
		return
	}

	factorial := make([]int, length+1)
	factorial[0] = 1
	for i := 1; i <= length; i++ {
		factorial[i] = i*factorial[i-1]
	}

	// If k is bigger than the possible number of permutations, then return
	// the last possible permutation which is reverse of the original input.
	if k > factorial[length] {
		reverse(array)
		return
	}

	k -= 1
	for i := 0; i < length-1; i++ {
		remainingFactorial := factorial[length-1-i]
		firstElementIndex := k/remainingFactorial

		// Move the firstElementIndex to the ith index and shift
		// other elements 1-right.
		shiftRight(array, i, i+firstElementIndex)

		// Number of permutations yet to be taken care of.
		k -= (remainingFactorial*firstElementIndex)
	}
}

func reverse(array []int) {
	length := len(array)
	mid := length/2
	for i := 0; i < mid; i++ {
		array[i], array[length-1-i] = array[length-1-i], array[i]
	}
}

func shiftRight(array []int, start, end int) {
	if start == end {
		return
	}

	moveToFirstElement := array[end]
	for i := end; i > start; i-- {
		array[i] = array[i-1]
	}
	array[start] = moveToFirstElement
}

func main() {
	input := []int{1, 2, 3, 4 }
	fmt.Println(input)
	kthPermutation(input, 16)
	fmt.Println(input)
}
