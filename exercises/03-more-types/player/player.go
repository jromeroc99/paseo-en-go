//go:build !solution

// Package player keeps the score of a game player.
//
// Write this type and these functions in this file:
//
//	type Player struct {
//		Name  string
//		Score int
//	}
//
//	func NewPlayer(name string) *Player
//	func AddPoints(p *Player, points int)
//
// Player is someone taking part in a game.
//
// NewPlayer returns a pointer to a new player with the given name and no
// points.
// Example: NewPlayer("Ana") returns &Player{Name: "Ana", Score: 0}.
//
// AddPoints adds points to the player's score.
// Example: if p.Score is 10, after AddPoints(p, 5) it is 15.
//
// Hint: NewPlayer returns a pointer to a Player literal. AddPoints changes
// the Score field through the pointer.
package player
