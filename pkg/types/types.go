package types

import (
	"github.com/google/go-github/v29/github"
)

// Repository maps a repo information
type Repository struct {
	URL         string `csv:"url"`
	Category    string `csv:"category"`
	Owner       string
	Name        string
	Description string
	Stargazers  int
	UpdatedAt   github.Timestamp // TODO make this human: Last update was X days/weeks/months/years ago
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
