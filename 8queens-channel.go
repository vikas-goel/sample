package main

import (
	"fmt"
	"time"
)

func main() {
	initBoard := []int{0, 1, 2, 3, 4, 5, 6, 7}

	valid := 0
	start := time.Now()
	ch := make(chan []int)
	wg := make(chan bool)

	go func() {
		defer func(){ wg <- true }()
		cadidateBoards(ch, initBoard, 0, len(initBoard)-1)
	}()

	go func() {
		<- wg
		close(ch)
	}()

	allBoards := collectBoards(ch, &valid)

	fmt.Println("Valid boards:", valid, "of", len(allBoards))
	fmt.Println("Execution time:", time.Since(start))
}

func cadidateBoards(ch chan []int, board []int, start, end int) {
	if start == end {
		newBoard := []int{0, 0, 0, 0, 0, 0, 0, 0}
		copy(newBoard, board)
		ch <- newBoard
	} else {
		for i := start; i <= end; i++ {
			swapPositions(board, start, i)
			cadidateBoards(ch, board, start+1, end)
			swapPositions(board, start, i)
		}
	}
}

func collectBoards(ch chan []int, valid *int) (boards [][]int) {
	for b := range ch {
		boards = append(boards, b)

		if validBoard(b) {
			fmt.Println(b)
			*valid++
		}
	}

	return
}

func validBoard(board []int) bool {
	for i := 0; i < len(board)-1; i++ {
		for j := i+1; j < len(board); j++ {
			if ! validPositions(i, j, board[i], board[j]) {
				return false
			}
		}
	}

	return true
}

func validPositions(row1, row2, col1, col2 int) bool {
	row := int8(row1 - row2)
	col := int8(col1 - col2)

	if (row == col) || (row == -1 * col) {
		return false
	}

	return true
}

func swapPositions(board []int, row1, row2 int) {
	board[row1], board[row2] = board[row2], board[row1]
}
