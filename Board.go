package main

import (
	"bufio"
	"errors"
	"fmt"
	"html/template"
	"os"
)

type square struct {
	row   int
	col   int
	value int
}

//IBoard TODO DOC
type IBoard interface {
	GetNumber(int, int) int
}

//Board TODO DOC
type Board struct {
	concreteBoard IBoard
	BoardSize     int
	numbers       [][]int // diagonalBoard doesn't need to hold the numbers in an array because of the formula in getValue
	visited       [][]bool
}

func (b *Board) visit(row int, col int) {
	b.visited[row][col] = true
}

//GetNumber TODO DOC
func (b Board) GetNumber(row int, column int) int {
	return b.concreteBoard.GetNumber(row, column)
}

func (b *Board) getKnightMovesFromPosition(row int, col int) []square {
	moves := make([]square, 0, 8) //empty slice to hold up to 8 moves
	if row > 1 {
		if col > 0 {
			moves = append(moves, square{row - 2, col - 1, b.GetNumber(row-2, col-1)})
		}
		if col < b.BoardSize-1 {
			moves = append(moves, square{row - 2, col + 1, b.GetNumber(row-2, col+1)})
		}
	}
	if row > 0 {
		if col > 1 {
			moves = append(moves, square{row - 1, col - 2, b.GetNumber(row-1, col-2)})
		}
		if col < b.BoardSize-2 {
			moves = append(moves, square{row - 1, col + 2, b.GetNumber(row-1, col+2)})
		}
	}
	if row < b.BoardSize-1 {
		if col > 1 {
			moves = append(moves, square{row + 1, col - 2, b.GetNumber(row+1, col-2)})
		}
		if col < b.BoardSize-2 {
			moves = append(moves, square{row + 1, col + 2, b.GetNumber(row+1, col+2)})
		}
	}
	if row < b.BoardSize-2 {
		if col > 0 {
			moves = append(moves, square{row + 2, col - 1, b.GetNumber(row+2, col-1)})
		}
		if col < b.BoardSize-1 {
			moves = append(moves, square{row + 2, col + 1, b.GetNumber(row+2, col+1)})
		}
	}
	return moves
}

func (b *Board) findLowestNonVisitedFromSortedList(moves []square) (square, error) {
	if len(moves) == 0 {
		return *new(square), errors.New("findLowestNonVisistedFromSortedList: empty list of moves")
	}
	for _, m := range moves {
		if !b.visited[m.row][m.col] {
			return m, nil
		}
	}
	return *new(square), errors.New("findLowestNonVisistedFromSortedList: All available squares already visisted")
}

func (b *Board) findLowestNonVisited(moves []square) (square, error) {
	sortMovesByValue(moves)
	return b.findLowestNonVisitedFromSortedList(moves)
}

func (b *Board) drawBoard() {
	for i := int(0); i < b.BoardSize; i++ {
		for j := int(0); j < b.BoardSize; j++ {
			fmt.Printf("%v ", b.GetNumber(i, j))
		}
		fmt.Println()
	}
}

func (b *Board) htmlBoardToFile() {
	t, _ := template.New("diagonalBoardTemplate.html").Funcs(template.FuncMap{"N": N}).ParseFiles("palako/trapped-knight/diagonalBoardTemplate.html")
	file, _ := os.Create("palako/trapped-knight/diagonalBoard.html")
	defer file.Close()
	bufferedWriter := bufio.NewWriter(file)
	t.Execute(bufferedWriter, b)
	bufferedWriter.Flush()
}
