package solver

import (
	"fmt"
	"os"
)

func handle_solutions(solutions <-chan Board, save_to_files bool) {
	num_solutions := 0
	for board := range solutions {
		num_solutions++
		handle_solved_board(board, save_to_files, num_solutions)
	}
	fmt.Println("Found", num_solutions, "total solutions")
}

func handle_solved_board(board Board, save_to_files bool, solution_num int) {
	if save_to_files {
		// Create dir
		dir := GetSolutionDir(board.solution_month, board.solution_day)
		os.MkdirAll(dir, 0755)

		// Create file
		path := fmt.Sprintf("%s/%d.txt", dir, solution_num)
		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		file.WriteString(board.String())
	} else {
		fmt.Println("Solved!")
		fmt.Println(board)
	}
}

func GetSolutionDir(month string, day string) string {
	return fmt.Sprintf("solutions/%s-%s", month, day)
}
