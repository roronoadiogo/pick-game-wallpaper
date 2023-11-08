package config

import "go.uber.org/zap"

type AppLogger struct {
	Logger *zap.Logger
}

func (al *AppLogger) Info(message string) {
	al.Logger.Info(message)
}

func (al *AppLogger) Error(message string) {
	al.Logger.Error(message)
}
