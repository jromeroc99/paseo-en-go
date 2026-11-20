//go:build !solution

// Package player keeps the score of a game player.
package player

// Player is someone taking part in a game.
type Player struct {
	Name  string
	Score int
}

// NewPlayer returns a pointer to a new player with the given name and no points.
// Example: NewPlayer("Ana") returns &Player{Name: "Ana", Score: 0}.
func NewPlayer(name string) *Player {
	// TODO: return a pointer to a Player literal.
	return nil
}

// AddPoints adds points to the player's score.
// Example: if p.Score is 10, after AddPoints(p, 5) it is 15.
func AddPoints(p *Player, points int) {
	// TODO: change the Score field through the pointer.
}
