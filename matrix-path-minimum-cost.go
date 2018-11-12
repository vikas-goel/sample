// https://www.geeksforgeeks.org/min-cost-path-dp-6/

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func minimumOf(nums ...int) (min int) {
        min = 1 << 31 -1
        for _, i := range nums {
                if i < min {
                        min = i
                }
        }

        return
}

func minCostPath(costCell [][]int, row, col int) int {
	costPath := make([][]int, row+1)

	costPath[0] = make([]int, col+1)
	copy(costPath[0], costCell[0])
	for r := 1; r <= row; r++ {
		costPath[r] = make([]int, col+1)
		copy(costPath[r], costCell[r][:col+1])
		costPath[r][0] += costPath[r-1][0]
	}

	for c := 1; c <= col; c++ {
		costPath[0][c] += costPath[0][c-1]
	}

	for r := 1; r <= row; r++ {
		for c := 1; c <= col; c++ {
			costPath[r][c] += minimumOf(
				costPath[r-1][c-1],
				costPath[r-1][c],
				costPath[r][c-1])
		}
	}

	return costPath[row][col]
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Usage", os.Args[0], "<row> <column>")
	}

	cost := [][]int{{1,2,3},{4,8,2},{1,5,3}}

	row, _ := strconv.Atoi(os.Args[1])
	column, _ := strconv.Atoi(os.Args[2])
	fmt.Println(minCostPath(cost, row, column))
}
