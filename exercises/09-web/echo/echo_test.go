package echo

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"testing"
	"time"
)

const timeout = 500 * time.Millisecond

func TestEcho(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
	}{
		{name: "one line", lines: []string{"hola"}},
		{name: "several lines", lines: []string{"uno", "dos", "tres"}},
		{name: "accents", lines: []string{"¿qué tal?"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, server := net.Pipe()
			defer client.Close()
			go Echo(server)

			client.SetDeadline(time.Now().Add(timeout))
			replies := bufio.NewScanner(client)
			for _, line := range tt.lines {
				fmt.Fprintln(client, line)
				if !replies.Scan() {
					t.Fatalf("sent %q and got no reply: %v", line, replies.Err())
				}
				if got := replies.Text(); got != line {
					t.Errorf("sent %q, got back %q", line, got)
				}
			}
		})
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

func TestEchoClosesTheConnection(t *testing.T) {
	client, server := net.Pipe()
	recorder := &closeRecorder{Conn: server, closed: make(chan struct{})}
	go Echo(recorder)

	client.Close()
	select {
	case <-recorder.closed:
	case <-time.After(timeout):
		t.Fatal("the client hung up, but Echo did not close its connection: add defer conn.Close()")
	}
}

func TestServe(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	stopped := make(chan error, 1)
	go func() { stopped <- Serve(l) }()

	// Two clients talk at the same time: the second one only gets an
	// answer if Serve does not wait for the first one to finish.
	var wg sync.WaitGroup
	for _, name := range []string{"ana", "luis"} {
		conn, err := net.DialTimeout("tcp", l.Addr().String(), timeout)
		if err != nil {
			t.Fatalf("could not connect to the server: %v", err)
		}
		defer conn.Close()
		conn.SetDeadline(time.Now().Add(timeout))
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Fprintln(conn, name)
			reply := bufio.NewScanner(conn)
			if !reply.Scan() || reply.Text() != name {
				t.Errorf("client %q got %q back, want %q: run Echo in its own goroutine for each connection", name, reply.Text(), name)
			}
		}()
	}
	wg.Wait()

	l.Close()
	select {
	case err := <-stopped:
		if err == nil {
			t.Error("Serve returned nil after the listener closed, want the error from Accept")
		}
	case <-time.After(timeout):
		t.Fatal("the listener is closed, but Serve did not return")
	}
}
