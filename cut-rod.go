// https://www.geeksforgeeks.org/dynamic-programming-set-13-cutting-a-rod/

package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func cutRod(length int, price []int) int {
	M := make([]int, length+1)

	for i := 1; i <= length; i++ {
		for j := 0; j < i; j++ {
			M[i] = max(M[i], price[j]+M[i-j-1])
		}
	}

	fmt.Println(M)
	return M[length]
}

func main() {
	price := []int{1, 5, 8, 9, 10, 17, 17, 20}
	fmt.Println(cutRod(len(price), price))
}
