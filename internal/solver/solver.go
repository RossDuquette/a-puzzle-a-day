package solver

import (
	"fmt"
	"sync"
)

func Solve(month string, day string) bool {
	board := newBoard(month, day)
	place_first_tile_on(&board)
    return true
}

func place_first_tile_on(board *Board) {
	orig_board := board.copy()
	tiles := get_tiles()
	var threads sync.WaitGroup
	for _, tile := range tiles {
		for range tile.num_flips {
			for range tile.num_rotations {
				for row, line := range board.cells {
					for col := range line {
						if place_tile_on_board_at(&tile, board, row, col) {
							new_board := board.copy()
							threads.Go(func() {
								place_tiles_on(&new_board)
							})
							*board = orig_board.copy()
						}
					}
				}
				tile.rotate_cw(1)
			}
			tile.flip()
		}
	}
	threads.Wait()
}

func place_tiles_on(board *Board) {
	orig_board := board.copy()
	tiles := get_tiles()
	for _, tile := range tiles {
		if board.has_tile(tile) {
			continue
		}
		for range tile.num_flips {
			for range tile.num_rotations {
				for row, line := range board.cells {
					for col := range line {
						if place_tile_on_board_at(&tile, board, row, col) {
							if is_solved(board) {
								fmt.Println("Solved!")
								board.print()
							} else {
								place_tiles_on(board)
								*board = orig_board.copy()
							}
						}
					}
				}
				tile.rotate_cw(1)
			}
			tile.flip()
		}
	}
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

func is_solved(board *Board) bool {
	for _, line := range board.cells {
		for _, cell := range line {
			if cell.is_free() {
				return false
			}
		}
	}
	return true
}
