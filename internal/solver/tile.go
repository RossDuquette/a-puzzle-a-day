package solver

import (
	"fmt"
)

type TileName string

const vacant_tile TileName = ""

type Tile struct {
	name TileName
	shape string
	rotations uint
	flippable bool
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

func get_tiles() map[TileName]Tile {
	tiles := map[TileName]Tile {
		"s": { name: "s", shape: "ldll",  rotations: 3, flippable: true },
		"y": { name: "y", shape: "rrudr", rotations: 3, flippable: true },
		"z": { name: "z", shape: "rddr",  rotations: 1, flippable: true },
		"u": { name: "u", shape: "drru",  rotations: 3, flippable: false },
		"p": { name: "p", shape: "rdld",  rotations: 3, flippable: true },
		"l": { name: "l", shape: "dddr",  rotations: 3, flippable: true },
		"v": { name: "v", shape: "ddrr",  rotations: 3, flippable: false },
		"b": { name: "b", shape: "ddruu", rotations: 1, flippable: false },
	}
	return tiles
}
