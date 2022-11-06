package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func configure() {
	cfg := getConfigFromEnv()
	//timeStampFormat := defaultTimeStampFormat
	level := log.DebugLevel

	if cfg != nil {
		if cfg.Level != nil {
			level = *cfg.Level
		}
		if cfg.TimeStampFormat != "" {
			//timeStampFormat = cfg.TimeStampFormat
		}
	}

	//formatter := log.JSONFormatter{
	//	TimestampFormat: timeStampFormat,
	//}
	//log.SetFormatter(&formatter)
	log.SetLevel(level)
	log.SetOutput(os.Stdout)
}

func GetLogger() *log.Logger {
	configure()
	return log.New()
}
