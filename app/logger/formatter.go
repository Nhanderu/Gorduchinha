package logger

import (
	"strconv"
	"strings"

	"github.com/labstack/gommon/color"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	colorer = color.New()
)

// coloredJSONFormatter formats output in JSON with colored strings
type coloredJSONFormatter struct {
	logrus.JSONFormatter
	shouldColor bool
}

// Format formats the log output, coloring status and level fields if output is not in a file
func (formatter *coloredJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	if formatter.shouldColor {

		level := entry.Level.String()
		levelColor := strings.ToUpper(level)

		switch level {
		case logrus.DebugLevel.String():
			levelColor = colorer.Magenta(levelColor)
		case logrus.InfoLevel.String():
			levelColor = colorer.Blue(levelColor)
		case logrus.WarnLevel.String():
			levelColor = colorer.Yellow(levelColor)
		case logrus.ErrorLevel.String():
			levelColor = colorer.Red(levelColor)
		case logrus.FatalLevel.String():
			levelColor = colorer.Bold(colorer.Red(levelColor))
		}

		entry.Data["level"] = levelColor

		rawStatus, hasStatus := entry.Data["status"]
		if hasStatus {

			status := rawStatus.(int)

			statusColor := strconv.Itoa(status)
			switch {
			case status >= 500:
				statusColor = colorer.Red(status)
			case status >= 400:
				statusColor = colorer.Yellow(status)
			case status >= 300:
				statusColor = colorer.Cyan(status)
			default:
				statusColor = colorer.Green(status)
			}

			entry.Data["level"] = statusColor
		}
	}

	log, err := formatter.JSONFormatter.Format(entry)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return log, nil
}
