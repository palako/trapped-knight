package main

import (
	"errors"
	"flag"
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

func main() {

	boardType := flag.String("type", "spiral", "Supported types are diagonal and spiral")
	boardSize := flag.Int("size", 3000, "Size of one side of the board, creates a board of size x size squares")
	printFlag := flag.String("print", "", "Supported values are console and html")

	flag.Parse()

	if *boardType != "spiral" && *boardType != "diagonal" {
		fmt.Println("ERROR: Unsupported board type. Supported types are \"spiral\" and \"diagonal\"")
		return
	}

	b := createBoard(*boardType, *boardSize)
	switch *printFlag {
	case "console":
		b.PrintBoard()
	case "html":
		b.HTMLBoardToFile(b)
	}

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
