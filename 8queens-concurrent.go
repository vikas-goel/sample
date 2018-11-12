package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	initBoard := []int{0, 1, 2, 3, 4, 5, 6, 7}

	valid := 0
	start := time.Now()
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		cadidateBoards(&wg, initBoard, 0, len(initBoard)-1, &valid)
	}()

	wg.Wait()

	fmt.Println("Valid boards:", valid)
	fmt.Println("Total time:", time.Since(start))
}

func cadidateBoards(wg *sync.WaitGroup, board []int, start, end int, valid *int) {
	if start == end {
		if validBoard(board) {
			fmt.Println(board)
			*valid++
		}
	} else {
		for i := start; i <= end; i++ {
			newBoard := swapPositions(board, start, i)
			wg.Add(1)
			go func() {
				defer wg.Done()
				cadidateBoards(wg, newBoard, start+1, end, valid)
			}()
		}
	}
}

func validBoard(board []int) bool {
	for i := 0; i < len(board)-1; i++ {
		for j := i+1; j < len(board); j++ {
			row := int8(i - j)
			col := int8(board[i] - board[j])

			if (row == col) || (row == -1 * col) {
				return false
			}
		}
	}

	return true
}

func swapPositions(board []int, row1, row2 int) (newBoard []int) {
	newBoard = []int{0, 0, 0, 0, 0, 0, 0, 0}
	copy(newBoard, board)
	newBoard[row1], newBoard[row2] = board[row2], board[row1]
	return
}
