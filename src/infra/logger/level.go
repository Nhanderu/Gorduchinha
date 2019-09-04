package logger

import (
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
