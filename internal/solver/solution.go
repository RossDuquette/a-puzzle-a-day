package solver

import (
	"fmt"
)

func handle_solutions(solutions chan Board) {
	num_solutions := 0
	for board := range solutions {
		handle_solved_board(board)
		num_solutions++
	}
	fmt.Println("Found", num_solutions, "total solutions")
}

func handle_solved_board(board Board) {
	fmt.Println("Solved!")
	fmt.Println(board)
}
