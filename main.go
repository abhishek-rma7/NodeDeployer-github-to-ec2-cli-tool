package main

import (
	"NodeDeployer/cmd"
	"NodeDeployer/config"
	"NodeDeployer/internal/utils"
	"os"
)

func main() {
	utils.InitLoggers()

	app := cmd.NewApp()
	config.RegisterFlags(app)

	err := app.Run(os.Args)
	if err != nil {
		utils.ErrorLogger.Fatal(err)
	}
}
