package kv

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"testing"
	"time"
)

const timeout = 500 * time.Millisecond

// client is one open connection to a Store.
type client struct {
	conn    net.Conn
	replies *bufio.Scanner
}

// connect opens a new connection that s handles.
func connect(t *testing.T, s *Store) *client {
	t.Helper()
	conn, server := net.Pipe()
	t.Cleanup(func() { conn.Close() })
	go s.Handle(server)
	conn.SetDeadline(time.Now().Add(timeout))
	return &client{conn: conn, replies: bufio.NewScanner(conn)}
}

// send writes one line and returns the line that comes back.
func (c *client) send(t *testing.T, line string) string {
	t.Helper()
	fmt.Fprintln(c.conn, line)
	if !c.replies.Scan() {
		t.Fatalf("sent %q and got no answer: answer every line with one line ending in \\n", line)
	}
	return c.replies.Text()
}

func TestCommands(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  []string
	}{
		{name: "set then get", lines: []string{"SET color azul", "GET color"}, want: []string{"OK", "azul"}},
		{name: "value with spaces", lines: []string{"SET saludo hola mundo", "GET saludo"}, want: []string{"OK", "hola mundo"}},
		{name: "unknown key", lines: []string{"GET nada"}, want: []string{"NOT FOUND"}},
		{name: "set again overwrites", lines: []string{"SET n 1", "SET n 2", "GET n"}, want: []string{"OK", "OK", "2"}},
		{name: "unknown command", lines: []string{"DEL color"}, want: []string{"ERROR"}},
		{name: "set without value", lines: []string{"SET color"}, want: []string{"ERROR"}},
		{name: "empty line", lines: []string{""}, want: []string{"ERROR"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := connect(t, NewStore())
			for i, line := range tt.lines {
				if got := c.send(t, line); got != tt.want[i] {
					t.Errorf("sent %q, got %q, want %q", line, got, tt.want[i])
				}
			}
		})
	}
}

func TestConnectionsShareTheStore(t *testing.T) {
	s := NewStore()
	first := connect(t, s)
	second := connect(t, s)

	first.send(t, "SET ciudad Madrid")
	if got := second.send(t, "GET ciudad"); got != "Madrid" {
		t.Errorf("SET on one connection, then GET on another: got %q, want %q", got, "Madrid")
	}
}

// closeRecorder is a connection that remembers whether it was closed.
type closeRecorder struct {
	net.Conn
	closed chan struct{}
}

func (c *closeRecorder) Close() error {
	close(c.closed)
	return c.Conn.Close()
}

func TestHandleClosesTheConnection(t *testing.T) {
	conn, server := net.Pipe()
	recorder := &closeRecorder{Conn: server, closed: make(chan struct{})}
	go NewStore().Handle(recorder)

	conn.Close()
	select {
	case <-recorder.closed:
	case <-time.After(timeout):
		t.Fatal("the client hung up, but Handle did not close its connection: add defer conn.Close()")
	}
}

func TestManyClientsAtOnce(t *testing.T) {
	s := NewStore()
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		c := connect(t, s)
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Fprintf(c.conn, "SET clave%d %d\n", i, i)
			c.replies.Scan()
			fmt.Fprintf(c.conn, "GET clave%d\n", i)
			c.replies.Scan()
			if got, want := c.replies.Text(), fmt.Sprint(i); got != want {
				t.Errorf("client %d: GET clave%d = %q, want %q", i, i, got, want)
			}
		}()
	}
	wg.Wait()
}
