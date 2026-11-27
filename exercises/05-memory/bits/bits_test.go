package bits

import (
	"math"
	"testing"
)

func TestBitsNeeded(t *testing.T) {
	tests := []struct {
		name string
		n    uint64
		want int
	}{
		{name: "zero fits in a byte", n: 0, want: 8},
		{name: "the largest uint8", n: 255, want: 8},
		{name: "one more than a uint8", n: 256, want: 16},
		{name: "the largest uint16", n: math.MaxUint16, want: 16},
		{name: "one more than a uint16", n: math.MaxUint16 + 1, want: 32},
		{name: "the largest uint32", n: math.MaxUint32, want: 32},
		{name: "one more than a uint32", n: math.MaxUint32 + 1, want: 64},
		{name: "the largest uint64", n: math.MaxUint64, want: 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BitsNeeded(tt.n); got != tt.want {
				t.Errorf("BitsNeeded(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
