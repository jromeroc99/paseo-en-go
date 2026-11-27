//go:build !solution

// Package inventory keeps track of the products in a small shop.
//
// Write this type and these functions in this file:
//
//	type Item struct {
//		Name     string
//		Quantity int
//	}
//
//	func Total(items []Item) int
//	func Restock(item *Item, amount int)
//	func ByName(items []Item) map[string]int
//	func Filter(items []Item, keep func(Item) bool) []Item
//
// Item is a product and how many units of it there are.
//
// Total returns the sum of the quantities of all items.
// Example: Total([]Item{{"manzana", 3}, {"pera", 2}}) returns 5.
//
// Restock adds amount units to the item.
// Example: with item := Item{"pera", 0}, after Restock(&item, 2)
// item.Quantity is 2.
//
// ByName returns the total quantity for each name, adding up repeated names.
// Example: ByName([]Item{{"manzana", 3}, {"manzana", 4}}) returns
// map[string]int{"manzana": 7}.
//
// Filter returns a new slice with the items for which keep returns true,
// in the same order. It does not modify items.
// Example: Filter(items, func(i Item) bool { return i.Quantity > 0 })
// returns only the items in stock.
//
// Hint: Total, ByName and Filter range over items. Restock changes the
// Quantity field through the pointer.
package inventory
