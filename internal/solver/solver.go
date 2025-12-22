package solver

import (
    "fmt"
)

func Solve(month string, day string) bool {
	board := newBoard()
	board.print()

    fmt.Println("Solved", month, day)
    return true
}
