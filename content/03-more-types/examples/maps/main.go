package main

import "fmt"

func main() {
	ages := make(map[string]int)

	ages["Ana"] = 31
	ages["Luis"] = 27

	fmt.Println(ages["Ana"])
	fmt.Println(len(ages))
}
