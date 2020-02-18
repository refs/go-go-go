package command

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/google/go-github/v29/github"
	"github.com/refs/go-go-go/pkg/config"
	"github.com/refs/go-go-go/pkg/templates"
	"github.com/urfave/cli/v2"
	"golang.org/x/oauth2"
)

var (
	client *github.Client
)

// GenerateCommand generates a README.md from sources
func GenerateCommand(c *config.Config) *cli.Command {
	return &cli.Command{
		Name: "generate",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "token",
				Usage:   "Github OAuth2 Token",
				Aliases: []string{"t"},
				EnvVars: []string{"GOGO_OAUTH2_TOKEN"},
			},
			&cli.StringFlag{
				Name:    "src",
				Usage:   "File containing a list of github repositories",
				Aliases: []string{"r"},
				EnvVars: []string{"GOGO_REPOS_SRC"},
				Value:   "repos.csv",
			},
		},
		Before: func(c *cli.Context) error {
			ctx := context.Background()
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: c.String("token")},
			)
			tc := oauth2.NewClient(ctx, ts)

			client = github.NewClient(tc)
			return nil
		},
		Action: func(c *cli.Context) error {
			store := initStore(c.String("src"))
			for i := range store {
				fmt.Printf("%+v\n", repoInfo(store[i]))
			}
			return nil
		},
	}
}

// initStore parses the csv into a Go type
func initStore(dst string) templates.Store {
	s := templates.Store{}
	repos, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer repos.Close()

	if err := gocsv.UnmarshalFile(repos, &s); err != nil {
		log.Fatal(err)
	}
	return s
}

func repoInfo(rec templates.Record) templates.Repository {
	owner, repo := deconstruct(rec.URL)

	ghrepo, _, err := client.Repositories.Get(context.Background(), owner, repo)
	if err != nil {
		log.Fatal(err)
	}

	r := templates.Repository{
		Name:        *ghrepo.Name,
		Description: *ghrepo.Description,
		Stars:       *ghrepo.StargazersCount,
		UpdatedAt:   ghrepo.GetUpdatedAt(),
	}

	return r
}

// returns the owner and repo name out of a github url
// i.e: https://github.com/refs/go-go-go // {"refs", "go-go-go"}
func deconstruct(rawurl string) (string, string) {
	parsed, err := url.Parse(rawurl)
	if err != nil {
		log.Fatal(err)
	}

	owner := strings.Split(parsed.Path, "/")[1]
	name := strings.Split(parsed.Path, "/")[2]

	return owner, name
}
