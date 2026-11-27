//go:build solution

// Package vat adds the value added tax to a price.
package vat

// VAT is the general rate of the tax in Spain: 21 %.
const VAT = 0.21

// WithVAT returns price with the tax added.
func WithVAT(price float64) float64 {
	return price + price*VAT
}
