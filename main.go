package main

import (
	"fmt"
	"github.com/Dogaev/algorithms/singlelinklist"
)

func main() {
	list := singlelinklist.New()
	list.AddAtEnd("123")

	err := list.Iterate(func(l *singlelinklist.Cell) error {
		fmt.Println(l.GetVal())
		return nil
	})
	if err != nil {
		panic(err)
	}
}
