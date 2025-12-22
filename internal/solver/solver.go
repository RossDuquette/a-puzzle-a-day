package solver

import (
    "fmt"
)

func Solve(month string, day string) bool {
	tiles := get_tiles()
	board := newBoard(month, day)
	for _, tile := range tiles {
		place_tile_on_board(&tile, &board)
	}
	board.print()

    fmt.Println("Solved", month, day)
    return true
}

func place_tile_on_board(tile *Tile, board *Board) bool {
	for row, line := range board.cells {
		for col, cell := range line {
			if cell.is_free() {
				board.cells[row][col].covered_by = tile.name
				return true
			}
		}
	}
	return false
}
