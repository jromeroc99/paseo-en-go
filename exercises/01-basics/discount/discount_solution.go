//go:build solution

// Package discount works out prices with decimals.
package discount

// Discount returns price after taking off percent per cent of it.
func Discount(price, percent float64) float64 {
	return price - price*percent/100
}
