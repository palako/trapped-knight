package main

import (
	"fmt"
)

type spiralBoard struct {
	Board
}

func (b *spiralBoard) LeftEdge() int {
	return -b.GetBoardSize() / 2
}
func (b *spiralBoard) RightEdge() int {
	return (b.GetBoardSize() / 2) - 1
}
func (b *spiralBoard) TopEdge() int {
	return -b.GetBoardSize() / 2
}
func (b *spiralBoard) BottomEdge() int {
	return (b.GetBoardSize() / 2) - 1
}

// getNumber Check docs/spiral.xlsx for an explanation on the following procedure to calculate
// the value of the corresponding square in the spiral in O(1)
func (b *spiralBoard) GetNumber(row int, column int) int {

	//flip the y axes. (0,0) will be in the middle of the board, with positive numbers going up
	row = -row

	// r is the outer ring of the spiral
	r := max(abs(row), abs(column))

	//accumInnerRings is the count of elements from all of the inner spirals
	accumInnerRings := (4 * r * r) - (4 * r) + 1

	totalOuterRing := 8 * r
	edgeSide := totalOuterRing / 4
	//bring (0,0) to the bottom right of the outer ring
	shiftColumn := abs(column - (edgeSide / 2))
	shiftRow := row + (edgeSide / 2)

	var previousEdges int
	var edge int
	switch {
	case column == r && row != -r:
		//right edge of the outer spiral, except for the bottom-right corner
		previousEdges = 0 * edgeSide
		edge = previousEdges + shiftRow

	case row == r && column != r:
		//top edge of the outer spiral
		previousEdges = 1 * edgeSide
		edge = previousEdges + shiftColumn

	case column == -r && row != r:
		//left edge of the outer spiral
		previousEdges = 2 * edgeSide
		edge = previousEdges + (edgeSide - shiftRow)

	case row == -r && column != -r:
		//bottom edge of the outer spiral
		previousEdges = 3 * edgeSide
		edge = previousEdges + (edgeSide - shiftColumn)
	}
	return accumInnerRings + edge
}

func (b *spiralBoard) PrintBoard() {
	half := b.GetBoardSize() / 2
	for i := -half; i < half; i++ {
		for j := -half; j < half; j++ {
			fmt.Printf("%v ", b.GetNumber(i, j))
		}
		fmt.Println()
	}
}
