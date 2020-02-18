package cmd

import (
	"os"

	"github.com/refs/go-go-go/pkg/command"
	"github.com/refs/go-go-go/pkg/config"
	"github.com/urfave/cli/v2"
)

// RootCommand is the entry point
func RootCommand() error {
	app := &cli.App{
		Commands: []*cli.Command{command.GenerateCommand(config.New())},
	}

	return app.Run(os.Args)
}
