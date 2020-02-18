package command

import (
	"github.com/refs/go-go-go/pkg/command/config"
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
		},
		Before: func(c *cli.Context) error {
			token = c.String("token")
			return nil
		},
		Action: func(c *cli.Context) error {
			// now we have the token loaded
			return nil
		},
	}
}
