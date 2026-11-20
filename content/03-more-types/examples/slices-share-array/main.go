package main

import "fmt"

func main() {
	names := [4]string{"Ana", "Luis", "Marta", "Pablo"}

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX" // b[0] is names[1], and so is a[1]
	fmt.Println(a, b)
	fmt.Println(names)
}
