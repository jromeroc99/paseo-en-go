package arena

import "testing"

// alloc is one call to Alloc and what it should return.
type alloc struct {
	n, align   int
	wantOffset int
	wantOK     bool
}

func TestAlloc(t *testing.T) {
	tests := []struct {
		name   string
		size   int
		allocs []alloc
	}{
		{
			name:   "the first piece starts at 0",
			size:   16,
			allocs: []alloc{{n: 4, align: 4, wantOffset: 0, wantOK: true}},
		},
		{
			name: "pieces go one after another",
			size: 16,
			allocs: []alloc{
				{n: 4, align: 4, wantOffset: 0, wantOK: true},
				{n: 4, align: 4, wantOffset: 4, wantOK: true},
			},
		},
		{
			name: "padding before an aligned piece",
			size: 16,
			allocs: []alloc{
				{n: 1, align: 1, wantOffset: 0, wantOK: true},
				{n: 8, align: 8, wantOffset: 8, wantOK: true},
			},
		},
		{
			name: "a piece that fills the arena exactly",
			size: 8,
			allocs: []alloc{
				{n: 8, align: 8, wantOffset: 0, wantOK: true},
				{n: 1, align: 1, wantOffset: 0, wantOK: false},
			},
		},
		{
			name: "a failed Alloc changes nothing",
			size: 12,
			allocs: []alloc{
				{n: 4, align: 4, wantOffset: 0, wantOK: true},
				{n: 8, align: 8, wantOffset: 0, wantOK: false},
				{n: 8, align: 4, wantOffset: 4, wantOK: true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := New(tt.size)
			for _, c := range tt.allocs {
				offset, ok := a.Alloc(c.n, c.align)
				if offset != c.wantOffset || ok != c.wantOK {
					t.Fatalf("Alloc(%d, %d) = %d, %v, want %d, %v", c.n, c.align, offset, ok, c.wantOffset, c.wantOK)
				}
			}
		})
	}
}

func TestReset(t *testing.T) {
	a := New(8)
	a.Alloc(8, 8)
	a.Reset()
	if offset, ok := a.Alloc(4, 4); offset != 0 || !ok {
		t.Errorf("after Reset, Alloc(4, 4) = %d, %v, want 0, true", offset, ok)
	}
}
