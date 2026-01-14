package solver

const vacantCell string = "vacant"

type Cell struct {
	name      string
	coveredBy string
}

func (c Cell) showing() string {
	if c.coveredBy == vacantCell {
		return c.name
	}
	return c.coveredBy
}

func (c Cell) isFree() bool {
	return c.name != "" && c.coveredBy == vacantCell
}

func (c *Cell) setVacant() {
	c.coveredBy = vacantCell
}
