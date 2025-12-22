package solver

type tile_id string

const tile_vacant tile_id = ""

func get_tile_ids() [9]tile_id {
	tiles := [...]tile_id {
		tile_vacant,
		"s",
		"y",
		"z",
		"u",
		"p",
		"l",
		"v",
		"b",
	}
	return tiles
}
