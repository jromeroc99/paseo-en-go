package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int8 = 1
	var b int32 = 2
	var c int64 = 3
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(c))
}
