package divmod

import "testing"

func TestDivMod(t *testing.T) {
	tests := []struct {
		a, b          int
		wantQuotient  int
		wantRemainder int
	}{
		{a: 7, b: 2, wantQuotient: 3, wantRemainder: 1},
		{a: 10, b: 5, wantQuotient: 2, wantRemainder: 0},
		{a: 3, b: 7, wantQuotient: 0, wantRemainder: 3},
		{a: -7, b: 2, wantQuotient: -3, wantRemainder: -1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			q, r := DivMod(tt.a, tt.b)
			if q != tt.wantQuotient || r != tt.wantRemainder {
				t.Errorf("DivMod(%d, %d) = (%d, %d), want (%d, %d)",
					tt.a, tt.b, q, r, tt.wantQuotient, tt.wantRemainder)
			}
		})
	}
}
