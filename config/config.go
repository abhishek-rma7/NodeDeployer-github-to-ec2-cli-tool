package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli/v2"
)

type Config struct {
	EC2User    string `json:"ec2_user"`
	EC2Address string `json:"ec2_address"`
	EC2KeyPath string `json:"ec2_key_path"`
	RepoPath   string `json:"repo_path"`
}

func GetConfig(c *cli.Context) (*Config, error) {
	configPath := c.String("config")
	if configPath != "" {
		return loadConfigFromFile(configPath)
	}

	ec2User := c.String("ec2-user")
	ec2Address := c.String("ec2-address")
	ec2KeyPath := c.String("ec2-key-path")
	repoPath := c.String("repo-path")

	if ec2User == "" || ec2Address == "" || ec2KeyPath == "" || repoPath == "" {
		return nil, fmt.Errorf("all EC2 connection details and repository path must be provided")
	}

	return &Config{
		EC2User:    ec2User,
		EC2Address: ec2Address,
		EC2KeyPath: ec2KeyPath,
		RepoPath:   repoPath,
	}, nil
}

func loadConfigFromFile(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(bytes, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func RegisterFlags(app *cli.App) {
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "config",
			Usage: "Path to configuration file",
		},
		&cli.StringFlag{
			Name:     "ec2-user",
			Usage:    "EC2 instance user",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "ec2-address",
			Usage:    "EC2 instance address",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "ec2-key-path",
			Usage:    "Path to the EC2 instance SSH key",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "repo-path",
			Usage:    "Path to the repository on the EC2 instance",
			Required: false,
		},
	}
}
