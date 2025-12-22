package solver

import (
	"fmt"
)

const (
	height int = 7
	width int = 7
)

type cell struct {
	name string
	covered_by tile_id
}

func (c cell) print() {
	if c.covered_by == tile_vacant {
		fmt.Printf("%-5s", c.name)
	} else {
		fmt.Printf("%-5s", c.covered_by)
	}
}

type board struct {
	cells [height][width]cell
}

func newBoard() *board {
	var b board
	cell_names := [height][width]string {
		{ "Jan", "Feb", "Mar", "Apr", "May", "Jun", "" },
		{ "Jul", "Aug", "Sep", "Oct", "Nov", "Dec", "" },
		{ "01",  "02",  "03",  "04",  "05",  "06",  "07" },
		{ "08",  "09",  "10",  "11",  "12",  "13",  "14" },
		{ "15",  "16",  "17",  "18",  "19",  "20",  "21" },
		{ "22",  "23",  "24",  "25",  "26",  "27",  "28" },
		{ "29",  "30",  "31",  "",    "",    "",    "" },
	}
	for row := range height {
		for col := range width {
			cell := &b.cells[row][col]
			cell.name = cell_names[row][col]
			cell.covered_by = tile_vacant
		}
	}
	return &b
}

func (b board) print() {
	for row := range height {
		for col := range width {
			b.cells[row][col].print()
		}
		fmt.Println()
	}
}
