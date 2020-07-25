package app

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"

	"github.com/Nhanderu/gorduchinha/app/cache"
	"github.com/Nhanderu/gorduchinha/app/config"
	"github.com/Nhanderu/gorduchinha/app/contract"
	"github.com/Nhanderu/gorduchinha/app/data"
	"github.com/Nhanderu/gorduchinha/app/logger"
)

type App struct {
	Config       config.Config
	Logger       logger.Logger
	DataManager  contract.DataManager
	CacheManager contract.CacheManager
	HTTPClient   *http.Client
}

func New(env string) App {

	cfg, err := config.Read(env)
	endAsErr(err, "Could not read configuration file.", os.Stdout, os.Stderr)

	log, err := logger.New(
		cfg.App.Name,
		cfg.App.Debug,
		cfg.Log.Path,
	)
	endAsErr(err, "Could not create logging structure.", os.Stdout, os.Stderr)

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

	log.Infof("Connecting to the cache server at %s:%d.", cfg.Cache.Host, cfg.Cache.Port)
	cache := cache.New(
		cfg.Cache.Host,
		cfg.Cache.Port,
		cfg.Cache.DB,
		cfg.Cache.Pass,
		cfg.Cache.Prefix,
		cfg.Cache.DefaultExpiration,
	)

	httpClient := &http.Client{Timeout: cfg.HTTPClient.Timeout}

	return App{
		Config:       cfg,
		Logger:       log,
		DataManager:  db,
		CacheManager: cache,
		HTTPClient:   httpClient,
	}

}

func (app App) AtInterruption(fn func()) {
	atInterruption(fn)
}

func (app App) EndAsErr(err error, message string, infow io.Writer, errorw io.Writer) {
	endAsErr(err, message, infow, errorw)
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
		os.Exit(1)
	}
}
