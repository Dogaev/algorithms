package doublylinkedlist

func (c *Cell) InsertCell(afterMeVal, val string)  error {
	findList := c.Find(afterMeVal)
	if findList == nil {
		return errNotFound
	}

	findList.AddAtBeginning(val)

	return nil
}

func (c *Cell) Find()
