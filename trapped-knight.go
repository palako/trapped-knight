package main

import (
	"errors"
	"fmt"
	"sort"
)

func sortMovesByValue(moves []square) {
	sort.Slice(moves, func(i, j int) bool {
		return moves[i].value < moves[j].value
	})
}

func lowestValue(moves []square) (square, error) {
	if len(moves) > 0 {
		lowest := moves[0]

		for _, m := range moves {
			if m.value < lowest.value {
				lowest = m
			}
		}
		return lowest, nil
	}
	return *new(square), errors.New("lowestValue: empty list of moves")
}

func createBoard(boardType string, size int) IBoard {
	v := make(map[string]bool)
	switch boardType {
	case "spiral":
		return &spiralBoard{Board{BoardSize: size, visited: v}}
	case "diagonal":
		return &diagonalBoard{Board{BoardSize: size, visited: v}}
	}
	panic("Unknown board type")
}

func main() {
	b := createBoard("spiral", 3000)
	//b.PrintBoard()
	//b.HTMLBoardToFile(b)
	x, y := int(0), int(0)
	for i := 1; ; i++ {
		b.visit(x, y)
		moves := (getKnightMovesFromPosition(b, x, y))
		lowest, err := findLowestNonVisited(b, moves)
		if err != nil {
			fmt.Printf("%d: {x:%v y:%v v:%v}\n", i, x, y, b.GetNumber(x, y))
			return
		}
		fmt.Printf("%d: %v\n", i, lowest)
		x = lowest.row
		y = lowest.col
	}
}
