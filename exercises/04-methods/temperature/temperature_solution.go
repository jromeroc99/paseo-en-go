//go:build solution

// Package temperature converts temperatures with methods on a type.
package temperature

// Celsius is a temperature in degrees Celsius.
type Celsius float64

// Fahrenheit returns c in degrees Fahrenheit.
func (c Celsius) Fahrenheit() float64 {
	return float64(c)*9/5 + 32
}

// IsFreezing reports whether water freezes at c: 0 degrees or below.
func (c Celsius) IsFreezing() bool {
	return c <= 0
}
