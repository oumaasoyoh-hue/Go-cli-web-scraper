# Go-cli-Scrape

A simple command-line web scraper built with Go.

## About

GoScrape is a command-line web scraper developed as a learning project to improve my understanding of Go and practical software development.

The project focuses on learning how Go can be used to work with HTTP requests, HTML documents, command-line arguments, error handling, data structures, file handling, testing, and eventually concurrency.

The goal is to start with a simple scraper and progressively improve it as I learn more Go.

## Features

### Current Features

* Accept a URL from the command line
* Validate user input
* Send HTTP GET requests
* Handle HTTP and network errors
* Read HTML responses
* Extract webpage titles
* Extract links
* Extract image URLs
* Display scraped information in the terminal

### Planned Features

* JSON output
* CSV output
* Command-line flags
* Multiple-page crawling
* Duplicate URL detection
* Configurable crawling depth
* Concurrent scraping
* Request timeouts
* Rate limiting
* Retry mechanism
* Unit and integration tests

## Requirements

* Go 1.25 or later
* Internet connection

Check your Go installation:

```bash
go version
```

## Installation

Clone the repository:

```bash
git clone <repository-url>
```

Move into the project directory:

```bash
cd goscrape
```

Initialize or download the project dependencies:

```bash
go mod tidy
```

## Usage

Run the scraper using:

```bash
go run . https://example.com
```

You can also build the application:

```bash
go build -o goscrape
```

Then run the compiled program:

```bash
./goscrape https://example.com
```

## Example

Running:

```bash
./goscrape https://example.com
```

could produce output similar to:

```text
GoScrape
--------

URL: https://example.com
Status: 200

Title:
Example Domain

Links:
1. https://example.com/about
2. https://example.com/contact

Images:
1. https://example.com/images/logo.png
```

The exact output depends on the webpage being scraped.

## Command-Line Interface

The basic command is:

```bash
goscrape <URL>
```

Example:

```bash
goscrape https://example.com
```

As the project develops, additional options will be introduced.

Planned examples:

```bash
goscrape https://example.com --links
goscrape https://example.com --images
goscrape https://example.com --metadata
goscrape https://example.com --all
goscrape https://example.com --output result.json
```

## Project Structure

The project will gradually evolve into a modular Go application.

A planned structure is:

```text
goscrape/
│
├── cmd/
│   └── goscrape/
│       └── main.go
│
├── internal/
│   ├── fetcher/
│   │   └── fetcher.go
│   │
│   ├── parser/
│   │   └── parser.go
│   │
│   └── scraper/
│       └── scraper.go
│
├── tests/
│
├── go.mod
├── go.sum
└── README.md
```

### `cmd/`

Contains the command-line application entry point.

### `internal/fetcher/`

Responsible for retrieving webpages through HTTP requests.

### `internal/parser/`

Responsible for processing HTML and extracting information.

### `internal/scraper/`

Coordinates the scraping process and combines the different components.

### `tests/`

Contains tests for the scraper.

## Learning Goals

This project is primarily designed as a practical Go learning project.

Through GoScrape, I aim to learn and practice:

* Go syntax and fundamentals
* Functions
* Structs
* Slices
* Maps
* Pointers
* Interfaces
* Packages
* Error handling
* Command-line applications
* HTTP requests
* URL parsing
* HTML parsing
* File handling
* JSON encoding and decoding
* Unit testing
* Integration testing
* Goroutines
* Channels
* Synchronization
* Concurrent programming
* Data structures and algorithms
* Project organization

## Development Roadmap

The project will be developed incrementally.

### Phase 1 — CLI

* [ ] Create the Go project
* [ ] Accept a URL from the command line
* [ ] Validate command-line input
* [ ] Display useful error messages

### Phase 2 — HTTP

* [ ] Send HTTP GET requests
* [ ] Read HTTP responses
* [ ] Handle HTTP status codes
* [ ] Handle network errors
* [ ] Add request timeouts

### Phase 3 — HTML Parsing

* [ ] Parse HTML
* [ ] Extract the page title
* [ ] Extract hyperlinks
* [ ] Extract image URLs
* [ ] Extract metadata

### Phase 4 — Output

* [ ] Improve terminal output
* [ ] Add JSON output
* [ ] Add CSV output
* [ ] Add output file support

### Phase 5 — Testing

* [ ] Add unit tests
* [ ] Test URL validation
* [ ] Test HTML parsing
* [ ] Test error handling
* [ ] Add HTTP test servers

### Phase 6 — Crawling

* [ ] Follow links
* [ ] Add crawling depth
* [ ] Track visited URLs
* [ ] Prevent duplicate requests
* [ ] Prevent infinite crawling

### Phase 7 — Concurrency

* [ ] Introduce goroutines
* [ ] Introduce channels
* [ ] Use worker pools
* [ ] Add synchronization
* [ ] Control the number of concurrent requests

### Phase 8 — Improvements

* [ ] Add rate limiting
* [ ] Add retries
* [ ] Improve logging
* [ ] Improve configuration
* [ ] Improve CLI experience

## Technologies

The project is written primarily in:

* Go
* HTTP
* HTML
* JSON

Go's standard library will be used wherever practical so that the project remains focused on understanding how the underlying functionality works.

## Error Handling

GoScrape should handle errors gracefully rather than terminating unexpectedly.

Examples include:

```text
ERROR: URL is required
ERROR: invalid URL
ERROR: unable to connect to server
ERROR: request timed out
ERROR: HTTP 404
ERROR: HTTP 500
ERROR: unable to parse HTML
```

The scraper should provide enough information for the user to understand what went wrong.

## Responsible Use

GoScrape is intended for learning and legitimate web data collection.

When using the scraper:

* Respect website terms of service.
* Respect `robots.txt` where applicable.
* Avoid sending excessive requests.
* Use reasonable request delays and rate limits.
* Do not attempt to bypass authentication or access controls.
* Do not use the scraper to overload or disrupt websites.
* Respect website owners and their data.

## Contributing

This is primarily a personal learning project, but suggestions, issues, and improvements are welcome.

If contributing:

1. Fork the repository.
2. Create a new branch.
3. Make your changes.
4. Add or update tests where appropriate.
5. Commit your changes.
6. Open a pull request.

## License

This project will be licensed under the MIT License.

## Project Status

GoScrape is an active learning project.

The project will evolve as I learn more about Go, networking, web scraping, algorithms, testing, and concurrent programming.

The roadmap will be updated as new features are completed.

---

**Built with Go as a practical project for learning, experimenting, and becoming a better Go developer.**
