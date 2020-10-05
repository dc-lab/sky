package logging

import (
	"fmt"
	"path"
	"runtime"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
)

type Config struct {
	LogFile  string
	LogLevel string
}

func InitDefaultLogging() error {
	return InitLogging(Config{
		LogFile:  "",
		LogLevel: "info",
	})
}

func InitLogging(config Config) error {
	logLevel := log.InfoLevel
	if config.LogLevel != "" {
		var err error
		logLevel, err = log.ParseLevel(config.LogLevel)
		if err != nil {
			return fmt.Errorf("failed to parse log level: %w", err)
		}
	}

	callerPrettyfier := func(f *runtime.Frame) (string, string) {
		filename := path.Base(f.File)
		return "", fmt.Sprintf("%s:%d", filename, f.Line)
	}

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:    true,
		CallerPrettyfier: callerPrettyfier,
		TimestampFormat:  time.RFC3339,
	})
	log.SetReportCaller(true)
	log.SetLevel(logLevel)

	if config.LogFile != "" {
		log.WithField("file", config.LogFile).Info("Initializing log file")
		hook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
			Filename: config.LogFile,
			Formatter: &log.JSONFormatter{
				CallerPrettyfier: callerPrettyfier,
				TimestampFormat:  time.RFC3339,
			},
			MaxSize:    50, // megabytes
			MaxBackups: 3,
			MaxAge:     365, // days
			Level:      logLevel,
		})
		if err != nil {
			return fmt.Errorf("failed to open log file: %w", err)
		}
		log.AddHook(hook)
	}

	return nil
}
