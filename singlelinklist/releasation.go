package singlelinklist

// New создание ограничителя
func New() *Cell {
	return &Cell{}
}

func (c *Cell) GetVal() string {
	return c.val
}

func (c *Cell) AddAtBeginning(val string) {
	newCell := &Cell{
		val:  val,
		next: c.next,
	}

	c.next = newCell
}

func (c *Cell) AddAtEnd(val string) {
	newCell := &Cell{
		val: val,
		next: nil,
	}

	list := c
	for list.next != nil {
		list = list.next
	}

	list.next = newCell
}

func (c *Cell) InsertCell(afterMeVal, val string)  error {
	findList := c.Find(afterMeVal)
	if findList == nil {
		return errNotFound
	}

	findList.AddAtBeginning(val)

	return nil
}

func (c *Cell) DeleteAfter() {
	c.next = c.next.next
}

func (c *Cell) Find(val string) *Cell{
	list := c
	for ; list.next != nil; list = list.next {
		if list.next.val == val {
			return list.next
		}
	}

	return nil
}

func(c *Cell) ClearAll() {
	c.next = nil
}

func (c *Cell) FindBefore(val string) *Cell{
	list := c
	for ; list.next != nil; list = list.next {
		if list.next.val == val {
			return list
		}
	}

	return nil
}

func (c *Cell) Iterate(function func (l *Cell) error ) error {
	list := c
	for ; list.next != nil; list = list.next {
		err := function(list.next)
		if err != nil {
			return err
		}
	}

	return nil
}

