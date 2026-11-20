package main

import "fmt"

func main() {
	// A tic-tac-toe board: three rows of three cells.
	board := make([][]string, 3)
	for y := range board {
		board[y] = make([]string, 3) // each row is its own slice
		for x := range board[y] {
			board[y][x] = "_"
		}
	}

	board[0][0] = "X"
	board[1][1] = "O"
	board[2][0] = "X"

	for _, row := range board {
		fmt.Println(row)
	}
}
