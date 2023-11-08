package config

import (
	"go.uber.org/zap"
)

type AppLogger struct {
	Logger *zap.Logger
}

func (al *AppLogger) Info(message string) {
	al.Logger.Info(message)
}

func (al *AppLogger) Error(message string, err error) {
	al.Logger.Error(message, zap.Error(err))
}
