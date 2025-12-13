# Web Crawler + Basic Search Engine using indexes

A basic search engine and web crawler in Go. It crawls pages starting from one or more seed URLs, tokenizes content, and builds an in-memory index to answer simple keyword searches. Concurrency keeps crawling and fetching snappy, and the code is split into small, focused modules.

## Features

- Concurrent crawling and fetching of web pages
- In-memory indexing of page content for keyword search
- Simple `stack` abstraction for crawl traversal
- `tokens` helpers for parsing/normalizing text
- Optional `dbOps` to persist results, if enabled
- Clear separation of concerns across files

## Project Structure

- `main.go`: Entrypoint; orchestrates crawling, indexing, and searching
- `crawler.go`: Crawl logic using stack + fetch
- `fetch.go`: HTTP client and response handling
- `stack.go`: Minimal stack implementation used by crawler
- `tokens.go`: Token parsing/normalization helpers
- `dbOps.go`: Optional persistence utilities (enable if needed)
- `go.mod`: Module definition and dependencies

## Requirements

- Go 1.20+ (recommended)
- Internet access for HTTP requests

## Setup

```bash
# From the project root
cd http_client

# Initialize (if not already)
go mod tidy
```

## Run

```bash
# Run the app
go run .

# Or explicitly
go run main.go
```

Typical flags (may vary depending on your `main.go`):

```bash
# Crawl starting from a seed URL with a max depth
go run . --seed https://example.com --depth 2

# After crawling, run a simple keyword search
go run . --search "golang concurrency"
```

## Usage Notes

- Set seed URLs and crawl depth in `main.go` or via flags.
- Tune concurrency in `fetch.go`/`crawler.go` to balance speed vs politeness.
- Consider adding URL de-duplication and robots.txt handling for production use.
- Enable `dbOps.go` if you want to persist the index to disk.

## Development

```bash
# Format code
gofmt -w .

# Vet for common issues
go vet ./...

# (Optional) Run static analysis
golangci-lint run
```

## Extending

- Add rate limiting/backoff and per-host concurrency in `fetch.go`.
- De-dup URLs, respect robots.txt/sitemaps in `crawler.go`.
- Improve tokenization (stop-words, stemming) in `tokens.go`.
- Persist and query index via `dbOps.go` (SQLite, Badger, etc.).
- Add a tiny HTTP API to serve search results.

## License

This repository is for learning/demo purposes. Add your own license if you plan to publish broadly.
