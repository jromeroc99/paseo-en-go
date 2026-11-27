//go:build solution

// Package robot tells an empty interface from one that holds a nil pointer.
package robot

// Speaker is anything that can say something.
type Speaker interface {
	Speak() string
}

// Robot is a machine with a name.
type Robot struct {
	Name string
}

// Speak introduces the robot, or snores when the robot is a nil pointer.
func (r *Robot) Speak() string {
	if r == nil {
		return "zzz"
	}
	return "I am " + r.Name
}

// Say returns what s says, or "nobody is there" when s is a nil interface.
func Say(s Speaker) string {
	if s == nil {
		return "nobody is there"
	}
	return s.Speak()
}
