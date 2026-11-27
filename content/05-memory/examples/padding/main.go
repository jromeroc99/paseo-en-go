package main

import (
	"fmt"
	"unsafe"
)

type Loose struct {
	A bool
	B int64
	C bool
}

type Tight struct {
	B int64
	A bool
	C bool
}

func main() {
	fmt.Println(unsafe.Sizeof(Loose{}), unsafe.Sizeof(Tight{}))
}
