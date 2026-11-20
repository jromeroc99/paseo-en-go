//go:build solution

// Package ipaddr prints IP addresses the way people write them.
package ipaddr

import "fmt"

// IPAddr is an IPv4 address: four bytes.
type IPAddr [4]byte

// String returns ip as four numbers separated by dots.
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}
