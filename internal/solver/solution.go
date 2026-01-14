package solver

import (
	"fmt"
	"os"
)

func handleSolutions(solutions <-chan Board, saveToFiles bool) {
	numSolutions := 0
	for board := range solutions {
		numSolutions++
		handleSolvedBoard(board, saveToFiles, numSolutions)
	}
	fmt.Println("Found", numSolutions, "total solutions")
}

func handleSolvedBoard(board Board, saveToFiles bool, solutionNum int) {
	if saveToFiles {
		// Create dir
		dir := GetSolutionDir(board.solutionMonth, board.solutionDay)
		os.MkdirAll(dir, 0755)

		// Create file
		path := fmt.Sprintf("%s/%d.txt", dir, solutionNum)
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
