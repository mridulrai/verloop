package logger

import (
	"go.uber.org/zap"
)

// InitSugarLogger returns a Zap SugaredLogger.
func InitSugarLogger() *zap.SugaredLogger {
	var logger *zap.Logger
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Cannot initialize zap logger")
	}
	return logger.Sugar()
}
