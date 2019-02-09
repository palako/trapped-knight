package main

import "fmt"

type diagonalBoard struct {
	Board
}

func (b diagonalBoard) LeftEdge() int {
	return 0
}
func (b diagonalBoard) RightEdge() int {
	return b.GetBoardSize() - 1
}
func (b diagonalBoard) TopEdge() int {
	return 0
}
func (b diagonalBoard) BottomEdge() int {
	return b.GetBoardSize() - 1
}

func (b diagonalBoard) GetNumber(row int, column int) int {
	n := row + column
	return row + ((n*n + n + 2) / 2)
}

func (b *diagonalBoard) PrintBoard() {
	for i := int(0); i < b.GetBoardSize(); i++ {
		for j := int(0); j < b.GetBoardSize(); j++ {
			fmt.Printf("%v ", b.GetNumber(i, j))
		}
		fmt.Println()
	}
}
