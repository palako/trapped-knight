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
	GetBoardSize() int
	IsVisited(int, int) bool
	visit(int, int)
	PrintBoard()
	HTMLBoardToFile(IBoard)
	LeftEdge() int
	RightEdge() int
	TopEdge() int
	BottomEdge() int
}

//Board TODO DOC
type Board struct {
	BoardSize int
	visited   map[string]bool
}

func (b *Board) visit(row int, col int) {
	key := fmt.Sprintf("%d,%d", row, col)
	b.visited[key] = true
}

//GetBoardSize TODO DOC
func (b *Board) GetBoardSize() int {
	return b.BoardSize
}

//IsVisited returns if the square at (row, col) has already been visited
func (b *Board) IsVisited(row int, col int) bool {
	key := fmt.Sprintf("%d,%d", row, col)
	return b.visited[key]
}

func getKnightMovesFromPosition(b IBoard, row int, col int) []square {
	moves := make([]square, 0, 8) //empty slice to hold up to 8 moves
	if row > b.TopEdge()+1 {
		if col > b.LeftEdge() {
			moves = append(moves, square{row - 2, col - 1, b.GetNumber(row-2, col-1)})
		}
		if col < b.RightEdge() {
			moves = append(moves, square{row - 2, col + 1, b.GetNumber(row-2, col+1)})
		}
	}
	if row > b.TopEdge() {
		if col > b.LeftEdge()+1 {
			moves = append(moves, square{row - 1, col - 2, b.GetNumber(row-1, col-2)})
		}
		if col < b.RightEdge()-1 {
			moves = append(moves, square{row - 1, col + 2, b.GetNumber(row-1, col+2)})
		}
	}
	if row < b.BottomEdge() {
		if col > b.LeftEdge()+1 {
			moves = append(moves, square{row + 1, col - 2, b.GetNumber(row+1, col-2)})
		}
		if col < b.RightEdge()-1 {
			moves = append(moves, square{row + 1, col + 2, b.GetNumber(row+1, col+2)})
		}
	}
	if row < b.BottomEdge()-1 {
		if col > b.LeftEdge() {
			moves = append(moves, square{row + 2, col - 1, b.GetNumber(row+2, col-1)})
		}
		if col < b.RightEdge() {
			moves = append(moves, square{row + 2, col + 1, b.GetNumber(row+2, col+1)})
		}
	}
	return moves
}

func findLowestNonVisitedFromSortedList(b IBoard, moves []square) (square, error) {
	if len(moves) == 0 {
		return *new(square), errors.New("findLowestNonVisistedFromSortedList: empty list of moves")
	}
	for _, m := range moves {
		if !b.IsVisited(m.row, m.col) {
			return m, nil
		}
	}
	return *new(square), errors.New("findLowestNonVisistedFromSortedList: All available squares already visisted")
}

func findLowestNonVisited(b IBoard, moves []square) (square, error) {
	sortMovesByValue(moves)
	return findLowestNonVisitedFromSortedList(b, moves)
}

//N allows integer ennumeration from templates
func N(b IBoard) (stream chan int) {
	stream = make(chan int)
	go func() {
		for i := b.LeftEdge(); i <= b.RightEdge(); i++ {
			stream <- i
		}
		defer close(stream)
	}()
	return stream
}

//HTMLBoardToFile TODO DOC
func (*Board) HTMLBoardToFile(b IBoard) {
	t, _ := template.New("boardTemplate.html").Funcs(template.FuncMap{"N": N}).ParseFiles("palako/trapped-knight/boardTemplate.html")
	file, _ := os.Create("palako/trapped-knight/board.html")
	defer file.Close()
	bufferedWriter := bufio.NewWriter(file)
	t.Execute(bufferedWriter, b)
	bufferedWriter.Flush()
}
