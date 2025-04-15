# Concurrent Scraper

Build a simple web scraper that demonstrates Go's concurrency model.

## Learning Goals

- Understand how to use goroutines and channels
- Use `sync.WaitGroup` to manage concurrency
- Perform basic HTTP requests with `net/http`
- Extract information from HTML using Go packages

## Prerequisites

- Go installed (1.20+)
- Basic understanding of goroutines and channels

## Task List

- [ ] Initialize a Go module
- [ ] Accept a list of URLs from command line or a file
- [ ] Use goroutines to fetch page content concurrently
- [ ] Extract and print page titles or meta descriptions
- [ ] Coordinate goroutines with channels and WaitGroup

## Stretch Goals

- [ ] Add CLI flags for concurrency level
- [ ] Save scraped data to a file (CSV or JSON)
- [ ] Handle and retry failed requests

## Helpful Resources

- https://gobyexample.com/goroutines
- https://gobyexample.com/channels
- https://pkg.go.dev/net/http
- https://pkg.go.dev/golang.org/x/net/html
