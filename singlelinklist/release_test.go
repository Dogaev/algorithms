package singlelinklist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkList_Iterate(t *testing.T) {
	list := initTestMock()
	err := list.Iterate(func(l *Cell) error {
		fmt.Println(l.GetVal())
		return nil
	})

	assert.Nil(t, err)
}

func TestCell_DeleteAfter(t *testing.T) {
	list := initTestMock()

	findListBefore := list.FindBefore("lol")
	assert.NotNil(t, findListBefore)

	findListBefore.DeleteAfter()
	assert.Equal(t, "kel", findListBefore.next.val)
}

func TestCell_ClearAll(t *testing.T) {
	list := initTestMock()

	list.ClearAll()
	assert.Nil(t, list.next)
}

func TestLinkList_Find(t *testing.T) {
	list := initTestMock()

	findList := list.Find("lol")
	assert.Equal(t, "lol", findList.GetVal())

	findList = list.Find("1471824791")
	assert.Nil(t, findList)
}

func TestLinkList_InsertCell(t *testing.T) {
	list := initTestMock()

	findList := list.Find("prikol")
	assert.Nil(t, findList)

	err := list.InsertCell("kek", "prikol")
	assert.Nil(t, err)

	findList = list.Find("prikol")
	assert.NotNil(t, findList)
	assert.Equal(t, "lol", findList.next.val)
}

func initTestMock() *Cell{
	// Создание ограничителя
	list := New()
	// Заполнение списка
	list.AddAtEnd("kek")
	list.AddAtEnd("lol")
	list.AddAtEnd("kel")

	return list
}