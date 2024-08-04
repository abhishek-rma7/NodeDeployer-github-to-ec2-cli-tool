package cmd

import (
	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	app := &cli.App{
		Name:  "deploy-tool",
		Usage: "A simple CLI tool to deploy Node.js projects from GitHub to EC2",
		Commands: []*cli.Command{
			{
				Name:   "update",
				Usage:  "Push updates to GitHub and deploy to EC2",
				Action: Update,
			},
		},
	}
	return app
}
