package chess

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	color Color
}

func (p *Player) String() string {
	return fmt.Sprint(p.color)
}

type Game struct {
	players [2]*Player
	turn    int
	board   Board
}

func NewGame() *Game {
	white := &Player{White} //White goes first #TODO make it configurable
	black := &Player{Black}

	return &Game{players: [2]*Player{white, black}, turn: 0, board: *NewBoard()}
}

func (g *Game) Start() {
	rand.Seed(time.Now().UnixNano())
	for {

		activePlayer := g.players[g.turn%2]
		fmt.Println("Turn:", g.turn, "Player:", activePlayer)
		// Get all alive pieces for the player

		boxes := g.board.pieces[activePlayer.color]
		print(boxes)

		// Get all posible moves for all the pieces

		var moves []Move

		for _, box := range boxes {

			moves = append(moves, box.piece.GetMoves(g.board, *box.pos)...)
		}
		if len(moves) == 0 {
			// Either stale or check mate
			fmt.Println("State/checkmate")
			return
		}

		fmt.Printf("Moves possible %d \n", len(moves))
		choice := rand.Intn(len(moves))

		g.board.Move(moves[choice])
		g.board.Print()
		time.Sleep(time.Millisecond * 100)

		// Check endGame

		g.turn += 1

	}
}
