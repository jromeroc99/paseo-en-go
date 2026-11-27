//go:build !solution

// Package page builds an HTML page from a template.
//
// Write this function in this file:
//
//	func Render(w io.Writer, title string, items []string) error
//
// Render writes to w an HTML page with the title in an <h1> and the items
// in a list, one <li> each:
//
//	<h1>Compra</h1>
//	<ul>
//	  <li>pan</li>
//	  <li>leche</li>
//	</ul>
//
// When items is empty there is no list: write <p>No hay nada.</p> instead.
//
// Use html/template with {{range}} and {{if}}, so that a title or an item
// such as "<b>sal</b>" is shown as text and never becomes HTML. Return the
// error from Execute.
package page
