// https://www.geeksforgeeks.org/search-an-element-in-a-sorted-and-pivoted-array/

package main

import "fmt"

func findPivot(array []int, start, end int) int {
	if end < start {
		return -1
	}

	if start == end {
		return start
	}

	mid := start + (end - start)/2
	if array[mid] > array[mid+1] {
		return mid
	} else if array[mid] < array[mid-1] {
		return mid-1
	}

	if array[start] > array[mid] {
		return findPivot(array, start, mid-1)
	} else {
		return findPivot(array, mid+1, end)
	}
}

func findIndexOf(key int, sortedArray []int, start, end int) int {
	if end < start {
		return -1
	} else if start == end {
		if sortedArray[start] == key {
			return start
		}
		return -1
	}

	mid := start + (end - start)/2
	if sortedArray[mid] == key {
		return mid
	} else if key < sortedArray[mid] {
		return findIndexOf(key, sortedArray, start, mid-1)
	} else  {
		return findIndexOf(key, sortedArray, mid+1, end)
	}
}

func main() {
	array := []int{5, 6, 7, 8, 9, 10, 1, 2, 3}

	pivot := findPivot(array, 0, len(array)-1)

	for index, key := range array {
		if pivot == -1 {
			index = findIndexOf(key, array, 0, len(array)-1)
		} else if key == array[pivot] {
			index = pivot
		} else if key < array[0] {
			index = findIndexOf(key, array, pivot+1, len(array)-1)
		} else {
			index = findIndexOf(key, array, 0, pivot)
		}
		fmt.Println(index)
	}
}
