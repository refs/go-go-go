package command

import (
	"context"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/gocarina/gocsv"
	"github.com/google/go-github/v29/github"
	"github.com/refs/go-go-go/pkg/config"
	"github.com/refs/go-go-go/pkg/templates"
	"github.com/refs/go-go-go/pkg/types"
	"github.com/urfave/cli/v2"
	"golang.org/x/oauth2"
)

var (
	client      *github.Client
	baseDir     string
	outFileName string
	err         error
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
			&cli.StringFlag{
				Name:    "out",
				Usage:   "Directory to write README.md",
				Aliases: []string{"o"},
				EnvVars: []string{"GOGO_OUT"},
			},
			&cli.StringFlag{
				Name:    "filename",
				Usage:   "Output file name",
				Aliases: []string{"fn"},
				EnvVars: []string{"GOGO_FILE_NAME"},
				Value:   "README.md",
			},
		},
		Before: func(c *cli.Context) error {
			ctx := context.Background()
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: c.String("token")},
			)
			tc := oauth2.NewClient(ctx, ts)
			client = github.NewClient(tc)

			if len(c.String("out")) > 0 {
				baseDir = c.String("out")
			} else {
				baseDir, err = filepath.Abs(filepath.Dir(os.Args[0]))
				if err != nil {
					log.Fatal(err)
				}
			}

			return nil
		},
		Action: func(c *cli.Context) error {
			repoStore := types.Repositories{}
			store := initStore(c.String("src"))
			for i := range store {
				repoStore = append(repoStore, store[i].Expand(client))
			}

			dst := path.Join(baseDir, c.String("filename"))

			f, err := os.Create(dst)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			return templates.Readme().Execute(f, repoStore.Categorize())
		},
	}
}

// initStore unmarshals the csv file
func initStore(dst string) types.Repositories {
	s := types.Repositories{}
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
