//go:build !solution

// Package temperature converts temperatures with methods on a type.
package temperature

// Celsius is a temperature in degrees Celsius.
type Celsius float64

// Fahrenheit returns c in degrees Fahrenheit.
// Example: Celsius(100).Fahrenheit() returns 212.
func (c Celsius) Fahrenheit() float64 {
	// TODO: multiply by 9, divide by 5 and add 32.
	return 0
}

// IsFreezing reports whether water freezes at c: 0 degrees or below.
// Example: Celsius(-5).IsFreezing() returns true.
func (c Celsius) IsFreezing() bool {
	// TODO: compare c with 0.
	return false
}
