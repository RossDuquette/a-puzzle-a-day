package solver

import (
	"fmt"
)

func handle_solutions(solutions <-chan Board, save_to_files bool) {
	num_solutions := 0
	for board := range solutions {
		handle_solved_board(board, save_to_files)
		num_solutions++
	}
	fmt.Println("Found", num_solutions, "total solutions")
}

func handle_solved_board(board Board, save_to_files bool) {
	fmt.Println("Solved!")
	fmt.Println(board)
}
