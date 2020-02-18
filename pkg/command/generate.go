package command

import (
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/refs/go-go-go/pkg/command/config"
	"github.com/refs/go-go-go/pkg/command/templates"
	"github.com/urfave/cli/v2"
)

var (
	token string
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
			token = c.String("token")
			return nil
		},
		Action: func(c *cli.Context) error {
			fmt.Println(initStore(c.String("src")))
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
