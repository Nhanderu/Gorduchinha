package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

// Level represents a logging level
type Level uint8

// Level standard values
const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

var (
	levelMap = map[logrus.Level]Level{
		logrus.DebugLevel: LevelDebug,
		logrus.InfoLevel:  LevelInfo,
		logrus.WarnLevel:  LevelWarn,
		logrus.ErrorLevel: LevelError,
		logrus.FatalLevel: LevelFatal,
		logrus.PanicLevel: LevelPanic,
	}
	setLevelMap = map[Level]logrus.Level{
		LevelDebug: logrus.DebugLevel,
		LevelInfo:  logrus.InfoLevel,
		LevelWarn:  logrus.WarnLevel,
		LevelError: logrus.ErrorLevel,
		LevelFatal: logrus.FatalLevel,
		LevelPanic: logrus.PanicLevel,
	}
)

// Level return the logger level
func (l Logger) Level() Level {
	return levelMap[l.entry.Level]
}

// SetLevel sets the logging level
func (l Logger) SetLevel(level Level) {
	l.entry.Level = setLevelMap[level]
}

// InfoWriter returns the io.Writer for info level
func (l Logger) InfoWriter() io.Writer {
	return l.entry.WriterLevel(logrus.InfoLevel)
}

// ErrorWriter returns the io.Writer for error level
func (l Logger) ErrorWriter() io.Writer {
	return l.entry.WriterLevel(logrus.ErrorLevel)
}

// FatalWriter returns the io.Writer for fatal level
func (l Logger) FatalWriter() io.Writer {
	return l.entry.WriterLevel(logrus.FatalLevel)
}

func (l Logger) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

func (l Logger) Info(args ...interface{}) {
	l.entry.Info(args...)
}

func (l Logger) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

func (l Logger) Error(args ...interface{}) {
	l.entry.Error(args...)
}

func (l Logger) Fatal(args ...interface{}) {
	l.entry.Fatal(args...)
}

func (l Logger) Panic(args ...interface{}) {
	l.entry.Panic(args...)
}

func (l Logger) Debugf(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

func (l Logger) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

func (l Logger) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

func (l Logger) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

func (l Logger) Fatalf(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

func (l Logger) Panicf(format string, args ...interface{}) {
	l.entry.Panicf(format, args...)
}

func (l Logger) Debugln(args ...interface{}) {
	l.entry.Debugln(args...)
}

func (l Logger) Infoln(args ...interface{}) {
	l.entry.Infoln(args...)
}

func (l Logger) Warnln(args ...interface{}) {
	l.entry.Warnln(args...)
}

func (l Logger) Errorln(args ...interface{}) {
	l.entry.Errorln(args...)
}

func (l Logger) Fatalln(args ...interface{}) {
	l.entry.Fatalln(args...)
}

func (l Logger) Panicln(args ...interface{}) {
	l.entry.Panicln(args...)
}
