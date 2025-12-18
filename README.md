# Web Crawler + Basic Search Engine using indexes

A basic search engine and web crawler in Go. It crawls pages starting from one or more seed URLs, tokenizes content, and builds an in-memory index to answer simple keyword searches. Concurrency keeps crawling and fetching snappy, and the code is split into small, focused modules.

## Features

- Concurrent crawling and fetching of web pages
- In-memory indexing of page content for keyword search
- **TF-IDF ranking** for relevance-based search results
- Simple `stack` abstraction for crawl traversal
- `tokens` helpers for parsing/normalizing text
- Optional `dbOps` to persist results, if enabled
- Clear separation of concerns across files

## Search Theory: TF-IDF

This search engine uses **TF-IDF (Term Frequency-Inverse Document Frequency)** to rank search results by relevance.

### How It Works

**TF-IDF** is a numerical statistic that reflects how important a word is to a document in a collection of documents. It combines two metrics:

#### 1. Term Frequency (TF)

Measures how often a term appears in a document:

```
TF(t, d) = (Number of times term t appears in document d) / (Total number of terms in document d)
```

A higher TF means the term is more prominent in that specific document.

#### 2. Inverse Document Frequency (IDF)

Measures how rare or common a term is across all documents:

```
IDF(t) = log(Total number of documents / Number of documents containing term t)
```

A higher IDF means the term is rarer and likely more meaningful (e.g., "algorithm" vs "the").

#### 3. TF-IDF Score

The final relevance score combines both:

```
TF-IDF(t, d) = TF(t, d) × IDF(t)
```

For multi-term queries, the scores are summed or averaged across all query terms.

### Why TF-IDF?

- **Simple & Effective**: Works well for keyword-based search without machine learning
- **Balances Frequency & Rarity**: Common words (like "the", "is") get low scores; meaningful terms get high scores
- **Fast Computation**: Can be pre-computed and stored in the index for quick lookups
- **Language-Agnostic**: Works with any tokenized text

### Example

For query "golang concurrency":

- Pages with both terms score highest
- Pages where these terms appear frequently (high TF) rank higher
- Pages where these terms are rare across the corpus (high IDF) also rank higher
- Common words are naturally de-emphasized

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
cd WebCrawler

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
