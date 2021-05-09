package doublylinkedlist

type Cell struct {
	prev, next *Cell
	val string
}

func New() *Cell {
	return &Cell{}
}

