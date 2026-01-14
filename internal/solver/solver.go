package solver

import (
	"sync"
)

func Solve(month string, day string, saveToFiles bool) {
	board := createBoard(month, day)
	solutions := make(chan Board)
	var threads sync.WaitGroup
	threads.Go(func() {
		handleSolutions(solutions, saveToFiles)
	})
	placeFirstTileOn(&board, solutions)
	close(solutions)
	threads.Wait()
}

func placeFirstTileOn(board *Board, solutions chan<- Board) {
	origBoard := board.copy()
	tiles := getTiles()
	tile := tiles["b"]
	delete(tiles, "b")
	var threads sync.WaitGroup
	for range tile.numRotations {
		for row, line := range board.cells {
			for col := range line {
				if placeTileOnBoardAt(&tile, board, row, col) {
					newBoard := board.copy()
					threads.Go(func() {
						placeNextTileOnBoard(tiles, &newBoard, solutions)
					})
					*board = origBoard.copy()
				}
			}
		}
		tile.rotateCW(1)
	}
	threads.Wait()
}

func placeNextTileOnBoard(tiles map[string]Tile, board *Board, solutions chan<- Board) {
	if board.isDead() {
		return
	}
	tile := getNextTile(tiles, board)
	origBoard := board.copy()
	for range tile.numFlips {
		for range tile.numRotations {
			for row, line := range board.cells {
				for col := range line {
					if placeTileOnBoardAt(&tile, board, row, col) {
						if isSolved(board) {
							solutions <- *board
						} else {
							placeNextTileOnBoard(tiles, board, solutions)
							*board = origBoard.copy()
						}
					}
				}
			}
			tile.rotateCW(1)
		}
		tile.flip()
	}
}

func placeTileOnBoardAt(tile *Tile, board *Board, row int, col int) bool {
	points := tile.getPoints()
	for _, point := range points {
		point.x += col
		point.y += row
		if !board.hasPoint(point) {
			return false
		}
		cell := board.cells[point.y][point.x]
		if !cell.isFree() {
			return false
		}
	}
	for _, point := range points {
		point.x += col
		point.y += row
		cell := &board.cells[point.y][point.x]
		cell.coveredBy = tile.name
	}
	return true
}

func getNextTile(tiles map[string]Tile, board *Board) Tile {
	for _, tile := range tiles {
		if !board.hasTile(tile) {
			return tile
		}
	}
	return tiles["b"]
}

func isSolved(board *Board) bool {
	for _, line := range board.cells {
		for _, cell := range line {
			if cell.isFree() {
				return false
			}
		}
	}
	return true
}
