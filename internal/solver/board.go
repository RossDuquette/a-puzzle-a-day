package solver

import (
	"fmt"
)

const (
	board_height int = 7
	board_width int = 7
)

type Board struct {
	cells [board_height][board_width]Cell
	solution_month string
	solution_day string
}

func create_board(month string, day string) Board {
	var b Board
	b.solution_month = month
	b.solution_day = day
	cell_names := [board_height][board_width]string {
		{ "Jan", "Feb", "Mar", "Apr", "May", "Jun", "" },
		{ "Jul", "Aug", "Sep", "Oct", "Nov", "Dec", "" },
		{ "01",  "02",  "03",  "04",  "05",  "06",  "07" },
		{ "08",  "09",  "10",  "11",  "12",  "13",  "14" },
		{ "15",  "16",  "17",  "18",  "19",  "20",  "21" },
		{ "22",  "23",  "24",  "25",  "26",  "27",  "28" },
		{ "29",  "30",  "31",  "",    "",    "",    "" },
	}
	for row := range board_height {
		for col := range board_width {
			cell := &b.cells[row][col]
			cell.name = cell_names[row][col]
			if cell.name == month || cell.name == day {
				cell.covered_by = cell.name
			} else {
				cell.set_vacant()
			}
		}
	}
	return b
}

func (b Board) copy() Board {
	var new_board Board
	new_board.solution_month = b.solution_month
	new_board.solution_day = b.solution_day
	for row := range board_height {
		for col := range board_width {
			new_board.cells[row][col] = b.cells[row][col]
		}
	}
	return new_board
}

func (b Board) String() string {
	var str string
	for row := range board_height {
		for col := range board_width {
			cell := b.cells[row][col]
			str = fmt.Sprintf("%s%-5s", str, cell.showing())
		}
		str = fmt.Sprintf("%s\n", str)
	}
	return str
}

func (b Board) has_point(point Point) bool {
	return point.x >= 0 &&
		point.x < board_width &&
		point.y >= 0 &&
		point.y < board_height
}

func (b Board) has_tile(tile Tile) bool {
	for row := range board_height {
		for col := range board_width {
			cell := &b.cells[row][col]
			if cell.covered_by == tile.name {
				return true
			}
		}
	}
	return false
}

func (b Board) is_dead() bool {
	// Label free spaces with a group number.
	labels := [board_height][board_width]int{}
	next_label := 1
	for row, line := range b.cells {
		for col, cell := range line {
			if cell.is_free() && (labels[row][col] == 0) {
				b.label_group(&labels, next_label, row, col)
				next_label++
			}
		}
	}
	// Once all empty spaces are labelled, count the number of cells with each label.
	label_counters := make(map[int]int)
	for _, line := range labels {
		for _, label := range line {
			if label != 0 {
				label_counters[label]++
			}
		}
	}
	// If the number is not a multiple of 5 for EVERY group, then the board has no
	// solution and is considered dead.
	for _, count := range label_counters {
		if (count % 5) != 0 {
			return true
		}
	}
	return false
}

func (b Board) label_group(labels *[board_height][board_width]int, label int, row int, col int) {
	if (row < 0) || (row >= board_height) || (col < 0) || (col >= board_width) ||
			!b.cells[row][col].is_free() || (labels[row][col] != 0) {
		return
	}

	labels[row][col] = label

	// Label all adjacent squares
	b.label_group(labels, label, row-1, col)
	b.label_group(labels, label, row+1, col)
	b.label_group(labels, label, row, col-1)
	b.label_group(labels, label, row, col+1)
}
