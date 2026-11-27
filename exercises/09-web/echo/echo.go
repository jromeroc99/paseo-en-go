//go:build !solution

// Package echo is a TCP server that sends back every line it receives.
//
// Write these functions in this file:
//
//	func Serve(l net.Listener) error
//	func Echo(conn net.Conn)
//
// Serve accepts connections from l, one after another, and hands each one
// to Echo in its own goroutine, so many clients can talk at the same time.
// When l is closed, Accept fails: Serve stops and returns that error.
//
// Echo reads from conn and writes back exactly what it reads, until the
// client stops sending. Then it closes conn.
//
// Example: a client sends "hola\n" and receives "hola\n".
package echo
