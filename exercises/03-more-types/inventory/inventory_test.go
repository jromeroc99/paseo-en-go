package inventory

import (
	"maps"
	"slices"
	"testing"
)

func TestTotal(t *testing.T) {
	tests := []struct {
		items []Item
		want  int
	}{
		{items: nil, want: 0},
		{items: []Item{{Name: "manzana", Quantity: 3}}, want: 3},
		{items: []Item{{Name: "manzana", Quantity: 3}, {Name: "pera", Quantity: 2}}, want: 5},
		{items: []Item{{Name: "pera", Quantity: 0}, {Name: "uva", Quantity: 10}}, want: 10},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Total(tt.items); got != tt.want {
				t.Errorf("Total(%v) = %d, want %d", tt.items, got, tt.want)
			}
		})
	}
}

func TestRestock(t *testing.T) {
	tests := []struct {
		item   Item
		amount int
		want   int
	}{
		{item: Item{Name: "manzana", Quantity: 3}, amount: 5, want: 8},
		{item: Item{Name: "pera", Quantity: 0}, amount: 2, want: 2},
		{item: Item{Name: "uva", Quantity: 4}, amount: 0, want: 4},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			item := tt.item
			Restock(&item, tt.amount)
			if item.Quantity != tt.want {
				t.Errorf("Restock(&%v, %d): Quantity = %d, want %d", tt.item, tt.amount, item.Quantity, tt.want)
			}
		})
	}
}

func TestByName(t *testing.T) {
	tests := []struct {
		items []Item
		want  map[string]int
	}{
		{items: nil, want: map[string]int{}},
		{items: []Item{{Name: "manzana", Quantity: 3}, {Name: "pera", Quantity: 2}},
			want: map[string]int{"manzana": 3, "pera": 2}},
		{items: []Item{{Name: "manzana", Quantity: 3}, {Name: "pera", Quantity: 2}, {Name: "manzana", Quantity: 4}},
			want: map[string]int{"manzana": 7, "pera": 2}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ByName(tt.items); !maps.Equal(got, tt.want) {
				t.Errorf("ByName(%v) = %v, want %v", tt.items, got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	inStock := func(item Item) bool { return item.Quantity > 0 }
	isApple := func(item Item) bool { return item.Name == "manzana" }
	items := []Item{{Name: "manzana", Quantity: 3}, {Name: "pera", Quantity: 0}, {Name: "manzana", Quantity: 0}}

	tests := []struct {
		items    []Item
		keepName string
		keep     func(Item) bool
		want     []Item
	}{
		{items: items, keepName: "inStock", keep: inStock, want: []Item{{Name: "manzana", Quantity: 3}}},
		{items: items, keepName: "isApple", keep: isApple,
			want: []Item{{Name: "manzana", Quantity: 3}, {Name: "manzana", Quantity: 0}}},
		{items: nil, keepName: "inStock", keep: inStock, want: []Item{}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			original := slices.Clone(tt.items)
			got := Filter(tt.items, tt.keep)
			if !slices.Equal(got, tt.want) {
				t.Errorf("Filter(%v, %s) = %v, want %v", original, tt.keepName, got, tt.want)
			}
			if !slices.Equal(tt.items, original) {
				t.Errorf("Filter(%v, %s) changed its input to %v; build a new slice instead",
					original, tt.keepName, tt.items)
			}
		})
	}
}
