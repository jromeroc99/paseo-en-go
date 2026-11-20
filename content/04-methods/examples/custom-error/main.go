package main

import (
	"errors"
	"fmt"
)

// ErrTooYoung says why someone cannot sign up.
type ErrTooYoung struct {
	Age int
}

func (e ErrTooYoung) Error() string {
	return fmt.Sprintf("demasiado joven: %d años", e.Age)
}

func signUp(name string, age int) error {
	if name == "" {
		return errors.New("falta el nombre")
	}
	if age < 18 {
		return ErrTooYoung{age}
	}
	return nil
}

func main() {
	fmt.Println(signUp("", 30))
	fmt.Println(signUp("Leo", 15))
	fmt.Println(signUp("Ana", 42))
}
