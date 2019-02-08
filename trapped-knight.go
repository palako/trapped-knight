package main

import (
	"bufio"
	"errors"
	"fmt"
	"html/template"
	"os"
	"sort"
)

type square struct {
	row   int
	col   int
	value int
}

type diagonalBoard struct {
	BoardSize int
	numbers   [][]int // diagonalBoard doesn't need to hold the numbers in an array because of the formula in getValue
	visited   [][]bool
}

func (b *diagonalBoard) visit(row int, col int) {
	b.visited[row][col] = true
}

func (b *diagonalBoard) GetNumber(row int, column int) int {
	n := row + column
	return row + ((n*n + n + 2) / 2)
}

func (b *diagonalBoard) getKnightMovesFromPosition(row int, col int) []square {
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

func (b *diagonalBoard) findLowestNonVisitedFromSortedList(moves []square) (square, error) {
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

func (b *diagonalBoard) findLowestNonVisited(moves []square) (square, error) {
	sortMovesByValue(moves)
	return b.findLowestNonVisitedFromSortedList(moves)
}

func (b *diagonalBoard) drawBoard() {
	for i := int(0); i < b.BoardSize; i++ {
		for j := int(0); j < b.BoardSize; j++ {
			fmt.Printf("%v ", b.GetNumber(i, j))
		}
		fmt.Println()
	}
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

func (b *diagonalBoard) htmlBoardToFile() {
	t, _ := template.New("diagonalBoardTemplate.html").Funcs(template.FuncMap{"N": N}).ParseFiles("palako/trapped-knight/diagonalBoardTemplate.html")
	file, _ := os.Create("palako/trapped-knight/diagonalBoard.html")
	defer file.Close()
	bufferedWriter := bufio.NewWriter(file)
	t.Execute(bufferedWriter, b)
	bufferedWriter.Flush()
}

func createBoard(size int) diagonalBoard {
	v := make([][]bool, size)
	for i := range v {
		v[i] = make([]bool, size)
	}
	return diagonalBoard{BoardSize: size, visited: v}
}

func main() {
	b := createBoard(100)
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
