package log

import (
	"strings"

	"github.com/676767ap/otus-go-hw/project/internal/config"

	"go.uber.org/zap"
)

func Init(cfg *config.Config) {
	if cfg.DevMode {
		ReInitLogger(true, "debug")
	} else {
		ReInitLogger(false, "warn")
	}
}

func ReInitLogger(dev bool, logLevel string) {
	initLogger(dev, strings.ToLower(logLevel))
}

func Debug(args ...interface{}) {
	zap.S().Debug(args...)
}

func Info(args ...interface{}) {
	zap.S().Info(args...)
}

func Warn(args ...interface{}) {
	zap.S().Warn(args...)
}

func Error(args ...interface{}) {
	zap.S().Error(args...)
}

func Fatal(args ...interface{}) {
	zap.S().Fatal(args...)
}

func Panic(args ...interface{}) {
	zap.S().Panic(args...)
}

func Debugf(format string, args ...interface{}) {
	zap.S().Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	zap.S().Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	zap.S().Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	zap.S().Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	zap.S().Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	zap.S().Panicf(format, args...)
}
