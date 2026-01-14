package solver

const vacant_cell string = "vacant"

type Cell struct {
	name       string
	covered_by string
}

func (c Cell) showing() string {
	if c.covered_by == vacant_cell {
		return c.name
	}
	return c.covered_by
}

func (c Cell) is_free() bool {
	return c.name != "" && c.covered_by == vacant_cell
}

func (c *Cell) set_vacant() {
	c.covered_by = vacant_cell
}
