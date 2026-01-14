package solver

import (
	"fmt"
)

type Tile struct {
	name         string
	shape        string
	numRotations uint
	numFlips     uint
}

func getTiles() map[string]Tile {
	// Shapes are rotated/flipped such that they can be placed in the top-left
	// available square.
	tiles := map[string]Tile{
		"s": {name: "s", shape: "rdrr", numRotations: 4, numFlips: 2},
		"y": {name: "y", shape: "rrdur", numRotations: 4, numFlips: 2},
		"z": {name: "z", shape: "rddr", numRotations: 2, numFlips: 2},
		"u": {name: "u", shape: "drru", numRotations: 4, numFlips: 1},
		"p": {name: "p", shape: "rdld", numRotations: 4, numFlips: 2},
		"l": {name: "l", shape: "dddr", numRotations: 4, numFlips: 2},
		"v": {name: "v", shape: "ddrr", numRotations: 4, numFlips: 1},
		"b": {name: "b", shape: "ddruu", numRotations: 2, numFlips: 1},
	}
	return tiles
}

func (t *Tile) rotateCW(rotations uint) {
	if rotations >= t.numRotations {
		msg := fmt.Sprintf("Cannot rotate %s %d times", t.name, rotations)
		panic(msg)
	}
	newShape := ""
	for _, c := range t.shape {
		var directions string
		switch c {
		case 'r':
			directions = "rdlu"
		case 'd':
			directions = "dlur"
		case 'l':
			directions = "lurd"
		case 'u':
			directions = "urdl"
		default:
			panic("Invalid shape direction")
		}
		newShape += string(directions[rotations])
	}
	t.shape = newShape
}

func (t *Tile) flip() {
	newShape := ""
	for _, c := range t.shape {
		switch c {
		case 'r':
			newShape += "l"
		case 'd':
			newShape += "d"
		case 'l':
			newShape += "r"
		case 'u':
			newShape += "u"
		default:
			panic("Invalid shape direction")
		}
	}
	t.shape = newShape
}

type Point struct {
	x int
	y int
}

func (t *Tile) getPoints() []Point {
	curPoint := Point{0, 0}
	points := []Point{curPoint}
	for _, c := range t.shape {
		switch c {
		case 'r':
			curPoint.x += 1
		case 'l':
			curPoint.x -= 1
		case 'd':
			curPoint.y += 1
		case 'u':
			curPoint.y -= 1
		default:
			panic("Invalid shape direction")
		}
		points = append(points, curPoint)
	}
	points = removeDuplicates(points)
	points = shiftOriginTopmostLeftmost(points)
	return points
}

func removeDuplicates(points []Point) []Point {
	found := make(map[Point]bool)
	newPoints := []Point{}

	for _, point := range points {
		_, exists := found[point]
		if !exists {
			found[point] = true
			newPoints = append(newPoints, point)
		}
	}
	return newPoints
}

func shiftOriginTopmostLeftmost(points []Point) []Point {
	tlPoint := Point{0, 0}
	for _, point := range points {
		if point.y < tlPoint.y || (point.y == tlPoint.y && point.x < tlPoint.x) {
			tlPoint = point
		}
	}
	// Adjust every point's offset
	for i := range points {
		points[i].x -= tlPoint.x
		points[i].y -= tlPoint.y
	}
	return points
}
