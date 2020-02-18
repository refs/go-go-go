package templates

import "github.com/google/go-github/v29/github"

// Record corresponds with a line from repos.csv (or source)
type Record struct {
	URL      string `csv:"url"`
	Category string `csv:"category"`
}

// Store are all records on source
type Store []Record

// Repository maps a repo information
type Repository struct {
	UpdatedAt   github.Timestamp // TODO make this human: Last update was X days/weeks/months/years ago
	Name        string
	Description string
	Stars       int
}

// Repositories are parsed repos ready to feed the templates
type Repositories []Repository
