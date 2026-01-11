package common

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.SugaredLogger

func InitLogger() *zap.SugaredLogger {
	var logger *zap.Logger
	var err error

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(parseLogLevel(logLevel))
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err = config.Build()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	globalLogger = logger.Sugar()
	return globalLogger
}

func GetLogger() *zap.SugaredLogger {
	if globalLogger == nil {
		return InitLogger()
	}
	return globalLogger
}

func parseLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
