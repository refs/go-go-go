package types

import (
	"github.com/google/go-github/v29/github"
)

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
	Owner       string
	Name        string
	URL         string
	Description string
	Category    string
	Stargazers  int
}

// Repositories are parsed repos ready to feed the templates
type Repositories []Repository

// Categorize sorts repositories by it's declared category
func (r Repositories) Categorize() map[string][]Repository {
	ret := map[string][]Repository{}

	for i := range r {
		cat := r[i].Category
		ret[cat] = append(ret[cat], r[i])
	}

	return ret
}
