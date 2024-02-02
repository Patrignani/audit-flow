package models

import "github.com/Patrignani/audit-flow/src/config"

type LogLevel string

const (
	DebugLevel LogLevel = "debug"
	InfoLevel  LogLevel = "info"
	WarnLevel  LogLevel = "warn"
	ErrorLevel LogLevel = "error"
)

type Logger struct {
	LogLevel LogLevel
}

func GetLoggingConfig() *Logger {
	return &Logger{
		LogLevel: LogLevel(config.Env.LogLevel),
	}
}
