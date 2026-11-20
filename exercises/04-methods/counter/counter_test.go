package counter

import "testing"

func TestIncrement(t *testing.T) {
	tests := []struct {
		increments int
		want       int
	}{
		{increments: 0, want: 0},
		{increments: 1, want: 1},
		{increments: 3, want: 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			var c Counter
			for i := 0; i < tt.increments; i++ {
				c.Increment()
			}
			if got := c.Value(); got != tt.want {
				t.Errorf("Value() after %d calls to Increment() = %d, want %d", tt.increments, got, tt.want)
			}
		})
	}
}

func TestReset(t *testing.T) {
	tests := []struct {
		increments int
	}{
		{increments: 0},
		{increments: 2},
		{increments: 5},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			var c Counter
			for i := 0; i < tt.increments; i++ {
				c.Increment()
			}
			c.Reset()
			c.Increment()
			if got := c.Value(); got != 1 {
				t.Errorf("Value() after %d increments, Reset() and one more Increment() = %d, want 1", tt.increments, got)
			}
		})
	}
}
