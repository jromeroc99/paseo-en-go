//go:build solution

// Package echo is a TCP server that sends back every line it receives.
package echo

import (
	"io"
	"net"
)

// Serve accepts connections from l and echoes each one in its own
// goroutine. It returns the error from Accept once l is closed.
func Serve(l net.Listener) error {
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go Echo(conn)
	}
}

// Echo writes back everything it reads from conn, then closes it.
func Echo(conn net.Conn) {
	defer conn.Close()
	io.Copy(conn, conn)
}
