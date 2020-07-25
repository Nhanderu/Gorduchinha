package logger

import (
	"io"
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	entry *logrus.Entry
}

func New(appName string, debug bool, filepath string) (Logger, error) {

	shouldColor := true
	if filepath != "" {

		shouldColor = false

		file, err := os.Create(filepath)
		if err != nil {
			return logrusLogger{}, errors.WithStack(err)
		}

		logrus.SetOutput(file)
	}

	formatter := new(coloredJSONFormatter)
	formatter.shouldColor = shouldColor
	logrus.SetFormatter(formatter)

	logrus.SetLevel(logrus.InfoLevel)
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	hostname, _ := os.Hostname()
	entry := logrus.WithFields(logrus.Fields{
		"app-name": appName,
		"hostname": hostname,
	})

	return logrusLogger{
		entry: entry,
	}, nil
}

// WithFields returns a new logger with the given fields.
func (l logrusLogger) WithFields(fields map[string]interface{}) Logger {
	return logrusLogger{
		entry: l.entry.WithFields(logrus.Fields(fields)),
	}
}

func (l logrusLogger) Output() io.Writer {
	return l.entry.Writer()
}

func (l logrusLogger) SetOutput(w io.Writer) {
	l.entry.Logger.SetOutput(w)
}

// Level return the logger level
func (l logrusLogger) Level() Level {
	return levelMap[l.entry.Level]
}

// SetLevel sets the logging level
func (l logrusLogger) SetLevel(level Level) {
	l.entry.Level = setLevelMap[level]
}

// InfoWriter returns the io.Writer for info level
func (l logrusLogger) InfoWriter() io.Writer {
	return l.entry.WriterLevel(logrus.InfoLevel)
}

// ErrorWriter returns the io.Writer for error level
func (l logrusLogger) ErrorWriter() io.Writer {
	return l.entry.WriterLevel(logrus.ErrorLevel)
}

// FatalWriter returns the io.Writer for fatal level
func (l logrusLogger) FatalWriter() io.Writer {
	return l.entry.WriterLevel(logrus.FatalLevel)
}

func (l logrusLogger) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

func (l logrusLogger) Info(args ...interface{}) {
	l.entry.Info(args...)
}

func (l logrusLogger) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

func (l logrusLogger) Error(args ...interface{}) {
	l.entry.Error(args...)
}

func (l logrusLogger) Fatal(args ...interface{}) {
	l.entry.Fatal(args...)
}

func (l logrusLogger) Panic(args ...interface{}) {
	l.entry.Panic(args...)
}

func (l logrusLogger) Debugf(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

func (l logrusLogger) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

func (l logrusLogger) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

func (l logrusLogger) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

func (l logrusLogger) Fatalf(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

func (l logrusLogger) Panicf(format string, args ...interface{}) {
	l.entry.Panicf(format, args...)
}

func (l logrusLogger) Debugln(args ...interface{}) {
	l.entry.Debugln(args...)
}

func (l logrusLogger) Infoln(args ...interface{}) {
	l.entry.Infoln(args...)
}

func (l logrusLogger) Warnln(args ...interface{}) {
	l.entry.Warnln(args...)
}

func (l logrusLogger) Errorln(args ...interface{}) {
	l.entry.Errorln(args...)
}

func (l logrusLogger) Fatalln(args ...interface{}) {
	l.entry.Fatalln(args...)
}

func (l logrusLogger) Panicln(args ...interface{}) {
	l.entry.Panicln(args...)
}
