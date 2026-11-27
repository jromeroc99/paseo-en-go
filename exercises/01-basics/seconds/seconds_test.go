package seconds

import "testing"

func TestToSeconds(t *testing.T) {
	tests := []struct {
		name                    string
		hours, minutes, seconds int
		want                    int
	}{
		{name: "nothing at all", want: 0},
		{name: "only seconds", seconds: 45, want: 45},
		{name: "one minute", minutes: 1, want: 60},
		{name: "one hour", hours: 1, want: 3600},
		{name: "a bit of everything", hours: 1, minutes: 2, seconds: 3, want: 3723},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSeconds(tt.hours, tt.minutes, tt.seconds); got != tt.want {
				t.Errorf("ToSeconds(%d, %d, %d) = %d, want %d", tt.hours, tt.minutes, tt.seconds, got, tt.want)
			}
		})
	}
}
