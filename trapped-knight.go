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

//N allows integer ennumeration from templates
func N(start, end int) (stream chan int) {
	stream = make(chan int)
	go func() {
		for i := start; i <= end; i++ {
			stream <- i
		}
		defer close(stream)
	}()
	return stream
}

func createBoard(boardType string, size int) Board {
	v := make([][]bool, size)
	for i := range v {
		v[i] = make([]bool, size)
	}
	switch boardType {
	case "diagonal":
		return Board{concreteBoard: diagonalBoard{}, BoardSize: size, visited: v}
	}
	panic("Unknown board type")
}

func main() {
	b := createBoard("diagonal", 100)
	//b.drawBoard()
	b.htmlBoardToFile()
	x, y := int(0), int(0)
	for {
		b.visit(x, y)
		moves := (b.getKnightMovesFromPosition(x, y))
		lowest, err := b.findLowestNonVisited(moves)
		if err != nil {
			return
		}
		fmt.Println(lowest)
		x = lowest.row
		y = lowest.col
	}
}
