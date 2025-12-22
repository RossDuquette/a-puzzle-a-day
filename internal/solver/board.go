package solver

import (
	"fmt"
)

const (
	board_height int = 7
	board_width int = 7
)

type Cell struct {
	name string
	covered_by string
}

func (c Cell) showing() string {
	if c.covered_by == vacant_tile {
		return c.name
	}
	return c.covered_by
}

type Board struct {
	cells [board_height][board_width]Cell
}

func newBoard() Board {
	var b Board
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
			cell.covered_by = vacant_tile
		}
	}
	return b
}

func (b Board) print() {
	for row := range board_height {
		for col := range board_width {
			cell := b.cells[row][col]
			fmt.Printf("%-5s", cell.showing())
		}
		fmt.Println()
	}
}
