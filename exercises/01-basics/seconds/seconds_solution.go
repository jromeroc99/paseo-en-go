//go:build solution

// Package seconds turns a time of day into seconds.
package seconds

// ToSeconds returns how many seconds there are in hours, minutes and
// seconds added together.
func ToSeconds(hours, minutes, seconds int) int {
	fromHours := hours * 3600
	fromMinutes := minutes * 60
	return fromHours + fromMinutes + seconds
}
