// https://www.geeksforgeeks.org/sort-array-applying-given-equation/

package main

import (
	"fmt"
	"sort"
)

type equation struct {
	value, index int
}

type Equations []equation

func (e Equations) Len() int {
	return len(e)
}

func (e Equations) Less(i, j int) bool {
	if i == j {
		return true
	}

	if e[j].value < e[i].value {
		return false
	}

	return true
}

func (e Equations) Swap(i, j int) {
	e[i].value, e[j].value = e[j].value, e[i].value
	e[i].index, e[j].index = e[j].index, e[i].index
}

func (e Equations) Sort() {
	sort.Sort(e)
}

func sortArray(array []int, length, A, B, C int) {
	eqarray := make([]equation, length)

	eval := func(num int) int {
		return A*num*num + B*num + C
	}

	for i, v := range array {
		eqarray[i].value, eqarray[i].index = eval(v), i
	}

	sort.Sort(Equations(eqarray))

	fmt.Print(array, eqarray)

	fmt.Print(" [")
	for i := 0; i < len(eqarray)-1; i++ {
		fmt.Print(array[eqarray[i].index], " ")
	}
	fmt.Print(array[eqarray[len(eqarray)-1].index])
	fmt.Println("]")
}

func main() {
	arr := []int{-1, 0, 1, 2, 3, 4}
	sortArray(arr, len(arr), -1, 2, -1)
}
