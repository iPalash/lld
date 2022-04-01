package chess

import (
	"fmt"
	"strings"
)

type Piece interface {
	IsWhite() bool
	Color() Color
	IsAlive() bool
	Kill()
	GetMoves(Board, Position) []Move
	Move(Board, Move)
	Short() string
}

type PieceType int

const (
	KingP PieceType = iota
	QueenP
	BishopP
	KnightP
	RookP
	PawnP
)

func MapToPiece(p PieceType) func(Color) Piece {
	switch p {
	case KingP:
		return NewKing
	case QueenP:
		return NewQueen
	case BishopP:
		return NewBishop
	case KnightP:
		return NewKnight
	case RookP:
		return NewRook
	case PawnP:
		return NewPawn

	}
	return NewPawn
}

type Color int

const (
	White Color = iota
	Black
)

func (c Color) String() string {
	switch c {
	case 0:
		return "White"
	case 1:
		return "Black"
	}
	return ""
}

type PieceCommon struct {
	color   Color
	isAlive bool
	sym     string
}

func NewPiece(c Color, sym string) *PieceCommon {
	if c == Black {
		sym = strings.ToLower(sym) // Symbol will be lower in case of black
	}
	return &PieceCommon{
		color:   c,
		isAlive: true,
		sym:     sym,
	}
}

func (p *PieceCommon) String() string {
	return fmt.Sprintf("Color:%s, Alive:%t", p.color, p.isAlive)
}

func (p *PieceCommon) IsWhite() bool {
	return p.color == White
}

func (p *PieceCommon) IsAlive() bool {
	return p.isAlive
}

func (p *PieceCommon) Kill() {
	p.isAlive = false
}

func (p *PieceCommon) Short() string {
	return p.sym
}

// King, Queen, Bishop, Knight, Rook

type King struct {
	*PieceCommon
	castled bool
	sym     string
}

func NewKing(c Color) Piece {
	k := &King{
		PieceCommon: NewPiece(c, `K`),
		castled:     false,
	}
	return k
}

func (k *King) String() string {
	return fmt.Sprintf("%s=>%v : castled:%t", "King", k.PieceCommon, k.castled)
}

func (k *King) Color() Color {
	return k.color
}

func (k *King) GetMoves(board Board, p Position) []Move {
	return []Move{}
}

func (k *King) Move(board Board, move Move) {

}

type Queen struct {
	*PieceCommon
}

func NewQueen(c Color) Piece {
	return &Queen{
		PieceCommon: NewPiece(c, "Q"),
	}
}

func (q *Queen) String() string {
	return fmt.Sprintf("%s=>%v", "Queen", q.PieceCommon)
}

func (k *Queen) Color() Color {
	return k.color
}

func (q *Queen) GetMoves(board Board, p Position) []Move {
	return []Move{}
}

func (q *Queen) Move(board Board, move Move) {

}

type Bishop struct {
	*PieceCommon
	castled bool
}

func NewBishop(c Color) Piece {
	return &Bishop{
		PieceCommon: NewPiece(c, "B"),
	}
}

func (k *Bishop) String() string {
	return fmt.Sprintf("%s=>%v", "Bishop", k.PieceCommon)
}

func (k *Bishop) Color() Color {
	return k.color
}

func (k *Bishop) GetMoves(board Board, p Position) []Move {
	return []Move{}
}

func (k *Bishop) Move(board Board, move Move) {

}

type Knight struct {
	*PieceCommon
}

func NewKnight(c Color) Piece {
	return &Knight{
		PieceCommon: NewPiece(c, "N"),
	}
}

func (k *Knight) String() string {
	return fmt.Sprintf("%s=>%v", "Knight", k.PieceCommon)
}

func (k *Knight) Color() Color {
	return k.color
}

func (k *Knight) GetMoves(board Board, p Position) []Move {
	return []Move{}
}

func (k *Knight) Move(board Board, move Move) {

}

type Rook struct {
	*PieceCommon
	castled bool
}

func NewRook(c Color) Piece {
	return &Rook{
		PieceCommon: NewPiece(c, "R"),
		castled:     false,
	}
}

func (k *Rook) String() string {
	return fmt.Sprintf("%s=>%v : castled:%t", "Rook", k.PieceCommon, k.castled)
}

func (k *Rook) Color() Color {
	return k.color
}

func (k *Rook) GetMoves(board Board, p Position) []Move {
	return []Move{}
}

func (k *Rook) Move(board Board, move Move) {

}

type Pawn struct {
	*PieceCommon
	castled bool
}

func NewPawn(c Color) Piece {
	return &Pawn{
		PieceCommon: NewPiece(c, "P"),
	}
}

func (k *Pawn) String() string {
	return fmt.Sprintf("%s=>%v", "Pawn", k.PieceCommon)
}

func (k *Pawn) Color() Color {
	return k.color
}

func (k *Pawn) GetMoves(board Board, p Position) []Move {

	var moves []Move

	// If white, move forward
	direction := 1
	if k.color == Black {
		direction = -1
	}

	// TODO: Handle 2 jumps at first pos later

	nxt, err := NewPosition(p.row+direction, p.col)
	if err == nil && board.At(nxt).piece == nil {
		moves = append(moves, Move{p, nxt})
	}
	firstCross, err := NewPosition(p.row+direction, p.col+1)

	if err == nil && board.At(firstCross).piece != nil && board.At(firstCross).piece.IsWhite() != k.IsWhite() {
		moves = append(moves, Move{p, firstCross})
	}

	secondCross, err := NewPosition(p.row+direction, p.col-1)

	if err == nil && board.At(secondCross).piece != nil && board.At(secondCross).piece.IsWhite() != k.IsWhite() {
		moves = append(moves, Move{p, secondCross})
	}

	// If diagonal has other color pawn move diagonal
	return moves
}

func (k *Pawn) Move(board Board, move Move) {

}
