package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/Nhanderu/gorduchinha/src/domain/service"
	"github.com/Nhanderu/gorduchinha/src/infra/cache"
	"github.com/Nhanderu/gorduchinha/src/infra/config"
	"github.com/Nhanderu/gorduchinha/src/infra/logger"
	"github.com/Nhanderu/gorduchinha/src/data"
	"github.com/Nhanderu/gorduchinha/src/server"
)

var (
	AppVersion string
)

func main() {

	serverPort := os.Getenv("PORT")
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	// Configuration.
	cfg, err := config.Read(env)
	cfg.App.Version = AppVersion
	cfg.Server.Port, _ = strconv.Atoi(serverPort)
	endAsErr(err, "Could not read configuration file.", os.Stdout, os.Stderr)

	// Logging structure.
	log, err := logger.New(
		cfg.App.Name,
		cfg.App.Version,
		cfg.App.Debug,
		cfg.Log.Path,
	)
	endAsErr(err, "Could not create logging structure.", os.Stdout, os.Stderr)

	// Data manager.
	log.Infof("Connecting to the database at %s:%d.", cfg.DB.Host, cfg.DB.Port)
	db, err := data.Connect(
		cfg.DB.User,
		cfg.DB.Pass,
		cfg.DB.Name,
		cfg.DB.Host,
		cfg.DB.Port,
	)
	endAsErr(err, "Could not connect to database.", log.InfoWriter(), log.ErrorWriter())
	atInterruption(func() {
		log.Infof("Closing DB Connection.")
		db.Close()
	})

	// Cache manager.
	log.Infof("Connecting to the cache server at %s:%d.", cfg.Cache.Host, cfg.Cache.Port)
	cache := cache.New(
		cfg.Cache.Host,
		cfg.Cache.Port,
		cfg.Cache.DB,
		cfg.Cache.Pass,
		cfg.Cache.Prefix,
		cfg.Cache.DefaultExpiration,
	)
	cache.CleanAll()

	// HTTP client.
	httpClient := &http.Client{Timeout: cfg.HTTPClient.Timeout}

	// Services;
	teamService := service.NewTeamService(db, cache)
	champService := service.NewChampService(db, cache)
	scraperService := service.NewScraperService(db, log, httpClient, teamService, champService)
	server.Run(cfg.App.Version, cfg.Server.Port, cfg.Server.Prefix, log, teamService, champService, scraperService)
}

func atInterruption(fn func()) {
	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, os.Interrupt)
		<-sc

		fn()
		os.Exit(0)
	}()
}

func endAsErr(err error, message string, infow io.Writer, errorw io.Writer) {
	if err != nil {
		fmt.Fprintln(errorw, "Error:", err)
		fmt.Fprintln(infow, message)
		time.Sleep(time.Millisecond * 50) // needed for printing all messages before exiting
		os.Exit(1)
	}
}
