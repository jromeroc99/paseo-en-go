package zero

import "testing"

func TestZeros(t *testing.T) {
	i, f, s, b := Zeros()
	if i != 0 {
		t.Errorf("Zeros() gives the int %d, want 0", i)
	}
	if f != 0 {
		t.Errorf("Zeros() gives the float64 %v, want 0", f)
	}
	if s != "" {
		t.Errorf("Zeros() gives the string %q, want %q", s, "")
	}
	if b {
		t.Errorf("Zeros() gives the bool %v, want false", b)
	}
}
