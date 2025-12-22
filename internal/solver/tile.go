package solver

import (
	"fmt"
)

type Tile struct {
	name string
	shape string
	rotations uint
	flippable bool
}

func get_tiles() map[string]Tile {
	// Shapes are rotated/flipped such that they can be placed in the top-left
	// available square.
	tiles := map[string]Tile {
		"s": { name: "s", shape: "rdrr",  rotations: 3, flippable: true },
		"y": { name: "y", shape: "rrdur", rotations: 3, flippable: true },
		"z": { name: "z", shape: "rddr",  rotations: 1, flippable: true },
		"u": { name: "u", shape: "drru",  rotations: 3, flippable: false },
		"p": { name: "p", shape: "rdld",  rotations: 3, flippable: true },
		"l": { name: "l", shape: "dddr",  rotations: 3, flippable: true },
		"v": { name: "v", shape: "ddrr",  rotations: 3, flippable: false },
		"b": { name: "b", shape: "ddruu", rotations: 1, flippable: false },
	}
	return tiles
}

func (t *Tile) rotate_cw(rotations uint) {
	if rotations > t.rotations {
		msg := fmt.Sprintf("Cannot rotate %s %d times", t.name, rotations)
		panic(msg)
	}
	new_shape := ""
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
		new_shape += string(directions[rotations])
	}
	t.shape = new_shape
}

func (t *Tile) flip() {
	new_shape := ""
	for _, c := range t.shape {
		switch c {
		case 'r':
			new_shape += "l"
		case 'd':
			new_shape += "d"
		case 'l':
			new_shape += "r"
		case 'u':
			new_shape += "u"
		default:
			panic("Invalid shape direction")
		}
	}
	t.shape = new_shape
}

type Point struct {
	x int
	y int
}

func (t *Tile) get_points() []Point {
	cur_point := Point{0, 0}
	points := []Point{cur_point}
	for _, c := range t.shape {
		switch c {
		case 'r':
			cur_point.x += 1
		case 'l':
			cur_point.x -= 1
		case 'd':
			cur_point.y += 1
		case 'u':
			cur_point.y -= 1
		default:
			panic("Invalid shape direction")
		}
		points = append(points, cur_point)
	}
	points = remove_duplicates(points)
	points = shift_origin_topmost_leftmost(points)
	return points
}

func remove_duplicates(points []Point) []Point {
	found := make(map[Point]bool)
    new_points := []Point{}

    for _, point := range points {
        _, exists := found[point]
        if !exists {
            found[point] = true
            new_points = append(new_points, point)
        }
    }
    return new_points
}

func shift_origin_topmost_leftmost(points []Point) []Point {
	tl_point := Point{0, 0}
    for _, point := range points {
		if point.y < tl_point.y || (point.y == tl_point.y && point.x < tl_point.x) {
			tl_point = point
		}
	}
	// Adjust every point's offset
    for i := range points {
		points[i].x -= tl_point.x
		points[i].y -= tl_point.y
	}
	return points
}
