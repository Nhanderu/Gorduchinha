package main

import (
	"os"

	"github.com/Nhanderu/gorduchinha/app"
	"github.com/Nhanderu/gorduchinha/cmd/api/server"
)

func main() {

	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	app := app.New(env)
	app.Logger.Infof("Running server at port %d.", app.Config.Server.Port)
	err := server.Run(
		app.Config.Server.Port,
		app.Config.Server.Prefix,
		app.Logger,
		app.Services().NewTeam(),
		app.Services().NewChamp(),
		app.Services().NewScraper(),
	)
	app.EndAsErr(err, "Could not run server.", app.Logger.InfoWriter(), app.Logger.ErrorWriter())
}
