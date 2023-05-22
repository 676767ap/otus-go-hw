package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initLogger(dev bool, logLevel string) {
	var config zap.Config

	if dev {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config = zap.NewProductionConfig()
	}

	config.Level = zap.NewAtomicLevelAt(zapLogLevel(logLevel))
	config.OutputPaths = []string{"stdout"}

	callerLevelsToSkip := 1

	logger, err := config.Build(
		zap.AddCallerSkip(callerLevelsToSkip),
		zap.AddStacktrace(zapLogLevel("error")),
	)
	if err != nil {
		panic(fmt.Errorf("building new logger failed, error %w", err))
	}

	zap.ReplaceGlobals(logger)
}

func zapLogLevel(logLevel string) zapcore.Level {
	switch logLevel {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.ErrorLevel
	}
}
