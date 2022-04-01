package chess

import (
	"fmt"
	"strings"
)

type Move struct {
	source      Position
	destination Position
}

type Position struct {
	row int
	col int
}

func NewPosition(row, col int) (Position, error) {
	if row >= 0 && col >= 0 && row <= 7 && col <= 7 {
		return Position{row, col}, nil
	} else {
		return Position{}, fmt.Errorf("Invalid")
	}
}

type Box struct {
	piece     Piece
	pos       *Position
	attackers []*Box
}

func (b *Box) String() string {
	return fmt.Sprintf("%v @ %v", b.piece, b.pos)
}

func NewEmptyBox(r, c int) *Box {
	return &Box{
		piece:     nil,
		attackers: []*Box{},
		pos:       &Position{r, c},
	}
}

func NewBox(p Piece, r, c int) *Box {
	return &Box{
		piece:     p,
		attackers: []*Box{},
		pos:       &Position{r, c},
	}
}

type Board struct {
	boxes  [8][8]*Box
	pieces map[Color][]*Box
}

func valid(p Position) bool {
	return p.row >= 0 && p.col >= 0 && p.row <= 7 && p.col <= 7
}

func (b *Board) At(p Position) *Box {
	if valid(p) {
		return b.boxes[p.row][p.col]
	}
	return nil

}

func (b *Board) Remove(p Position) {
	end := b.At(p)
	if end.piece != nil {
		color := Black
		if end.piece.IsWhite() {
			color = White
		}
		i := 0
		for _, box := range b.pieces[color] {
			if box != end {
				b.pieces[color][i] = box
				i++
			}
		}
		b.pieces[color] = b.pieces[color][:i]
		end.piece = nil
	}
}

func (b *Board) Check(c Color) bool {

	otherColor := White
	if c == White {
		otherColor = Black
	}
	check := false
	for _, box := range b.pieces[otherColor] {
		for _, move := range box.piece.GetMoves(*b, *box.pos) {
			attacked := b.At(move.destination)
			if attacked.piece != nil {
				if _, ok := (attacked.piece).(*King); ok {
					fmt.Println(box.piece, " attacking ", attacked.piece)
					check = true
				}
			}
		}
	}

	return check
}

func (b *Board) Move(move Move) error {
	fmt.Println("Making a move", move)
	piece := b.At(move.source).piece
	b.Remove(move.source)
	capture := b.At(move.destination)
	if capture.piece != nil {
		fmt.Println("Capturing:", piece, capture.piece)
		if strings.ToLower(capture.piece.Short()) == "k" {
			return fmt.Errorf("Game over")

		}
		capture.piece.Kill()
	}
	b.Remove(move.destination)

	capture.piece = piece

	b.pieces[piece.Color()] = append(b.pieces[piece.Color()], capture)

	return nil
}

func (b *Board) updateAttackers(move Move) {

}

func (b *Board) AddPiece(pieceType PieceType, color Color, row, col int) {
	box := b.boxes[row][col]
	box.piece = MapToPiece(pieceType)(color)
	b.pieces[color] = append(b.pieces[color], box)
}

func (b *Board) Print() {
	buf := ""
	for i := 7; i >= 0; i-- {
		for j := 0; j < 8; j++ {
			if b.At(Position{i, j}).piece != nil {
				buf += fmt.Sprintf(" %s ", b.At(Position{i, j}).piece.Short())
			} else {
				buf += " . "
			}
		}
		buf += "\n"
	}
	fmt.Println(buf)
}

func NewBoard() *Board {
	// Initialize all white pieces at row 0 and 1
	var boxes [8][8]*Box
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			boxes[i][j] = NewEmptyBox(i, j)
		}
	}
	colorMap := make(map[Color][]*Box)

	colorMap[White] = []*Box{}
	colorMap[Black] = []*Box{}
	b := &Board{
		boxes:  boxes,
		pieces: colorMap,
	}

	b.AddPiece(RookP, White, 0, 0)
	b.AddPiece(RookP, White, 0, 7)

	b.AddPiece(KnightP, White, 0, 1)
	b.AddPiece(KnightP, White, 0, 6)

	b.AddPiece(BishopP, White, 0, 2)
	b.AddPiece(BishopP, White, 0, 5)

	b.AddPiece(QueenP, White, 0, 3)
	b.AddPiece(KingP, White, 0, 4)

	for i := 0; i < 8; i++ {
		b.AddPiece(PawnP, White, 1, i)
	}

	// init all black pieces

	b.AddPiece(RookP, Black, 7, 0)
	b.AddPiece(RookP, Black, 7, 7)

	b.AddPiece(KnightP, Black, 7, 1)
	b.AddPiece(KnightP, Black, 7, 6)

	b.AddPiece(BishopP, Black, 7, 2)
	b.AddPiece(BishopP, Black, 7, 5)

	b.AddPiece(QueenP, Black, 7, 3)
	b.AddPiece(KingP, Black, 7, 4)

	for i := 0; i < 8; i++ {
		b.AddPiece(PawnP, Black, 6, i)
	}

	return b
}
