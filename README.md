# Paseo en Go

A modern, beginner-friendly tour of Go, in Spanish. Lessons live on a static
website; exercises are solved locally and validated with `go test`.

Requires Go 1.22+ (1.26 for developing the site). No other dependencies.

## Development

```bash
go run ./cmd/paseo serve        # http://localhost:8080
```

Content is embedded in the binary: restart `serve` to see changes.

```bash
go vet ./...
go test ./...                                   # site, domain and content checks
cd exercises && go test -tags solution ./...    # reference solutions must pass
```

## Production

The site is fully static. Build it and upload `dist/` to any static host
(GitHub Pages, Netlify, Cloudflare Pages):

```bash
go run ./cmd/paseo build -out dist                     # served from the domain root
go run ./cmd/paseo build -out dist -base /paseo-en-go/ # served from a subpath
```

## For students

```bash
git clone <repo-url>
cd paseo-en-go/exercises
go test ./00-setup/hello    # fails until you solve the exercise
```

See [CLAUDE.md](CLAUDE.md) for architecture and contribution rules.
