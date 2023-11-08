package config

type Logger interface {
	Info(message string)
	Error(message string)
}
