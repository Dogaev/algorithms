package singlelinkedlist

type Cell struct {
	value int
	next *Cell
}

func CreateCell(val int) *Cell {
	return &Cell{
		value: val,
		next: nil,
	}
}

func (c *Cell) GetValue() int {
	return c.value
}

func (c *Cell) AddAfter(cell *Cell) {
	cell.next = c.next
	c.next = cell
}

func (c *Cell) DeleteNext() {
	c.next = c.next.next
}