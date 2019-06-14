package logger

import (
	"io"
	"os"

	"github.com/Nhanderu/gorduchinha/src/infra/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Logger is the default application logger
type Logger struct {
	cfg   config.Config
	entry *logrus.Entry
}

// New returns a new Logger instance
func New(cfg config.Config) (Logger, error) {

	if cfg.Log.LogToFile {

		file, err := os.Create(cfg.Log.Path)
		if err != nil {
			return Logger{}, errors.WithStack(err)
		}

		logrus.SetOutput(file)
	}

	logrus.SetFormatter(&coloredJSONFormatter{cfg: cfg})
	logrus.SetLevel(logrus.InfoLevel)
	if cfg.App.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	hostname, _ := os.Hostname()
	entry := logrus.WithFields(logrus.Fields{
		"app-name":    cfg.App.Name,
		"app-version": cfg.App.Version,
		"hostname":    hostname,
	})

	return Logger{
		cfg:   cfg,
		entry: entry,
	}, nil
}

// WithFields returns a new logger with the given fields.
func (l Logger) WithFields(fields map[string]interface{}) Logger {
	return Logger{
		cfg:   l.cfg,
		entry: l.entry.WithFields(logrus.Fields(fields)),
	}
}

func (l Logger) Output() io.Writer {
	return l.entry.Writer()
}

func (l Logger) SetOutput(w io.Writer) {
	l.entry.Logger.SetOutput(w)
}
