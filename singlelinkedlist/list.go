package singlelinkedlist

import "errors"

type List struct {
	top *Cell
	len int
}

// CreateList создание односвязного списка
func CreateList() *List {
	return &List{
		top: &Cell{
			value: 0,
			next: nil,
		},
	}
}

// Iterate Передвижение по связному списку. Сложность O(N)
func (l *List) Iterate(function func (l *Cell) error) {
	top := l.top
	for ; top.next != nil; top = top.next {
		function(top.next)
	}
}

// FindCell Поиск ячейки в списке. Сложность O(N)
func (l *List) FindCell(target int) *Cell {
	// пропускаем ячейку ограничитель
	top := l.top
	if top.next != nil {
		top = top.next
	}

	for ; top != nil; top = top.next {
		if top.value == target {
			return top
		}
	}

	return nil
}

// FindPrevCell Поиск предыдущей ячейки в списке. Сложность O(N)
func (l *List) FindPrevCell(target int) *Cell {
	top := l.top
	for ; top.next != nil; top = top.next {
		if top.next.value == target {
			return top
		}
	}

	return nil
}

// AddAtBeginning Добавление ячейки в начало списка. Сложность O(1)
func (l *List) AddAtBeginning(newCell *Cell) {
	newCell.next = l.top.next
	l.top.next = newCell
	l.len++
}

// AddAtEnd Добавление ячейки в конец списка. Сложность O(N)
func (l *List) AddAtEnd(newCell *Cell) {
	top := l.top
	for ; top.next != nil; {
		top = top.next
	}

	top.next = newCell
	newCell.next = nil
	l.len++
}

func (l *List) GetLength() (cnt int) {
	return l.len
}

func (l *List) DeleteByValue(target int) error {
	fc := l.FindPrevCell(target)
	if fc == nil {
		return errors.New("not found any cell")
	}

	fc.DeleteNext()
	l.len--

	return nil
}

func (l *List) AddAfterByVal(target int, cell *Cell) error {
	fc := l.FindCell(target)
	if fc == nil {
		return errors.New("not found any cell")
	}

	fc.AddAfter(cell)
	l.len++

	return nil
}