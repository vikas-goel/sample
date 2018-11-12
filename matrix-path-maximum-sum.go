// https://www.geeksforgeeks.org/maximum-path-sum-starting-cell-0-th-row-ending-cell-n-1-th-row/

package main

import "fmt"

func max(nums ...int) (maxNum int) {
	for _, i := range nums {
		if i > maxNum {
			maxNum = i
		}
	}

	return
}

func matrixMax(matrix [][]int) (largestSum int) {
	dimension := len(matrix)

	path := make([][]int, dimension, dimension)
	for i := 0; i < dimension; i++ {
		path[i] = make([]int, dimension+2)
		path[0][i+1] = matrix[0][i]
	}

	for i := 1; i < dimension; i++ {
		for j := 1; j <= dimension; j++ {
			path[i][j] = max(path[i-1][j], path[i-1][j-1], path[i-1][j+1]) + matrix[i][j-1]
		}
	}

	fmt.Println(path)

	for i := 1; i <= dimension; i++ {
		if path[dimension-1][i] > largestSum {
			largestSum = path[dimension-1][i]
		}
	}

	return
}

func main() {
	matrix := [][]int{{4, 2, 3, 4}, {2, 9, 1, 10}, {15, 1, 3, 0}, {16, 92, 41, 44}}
	fmt.Println("Largest sum = ", matrixMax(matrix))
}
