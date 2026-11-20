package areader

import "testing"

func TestRead(t *testing.T) {
	tests := []struct {
		size int
	}{
		{size: 0},
		{size: 1},
		{size: 8},
		{size: 1024},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			b := make([]byte, tt.size)
			n, err := AReader{}.Read(b)
			if n != tt.size || err != nil {
				t.Fatalf("Read(buffer of %d bytes) = (%d, %v), want (%d, <nil>)", tt.size, n, err, tt.size)
			}
			for i, c := range b {
				if c != 'A' {
					t.Fatalf("after Read(buffer of %d bytes), byte %d is %q, want 'A'", tt.size, i, c)
				}
			}
		})
	}
}
