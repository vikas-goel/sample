package main

import "fmt"

type sortMethod func([]int, int) (string, int, int)

func main() {
	sortIntegersAll(11, 12, 22, 25, 34, 64, 90)
	sortIntegersAll(64, 34, 25, 12, 22, 11, 90)
	sortIntegersAll(90, 64, 34, 25, 22, 12, 11)
	sortIntegersAll(10, 7, 8, 9, 1, 5)
	sortIntegersAll(1, 9, 5, 13, 3, 11, 7, 15, 2, 10, 6, 14, 4, 12, 16, 8)
}

func sortIntegersAll(nums ...int) {
	for _, sm := range []sortMethod{selection, bubble, insertion, merge, quick, heap} {
		sortIntegers(sm, nums...)
	}
	fmt.Println()
}

func sortIntegers(fn sortMethod, nums ...int) {
	array := make([]int, len(nums))
	for i, n := range nums {
		array[i] = n
	}

	fmt.Print(array, "->")
	name, iters, swaps := fn(array, len(array))
	fmt.Printf("%v:(%2v,%2v):%v sort\n", array, iters, swaps, name)
}

func heap(array []int, length int) (name string, numIter, numSwap int) {
	name = "Heap"

	var heapify func([]int, int, int) (int, int)
	heapify = func(array []int, length, index int) (iters, swaps int) {
		if index < 0 || index >= length {
			return
		}

		lchild, rchild := 2*index+1, 2*index+2
		max := index
		if lchild < length && array[lchild] > array[max] {
			max = lchild
		}

		if rchild < length && array[rchild] > array[max] {
			max = rchild
		}

		if index != max {
			array[max], array[index] = array[index], array[max]
			iters, swaps = heapify(array, length, max)
			swaps++
		}

		iters++
		return
	}

	for i := length/2-1; i >= 0; i-- {
		it, sw := heapify(array, length, i)
		numIter += it
		numSwap += sw
	}

	for i := length-1; i >= 0; i-- {
		array[i], array[0] = array[0], array[i]
		it, sw := heapify(array, i, 0)
		numIter += it
		numSwap += sw
	}

	return
}

func quick(array []int, length int) (name string, numIter, numSwap int) {
	name = "Quick"

	partition := func(array []int, start, end int) (pi, iters, swaps int) {
		pi = start

		if start == end {
			return
		}

		pivot := array[end]
		for j := start; j < end; j++ {
			if array[j] <= pivot {
				array[pi], array[j] = array[j], array[pi]
				pi++
				swaps++
			}
			iters++
		}
		array[pi], array[end] = array[end], array[pi]
		swaps++
		return
	}
	
	var splitSort func([]int, int, int) (int, int)
	splitSort = func(array []int, start, end int) (iters, swaps int) {
		if start >= end {
			return
		}

		pi, it, sw := partition(array, start, end)
		iters = it
		swaps = sw

		it, sw = splitSort(array, 0, pi-1)
		iters += it
		swaps += sw

		it, sw = splitSort(array, pi+1, end)
		iters += it
		swaps += sw

		return
	}

	numIter, numSwap = splitSort(array, 0, length-1)
	return
}

func merge(array []int, length int) (name string, numIter, numSwap int) {
	name = "Merge"

	var splitSort func([]int, int, int) (int, int)
	splitSort = func(array []int, start, end int) (iters, swaps int) {
		if start >= end {
			return
		}

		middle := start + (end-start)/2
		i1, s1 := splitSort(array, start, middle)
		i2, s2 := splitSort(array, middle+1, end)

		iters = i1+i2
		swaps = s1+s2

		size1, size2 := middle-start+1, end-middle
		temp1, temp2 := make([]int, size1), make([]int, size2)
		copy(temp1, array[start:middle+1])
		copy(temp2, array[middle+1:end+1])
		t1, t2 := 0, 0
		for ; t1 < size1 && t2 < size2; start++ {
			iters++
			swaps++

			if temp2[t2] < temp1[t1] {
				array[start] = temp2[t2]
				t2++
			} else {
				array[start] = temp1[t1]
				t1++
			}
		}

		for _, elem := range temp1[t1:] {
			array[start] = elem
			start++
			iters++
			swaps++
		}

		for _, elem := range temp2[t2:] {
			array[start] = elem
			start++
			iters++
			swaps++
		}

		return
	}

	numIter, numSwap = splitSort(array, 0, length-1)
	return
}

func insertion(array []int, length int) (name string, numIter, numSwap int) {
	name = "Insertion"
	for i := 1; i < length; i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			numIter++
			numSwap++
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
	return
}

func bubble(array []int, length int) (name string, numIter, numSwap int) {
	name = "Bubble"
	for sorted := false; !sorted; {
		sorted = true
		length--
		for i := 0; i < length; i++ {
			numIter++
			if array[i] > array[i+1] {
				numSwap++
				array[i], array[i+1] = array[i+1], array[i]
				sorted = false
			}
		}
	}
	return
}

func selection(array []int, length int) (name string, numIter, numSwap int) {
	name = "Selection"
	for i := 0; i < length-1; i++ {
		min := i
		for j := i+1; j < length; j++ {
			numIter++
			if array[j] < array[min] {
				min = j
			}
		}

		if i != min {
			numSwap++
			array[i], array[min] = array[min], array[i]
		}
	}
	return
}
