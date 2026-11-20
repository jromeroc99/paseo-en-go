//go:build !solution

// Package season maps months to seasons of the northern hemisphere.
package season

// Season returns "winter", "spring", "summer" or "autumn" for a month from
// 1 (January) to 12 (December), and "unknown" for any other number.
// Example: Season(7) returns "summer".
func Season(month int) string {
	// TODO: use a switch with several values per case and a default.
	return ""
}
