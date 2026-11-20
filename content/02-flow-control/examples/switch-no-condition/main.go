package main

import "fmt"

// grade turns a score from 0 to 10 into a word.
func grade(score int) string {
	switch {
	case score >= 9:
		return "sobresaliente"
	case score >= 7:
		return "notable"
	case score >= 5:
		return "aprobado"
	default:
		return "suspenso"
	}
}

func main() {
	fmt.Println(grade(10), grade(7), grade(5), grade(2))
}
