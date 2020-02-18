package types

import (
	"context"
	"log"
	"net/url"
	"strings"

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

// Deconstruct returns the owner and repo name out of a github url
// i.e: https://github.com/refs/go-go-go // {"refs", "go-go-go"}
func (r Repository) Deconstruct() (string, string) {
	parsed, err := url.Parse(r.URL)
	if err != nil {
		log.Fatal(err)
	}

	owner := strings.Split(parsed.Path, "/")[1]
	name := strings.Split(parsed.Path, "/")[2]

	return owner, name
}

// Expand feeds the repository with information pulled from Github
func (r Repository) Expand(c *github.Client) Repository {
	owner, repo := r.Deconstruct()

	ghrepo, _, err := c.Repositories.Get(context.Background(), owner, repo)
	if err != nil {
		log.Fatal(err)
	}

	return Repository{
		Name:        *ghrepo.Name,
		Owner:       owner,
		URL:         r.URL,
		Category:    r.Category,
		Description: *ghrepo.Description,
		Stargazers:  *ghrepo.StargazersCount,
		UpdatedAt:   ghrepo.GetUpdatedAt(),
	}
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
