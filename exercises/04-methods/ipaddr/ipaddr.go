//go:build !solution

// Package ipaddr prints IP addresses the way people write them.
package ipaddr

// IPAddr is an IPv4 address: four bytes.
type IPAddr [4]byte

// String returns ip as four numbers separated by dots, so fmt prints it nicely.
// Example: IPAddr{127, 0, 0, 1}.String() returns "127.0.0.1".
func (ip IPAddr) String() string {
	// TODO: use fmt.Sprintf with the four bytes.
	return ""
}
