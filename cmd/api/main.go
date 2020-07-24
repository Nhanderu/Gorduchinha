package main

import (
	"os"
	"strconv"

	"github.com/Nhanderu/gorduchinha/app"
	"github.com/Nhanderu/gorduchinha/cmd/api/server"
)

var (
	AppVersion string
)

func main() {

	serverPort, _ := strconv.Atoi(os.Getenv("PORT"))
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	app := app.New(env, AppVersion)
	if serverPort != 0 {
		app.Config.Server.Port = serverPort
	}

	app.Logger.Infof("Running server at port %d.", app.Config.Server.Port)
	err := server.Run(
		app.Config.App.Version,
		app.Config.Server.Port,
		app.Config.Server.Prefix,
		app.Logger,
		app.Services().NewTeam(),
		app.Services().NewChamp(),
		app.Services().NewScraper(),
	)
	app.EndAsErr(err, "Could not run server.", app.Logger.InfoWriter(), app.Logger.ErrorWriter())
}
