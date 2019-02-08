package main

import (
	"reflect"
	"testing"
)

func Test_diagonalBoard_getKnightMovesFromPosition(t *testing.T) {
	type fields struct {
		boardSize int
		numbers   [][]int
		visited   [][]bool
	}
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []square
	}{
		{
			name:   "leftTopCorner",
			fields: fields{boardSize: 5},
			args:   args{0, 0},
			want:   []square{{1, 2, 8}, {2, 1, 9}},
		},
		{
			name:   "rightTop",
			fields: fields{boardSize: 5},
			args:   args{0, 4},
			want:   []square{{1, 2, 8}, {2, 3, 18}},
		},
		{
			name:   "leftBottom",
			fields: fields{boardSize: 5},
			args:   args{4, 0},
			want:   []square{{2, 1, 9}, {3, 2, 19}},
		},
		{
			name:   "rightBottom",
			fields: fields{boardSize: 5},
			args:   args{4, 4},
			want:   []square{{2, 3, 18}, {3, 2, 19}},
		},
		{
			name:   "center",
			fields: fields{boardSize: 5},
			args:   args{2, 2},
			want:   []square{{0, 1, 2}, {0, 3, 7}, {1, 0, 3}, {1, 4, 17}, {3, 0, 10}, {3, 4, 32}, {4, 1, 20}, {4, 3, 33}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := createBoard("diagonal", tt.fields.boardSize)
			//b := &diagonalBoard{
			//	BoardSize: tt.fields.boardSize,
			//	numbers:   tt.fields.numbers,
			//	visited:   tt.fields.visited,
			//}
			if got := b.getKnightMovesFromPosition(tt.args.row, tt.args.col); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diagonalBoard.getKnightMovesFromPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
