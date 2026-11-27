//go:build !solution

// Package ipaddr prints IP addresses the way people write them.
//
// Write this type and this method in this file:
//
//	type IPAddr [4]byte
//
//	func (ip IPAddr) String() string
//
// IPAddr is an IPv4 address: four bytes.
//
// String returns ip as four numbers separated by dots, so fmt prints it
// nicely.
// Example: IPAddr{127, 0, 0, 1}.String() returns "127.0.0.1".
//
// Hint: use fmt.Sprintf with the four bytes.
package ipaddr
