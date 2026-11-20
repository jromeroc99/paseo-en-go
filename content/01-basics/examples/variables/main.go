package main

import "fmt"

// Package-level variables are visible to every function in the package.
var language = "Go"

var c, python, java bool

func main() {
	var i, j int = 1, 2
	fmt.Println(language, i, j, c, python, java)
}
