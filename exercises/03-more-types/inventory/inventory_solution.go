//go:build solution

// Package inventory keeps track of the products in a small shop.
package inventory

// Item is a product and how many units of it there are.
type Item struct {
	Name     string
	Quantity int
}

// Total returns the sum of the quantities of all items.
func Total(items []Item) int {
	total := 0
	for _, item := range items {
		total += item.Quantity
	}
	return total
}

// Restock adds amount units to the item.
func Restock(item *Item, amount int) {
	item.Quantity += amount
}

// ByName returns the total quantity for each name, adding up repeated names.
func ByName(items []Item) map[string]int {
	totals := make(map[string]int)
	for _, item := range items {
		totals[item.Name] += item.Quantity
	}
	return totals
}

// Filter returns a new slice with the items for which keep returns true,
// in the same order.
func Filter(items []Item, keep func(Item) bool) []Item {
	var kept []Item
	for _, item := range items {
		if keep(item) {
			kept = append(kept, item)
		}
	}
	return kept
}
