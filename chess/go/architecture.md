# Chess Design in Go

Entities are:

1. Game
2. Player
3. Board
4. Piece
5. Move

## 1. Game

- Controls the flow of the game, including turn management

## 2. Player

- Represents player involved in the game

## 3. Board

- Represents the state of the board
- Evaluates end games, blocking positions, feasibility moves, etc

## 4. Piece

- An interface which can be extended for every piece of the board
- Given a position and board, each piece itself can define what moves are possible
- We can put this functionality in board as well since a move is never isolated and always impacts the entire board
- What we did here is a piece simply tells what moves are legally possible  for the piece and the board evaluates which of them can be played considering check scenarios by mock-playing the move & checking the state of the board

## 5. Move

- Represents a legal move from one pos of the board to another
