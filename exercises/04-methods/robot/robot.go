//go:build !solution

// Package robot tells an empty interface from one that holds a nil pointer.
//
// Write these types, this method and this function in this file:
//
//	type Speaker interface {
//		Speak() string
//	}
//
//	type Robot struct {
//		Name string
//	}
//
//	func (r *Robot) Speak() string
//	func Say(s Speaker) string
//
// Speaker is anything that can say something.
//
// Speak returns "I am " followed by the robot's name. A nil *Robot is a
// robot that is switched off, so it returns "zzz" instead.
// Example: (&Robot{Name: "R2"}).Speak() returns "I am R2".
//
// Say returns what s says, or "nobody is there" when s is a nil interface.
// Example: Say(nil) returns "nobody is there", but with var r *Robot,
// Say(r) returns "zzz": the interface holds a type, so it is not nil.
//
// Hint: in Speak, check r == nil before reading r.Name. In Say, check
// s == nil before calling s.Speak().
package robot
