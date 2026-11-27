//go:build !solution

// Package vat adds the value added tax to a price.
//
// Write this constant and this function in this file:
//
//	const VAT = 0.21
//
//	func WithVAT(price float64) float64
//
// VAT is the general rate of the tax in Spain: 21 %. It never changes
// while the program runs, so it is a constant, not a variable.
// WithVAT returns price with the tax added.
// Example: WithVAT(100) returns 121.
package vat
