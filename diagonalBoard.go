package main

type diagonalBoard struct {
	b Board
}

func (b diagonalBoard) GetNumber(row int, column int) int {
	n := row + column
	return row + ((n*n + n + 2) / 2)
}
