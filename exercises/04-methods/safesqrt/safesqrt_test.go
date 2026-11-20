package safesqrt

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	tests := []struct {
		x       float64
		want    float64
		wantErr error
	}{
		{x: 4, want: 2, wantErr: nil},
		{x: 0, want: 0, wantErr: nil},
		{x: 2, want: math.Sqrt2, wantErr: nil},
		{x: -2, want: 0, wantErr: ErrNegativeSqrt(-2)},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := Sqrt(tt.x)
			if math.Abs(got-tt.want) > 1e-9 || err != tt.wantErr {
				t.Errorf("Sqrt(%v) = (%v, %v), want (%v, %v)", tt.x, got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestErrorMessage(t *testing.T) {
	tests := []struct {
		err  ErrNegativeSqrt
		want string
	}{
		{err: ErrNegativeSqrt(-2), want: "cannot Sqrt negative number: -2"},
		{err: ErrNegativeSqrt(-0.5), want: "cannot Sqrt negative number: -0.5"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("ErrNegativeSqrt(%v).Error() = %q, want %q", float64(tt.err), got, tt.want)
			}
		})
	}
}
