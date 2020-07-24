package main

import (
	"os"

	"github.com/Nhanderu/gorduchinha/app"
)

var (
	AppVersion string
)

func main() {

	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	app := app.New(env, AppVersion)
	err := app.Services().NewScraper().ScrapeAndUpdate()
	app.EndAsErr(err, "Could not execute service.", app.Logger.InfoWriter(), app.Logger.ErrorWriter())
}
