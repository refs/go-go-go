package types

import (
	"context"
	"log"
	"net/url"
	"strings"

	"github.com/google/go-github/v29/github"
)

const (
	// DefaultFilename is the output file name
	DefaultFilename = "README.md"
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

// IRepo capture the basic operations on a repository
type IRepo interface {
	ROwner() string
	RName() string
	Expand(*github.Client) Repository
}

// ROwner returns the Github repository owner's name
func (r Repository) ROwner() string {
	parsed, err := url.Parse(r.URL)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(parsed.Path, "/")[1]
}

// RName returns the Github repository name
func (r Repository) RName() string {
	parsed, err := url.Parse(r.URL)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(parsed.Path, "/")[2]
}

// Expand feeds the repository with information pulled from Github
func (r Repository) Expand(c *github.Client) Repository {
	owner := r.ROwner()
	name := r.RName()

	ghrepo, _, err := c.Repositories.Get(context.Background(), owner, name)
	if err != nil {
		log.Fatal(err)
	}

	return Repository{
		Name:        *ghrepo.Name,
		Owner:       owner,
		URL:         r.URL,
		Category:    r.Category,
		Description: ghrepo.GetDescription(),
		Stargazers:  ghrepo.GetStargazersCount(),
		UpdatedAt:   ghrepo.GetUpdatedAt(),
	}
}

// Repositories are parsed repos ready to feed the templates
type Repositories []Repository

// Categorize sorts repositories by it's declared category
func (r Repositories) Categorize() map[string][]Repository {
	repo := map[string][]Repository{}

	for i := range r {
		repo[r[i].Category] = append(repo[r[i].Category], r[i])
	}

	return repo
}
