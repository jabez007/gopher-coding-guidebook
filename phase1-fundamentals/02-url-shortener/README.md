# URL Shortener

Build a simple URL shortener using Go maps, file I/O, and JSON.

## Learning Goals

- Practice working with maps and basic I/O
- Persist data using JSON files
- Understand basic program structure and modularity
- Parse and normalize URLs with Go's `net/url` package

## Prerequisites

- Go installed (1.20+)
- Basic knowledge of CLI development in Go

## Task List

- [ ] Initialize a Go module
- [ ] Define a `map[string]string` for short -> full URLs
- [ ] Load data from a JSON file into memory
- [ ] Parse CLI args to `add`, `get`, or `list` shortlinks
- [ ] Write updated data back to the JSON file

## Stretch Goals

- [ ] Validate and normalize URLs with `net/url`
- [ ] Add support for deleting or updating shortlinks
- [ ] Implement a simple redirect HTTP server (preview of Phase 2)
- [ ] Add unit tests

## Helpful Resources

- https://pkg.go.dev/encoding/json
- https://pkg.go.dev/os
- https://pkg.go.dev/net/url
- https://gobyexample.com
