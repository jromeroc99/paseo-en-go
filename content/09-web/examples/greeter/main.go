package main

import (
	"fmt"
	"log"
	"net"
)

func greet(conn net.Conn) {
	defer conn.Close()
	fmt.Fprintf(conn, "Hola, %s\n", conn.RemoteAddr())
}

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go greet(conn)
	}
}
