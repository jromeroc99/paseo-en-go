//go:build solution

// Package player keeps the score of a game player.
package player

// Player is someone taking part in a game.
type Player struct {
	Name  string
	Score int
}

// NewPlayer returns a pointer to a new player with the given name and no points.
func NewPlayer(name string) *Player {
	return &Player{Name: name}
}

// AddPoints adds points to the player's score.
func AddPoints(p *Player, points int) {
	p.Score += points
}
