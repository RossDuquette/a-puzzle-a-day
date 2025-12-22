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
		for col := range line {
			if place_tile_on_board_at(tile, board, row, col) {
				return true
			}
		}
	}
	return false
}

func place_tile_on_board_at(tile *Tile, board *Board, row int, col int) bool {
	points := tile.get_points()
	for _, point := range points {
		point.x += col
		point.y += row
		if !board.has_point(point) {
			return false
		}
		cell := board.cells[point.y][point.x]
		if !cell.is_free() {
			return false
		}
	}
	for _, point := range points {
		point.x += col
		point.y += row
		cell := &board.cells[point.y][point.x]
		cell.covered_by = tile.name
	}
	return true
}
