package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/Nhanderu/gorduchinha/src/domain/service"
	"github.com/Nhanderu/gorduchinha/src/infra/config"
	"github.com/Nhanderu/gorduchinha/src/infra/logger"
	"github.com/Nhanderu/gorduchinha/src/integration/cache"
	"github.com/Nhanderu/gorduchinha/src/integration/data"
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
	log, err := logger.New(cfg)
	endAsErr(err, "Could not create logging structure.", os.Stdout, os.Stderr)

	// Data manager.
	log.Infof("Connecting to the database at %s:%d.", cfg.DB.Host, cfg.DB.Port)
	db, err := data.Connect(cfg)
	endAsErr(err, "Could not connect to database.", log.InfoWriter(), log.ErrorWriter())
	atInterruption(func() {
		log.Infof("Closing DB Connection.")
		db.Close()
	})

	// Cache manager.
	log.Infof("Connecting to the cache server at %s:%d.", cfg.Cache.Host, cfg.Cache.Port)
	cache := cache.New(cfg)
	cache.CleanAll()

	log.Infof("Starting service.")
	svc := service.New(db, cache, cfg, log, nil, nil, nil)
	server.Run(svc, cfg, log)
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
