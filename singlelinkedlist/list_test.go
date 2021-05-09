package singlelinkedlist_test

import (
	"fmt"
	"github.com/Dogaev/algorithms/singlelinkedlist"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_AddAtBeginning(t *testing.T) {
	list := singlelinkedlist.CreateList()
	list.AddAtBeginning(singlelinkedlist.CreateCell(1))
	assert.Equal(t, 1, list.FindCell(1).GetValue())
	assert.Equal(t, 1, list.GetLength())
	list.AddAtBeginning(singlelinkedlist.CreateCell(2))
	assert.Equal(t, 2, list.FindCell(2).GetValue())
	assert.Equal(t, 2, list.GetLength())

	list.Iterate(printCell)
}

func TestList_AddAtEnd(t *testing.T) {
	list := singlelinkedlist.CreateList()
	list.AddAtEnd(singlelinkedlist.CreateCell(1))
	assert.Equal(t, 1, list.FindCell(1).GetValue())
	assert.Equal(t, 1, list.GetLength())
	list.AddAtEnd(singlelinkedlist.CreateCell(2))
	assert.Equal(t, 2, list.FindCell(2).GetValue())
	assert.Equal(t, 2, list.GetLength())

	list.Iterate(printCell)
}

func TestList_DeleteByValue(t *testing.T) {
	list := singlelinkedlist.CreateList()
	list.AddAtEnd(singlelinkedlist.CreateCell(1))
	list.AddAtEnd(singlelinkedlist.CreateCell(2))
	list.AddAtEnd(singlelinkedlist.CreateCell(3))
	assert.Equal(t, 3, list.FindCell(3).GetValue())
	assert.Equal(t, 3, list.GetLength())

	err := list.DeleteByValue(5)
	assert.NotNil(t, err)

	err = list.DeleteByValue(1)
	assert.Nil(t, err)
	assert.Nil(t, list.FindCell(1))
	assert.Equal(t, 2, list.GetLength())
}

func TestList_AddAfterByVal(t *testing.T) {
	list := singlelinkedlist.CreateList()
	list.AddAtEnd(singlelinkedlist.CreateCell(1))
	list.AddAtEnd(singlelinkedlist.CreateCell(2))
	list.AddAtEnd(singlelinkedlist.CreateCell(4))
	assert.Equal(t, 4, list.FindCell(4).GetValue())
	assert.Equal(t, 3, list.GetLength())

	err := list.AddAfterByVal(3, singlelinkedlist.CreateCell(3))
	assert.NotNil(t, err)

	err = list.AddAfterByVal(2, singlelinkedlist.CreateCell(3))
	assert.Nil(t, err)
	assert.Equal(t, 3, list.FindCell(3).GetValue())
	assert.Equal(t, 4, list.GetLength())
}

func printCell(c *singlelinkedlist.Cell) error {
	fmt.Println(c.GetValue())
	return nil
}