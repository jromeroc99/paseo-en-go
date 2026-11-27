package main

import "fmt"

// disk stands in for a file: it counts the writes that reach it.
type disk struct {
	writes int
}

func (d *disk) Write(p []byte) (int, error) {
	d.writes++
	return len(p), nil
}

func main() {
	words := []string{"uno ", "dos ", "tres "}

	slow := &disk{}
	for _, word := range words {
		slow.Write([]byte(word))
	}

	fast := &disk{}
	var buf []byte
	for _, word := range words {
		buf = append(buf, word...)
	}
	fast.Write(buf)

	fmt.Println(slow.writes, fast.writes)
}
