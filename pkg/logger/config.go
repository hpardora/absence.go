package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

const (
	timeStampKey = "LOG_TS_FORMAT"
	levelKey     = "LOG_LEVEL"
)

const defaultTimeStampFormat = "2006-01-02 15:04:05.000"

type config struct {
	TimeStampFormat string
	Level           *logrus.Level
}

func getConfigFromEnv() *config {
	tsFormat, ok := os.LookupEnv(timeStampKey)
	if !ok {
		tsFormat = defaultTimeStampFormat
	}

	lvlString, ok := os.LookupEnv(levelKey)
	if !ok {
		lvlString = "debug"
	}
	lvlParsed, err := logrus.ParseLevel(lvlString)
	if err != nil {
		lvlParsed = logrus.DebugLevel
	}

	return &config{
		TimeStampFormat: tsFormat,
		Level:           &lvlParsed,
	}

}
