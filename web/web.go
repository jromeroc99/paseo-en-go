// Package web embeds the page templates and static assets of the site.
package web

import "embed"

// FS contains templates/ and static/.
//
//go:embed templates static
var FS embed.FS
