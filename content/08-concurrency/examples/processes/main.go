package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("proceso:", os.Getpid())
	fmt.Println("núcleos:", runtime.NumCPU())
}
