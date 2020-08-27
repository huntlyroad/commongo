// Package logger is where we explain the package.
// Some other stuff.
package logger

import (
	"fmt"
	"os"
	"strings"

	LOG "github.com/alecthomas/log4go"
)

func parseLogLevel(level string) LOG.Level {
	switch strings.ToLower(level) {
	case "debug":
		return LOG.DEBUG
	case "info":
		return LOG.INFO
	case "warning":
		return LOG.WARNING
	case "error":
		return LOG.ERROR
	default:
		return LOG.DEBUG
	}
}

// InitialLogger setup logger with log level etc.
func InitialLogger(logDir, logFile, level string, logFlush bool, verbose bool) error {
	logLevel := parseLogLevel(level)

	if verbose {
		LOG.AddFilter("stdout", logLevel, LOG.NewConsoleLogWriter())
	}

	if len(logDir) == 0 {
		logDir = "logs"
	}
	// check directory exists
	if _, err := os.Stat(logDir); err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(logDir, os.ModeDir|os.ModePerm); err != nil {
			return fmt.Errorf("create log.dir[%v] failed[%v]", logDir, err)
		}
	}

	if len(logFile) != 0 {
		if !logFlush {
			LOG.LogBufferLength = 32
		} else {
			LOG.LogBufferLength = 0
		}
		fileLogger := LOG.NewFileLogWriter(fmt.Sprintf("%s/%s", logDir, logFile), true)
		fileLogger.SetRotateDaily(true)
		fileLogger.SetFormat("[%D %T] [%L] [%s] %M")
		fileLogger.SetRotateMaxBackup(7)
		LOG.AddFilter("file", logLevel, fileLogger)
	} else {
		return fmt.Errorf("log.file[%v] shouldn't be empty", logFile)
	}

	return nil
}
