package main

import "fmt"

// kind classifies a day of the week, where 1 is Monday.
func kind(day int) string {
	switch day {
	case 1, 2, 3, 4:
		return "laborable"
	case 5:
		return "viernes"
	case 6, 7:
		return "fin de semana"
	default:
		return "no es un día"
	}
}

func main() {
	fmt.Println(kind(2))
	fmt.Println(kind(5))
	fmt.Println(kind(7))
	fmt.Println(kind(9))
}
