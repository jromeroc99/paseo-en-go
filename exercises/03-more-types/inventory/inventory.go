//go:build !solution

// Package inventory keeps track of the products in a small shop.
package inventory

// Item is a product and how many units of it there are.
type Item struct {
	Name     string
	Quantity int
}

// Total returns the sum of the quantities of all items.
// Example: Total([]Item{{"manzana", 3}, {"pera", 2}}) returns 5.
func Total(items []Item) int {
	// TODO: range over items and add up each Quantity.
	return 0
}

// Restock adds amount units to the item.
// Example: with item := Item{"pera", 0}, after Restock(&item, 2) item.Quantity is 2.
func Restock(item *Item, amount int) {
	// TODO: change the Quantity field through the pointer.
}

// ByName returns the total quantity for each name, adding up repeated names.
// Example: ByName([]Item{{"manzana", 3}, {"manzana", 4}}) returns map[string]int{"manzana": 7}.
func ByName(items []Item) map[string]int {
	// TODO: make a map and add each Quantity to its Name.
	return nil
}

// Filter returns a new slice with the items for which keep returns true,
// in the same order. It does not modify items.
// Example: Filter(items, func(i Item) bool { return i.Quantity > 0 })
// returns only the items in stock.
func Filter(items []Item, keep func(Item) bool) []Item {
	// TODO: range over items and append the ones that keep accepts.
	return nil
}
