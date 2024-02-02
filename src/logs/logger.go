package logs

import (
	"fmt"
	"sync"

	"github.com/Patrignani/audit-flow/src/models"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var onceLogger = new(sync.Once)

const (
	RequestIDLogFieldKey string = "request_id"
	GloboID              string = "globo_id"
)

func NewLogger(logConfig *models.Logger) *zap.Logger {
	onceLogger.Do(func() {
		logger, err := configureLogger(logConfig)
		if err != nil {
			panic(fmt.Errorf("failed to initialize logger: %s", err.Error()))
		}

		zap.ReplaceGlobals(logger)
	})

	return zap.L()
}

func configureLogger(logConfig *models.Logger) (*zap.Logger, error) {
	var zapConfig zap.Config

	env := models.GetEnvironmentConfig()

	logLevel := zapLogLevel(logConfig.LogLevel)

	encodingConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "msg",
		CallerKey:      "caller",
		StacktraceKey:  "stackTrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}

	zapConfig = zap.Config{
		Level:            zap.NewAtomicLevelAt(logLevel),
		Encoding:         "json",
		EncoderConfig:    encodingConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	if env == models.ProductionEnv {
		zapConfig.Sampling = &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		}
	}

	if env == models.LocalEnv {
		zapConfig.Encoding = "console"
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func zapLogLevel(level models.LogLevel) zapcore.Level {
	switch level {
	case models.DebugLevel:
		return zapcore.DebugLevel

	case models.InfoLevel:
		return zapcore.InfoLevel

	case models.WarnLevel:
		return zapcore.WarnLevel

	default:
		return zapcore.ErrorLevel
	}
}
