//go:build !solution

// Package counter counts the bytes that pass through a Writer.
//
// Write this type and this method in this file:
//
//	type CountingWriter struct {
//		W     io.Writer
//		Count int
//	}
//
//	func (c *CountingWriter) Write(p []byte) (int, error)
//
// Write passes p on to c.W and adds to c.Count the bytes that c.W says it
// wrote: the n it returns, which may be less than len(p) when it fails.
// Then it returns that same n and error. With this method, *CountingWriter
// is an io.Writer, so io.Copy can copy into it.
//
// Example:
//
//	var buf bytes.Buffer
//	cw := &CountingWriter{W: &buf}
//	io.Copy(cw, strings.NewReader("hola"))
//	// buf holds "hola" and cw.Count is 4.
package counter
