package solver

import (
	"fmt"
)

const (
	boardHeight int = 7
	boardWidth  int = 7
)

type Board struct {
	cells         [boardHeight][boardWidth]Cell
	solutionMonth string
	solutionDay   string
}

func createBoard(month string, day string) Board {
	var b Board
	b.solutionMonth = month
	b.solutionDay = day
	cellNames := [boardHeight][boardWidth]string{
		{"Jan", "Feb", "Mar", "Apr", "May", "Jun", ""},
		{"Jul", "Aug", "Sep", "Oct", "Nov", "Dec", ""},
		{"01", "02", "03", "04", "05", "06", "07"},
		{"08", "09", "10", "11", "12", "13", "14"},
		{"15", "16", "17", "18", "19", "20", "21"},
		{"22", "23", "24", "25", "26", "27", "28"},
		{"29", "30", "31", "", "", "", ""},
	}
	for row := range boardHeight {
		for col := range boardWidth {
			cell := &b.cells[row][col]
			cell.name = cellNames[row][col]
			if cell.name == month || cell.name == day {
				cell.coveredBy = cell.name
			} else {
				cell.setVacant()
			}
		}
	}
	return b
}

func (b Board) copy() Board {
	var newBoard Board
	newBoard.solutionMonth = b.solutionMonth
	newBoard.solutionDay = b.solutionDay
	for row := range boardHeight {
		for col := range boardWidth {
			newBoard.cells[row][col] = b.cells[row][col]
		}
	}
	return newBoard
}

func (b Board) String() string {
	var str string
	for row := range boardHeight {
		for col := range boardWidth {
			cell := b.cells[row][col]
			str = fmt.Sprintf("%s%-5s", str, cell.showing())
		}
		str = fmt.Sprintf("%s\n", str)
	}
	return str
}

func (b Board) hasPoint(point Point) bool {
	return point.x >= 0 &&
		point.x < boardWidth &&
		point.y >= 0 &&
		point.y < boardHeight
}

func (b Board) hasTile(tile Tile) bool {
	for row := range boardHeight {
		for col := range boardWidth {
			cell := &b.cells[row][col]
			if cell.coveredBy == tile.name {
				return true
			}
		}
	}
	return false
}

func (b Board) isDead() bool {
	// Label free spaces with a group number.
	labels := [boardHeight][boardWidth]int{}
	nextLabel := 1
	for row, line := range b.cells {
		for col, cell := range line {
			if cell.isFree() && (labels[row][col] == 0) {
				b.labelGroup(&labels, nextLabel, row, col)
				nextLabel++
			}
		}
	}
	// Once all empty spaces are labelled, count the number of cells with each label.
	labelCounters := make(map[int]int)
	for _, line := range labels {
		for _, label := range line {
			if label != 0 {
				labelCounters[label]++
			}
		}
	}
	// If the number is not a multiple of 5 for EVERY group, then the board has no
	// solution and is considered dead.
	for _, count := range labelCounters {
		if (count % 5) != 0 {
			return true
		}
	}
	return false
}

func (b Board) labelGroup(labels *[boardHeight][boardWidth]int, label int, row int, col int) {
	if (row < 0) || (row >= boardHeight) || (col < 0) || (col >= boardWidth) ||
		!b.cells[row][col].isFree() || (labels[row][col] != 0) {
		return
	}

	labels[row][col] = label

	// Label all adjacent squares
	b.labelGroup(labels, label, row-1, col)
	b.labelGroup(labels, label, row+1, col)
	b.labelGroup(labels, label, row, col-1)
	b.labelGroup(labels, label, row, col+1)
}
