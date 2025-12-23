package solver

import (
	"sync"
)

func Solve(month string, day string, save_to_files bool) {
	board := newBoard(month, day)
	solutions := make(chan Board)
	var threads sync.WaitGroup
	threads.Go(func() {
		handle_solutions(solutions, save_to_files)
	})
	place_first_tile_on(&board, solutions)
	close(solutions)
	threads.Wait()
}

func place_first_tile_on(board *Board, solutions chan<- Board) {
	orig_board := board.copy()
	tiles := get_tiles()
	tile := tiles["b"]
	delete(tiles, "b")
	var threads sync.WaitGroup
	for range tile.num_rotations {
		for row, line := range board.cells {
			for col := range line {
				if place_tile_on_board_at(&tile, board, row, col) {
					new_board := board.copy()
					threads.Go(func() {
						place_next_tile_on_board(tiles, &new_board, solutions)
					})
					*board = orig_board.copy()
				}
			}
		}
		tile.rotate_cw(1)
	}
	threads.Wait()
}

func place_next_tile_on_board(tiles map[string]Tile, board *Board, solutions chan<- Board) {
	if board.is_dead() {
		return
	}
	tile := get_next_tile(tiles, board)
	orig_board := board.copy()
	for range tile.num_flips {
		for range tile.num_rotations {
			for row, line := range board.cells {
				for col := range line {
					if place_tile_on_board_at(&tile, board, row, col) {
						if is_solved(board) {
							solutions <- *board
						} else {
							place_next_tile_on_board(tiles, board, solutions)
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

func get_next_tile(tiles map[string]Tile, board *Board) Tile {
	for _, tile := range tiles {
		if !board.has_tile(tile) {
			return tile
		}
	}
	return tiles["b"]
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
