//go:build solution

// Package season maps months to seasons of the northern hemisphere.
package season

// Season returns "winter", "spring", "summer" or "autumn" for a month from
// 1 (January) to 12 (December), and "unknown" for any other number.
func Season(month int) string {
	switch month {
	case 12, 1, 2:
		return "winter"
	case 3, 4, 5:
		return "spring"
	case 6, 7, 8:
		return "summer"
	case 9, 10, 11:
		return "autumn"
	default:
		return "unknown"
	}
}
