package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	loggerLevelFile    = zap.ErrorLevel
	loggerLevelConsole = zap.FatalLevel
)

var logger *zap.Logger

func init() {

	logger = startLogger()
	loadEnv(logger)

}

func InitializeConfigs() *zap.Logger {
	return logger
}

func loadEnv(logger *zap.Logger) error {

	logger.Info("Getting the Enviroment definitions")

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	err := godotenv.Load(fmt.Sprintf("./app/config/.env.%s", env))

	if err != nil {
		logger.Error("Error in env params definitions, application should be stopped",
			zapcore.Field{Key: "error: ", Interface: err})
		panic(err)
	}

	return nil
}

func startLogger() *zap.Logger {

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	logFile, err := os.OpenFile("log_ "+time.Now().Format("01022006150405")+".json",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal("Error in initialize writter log local, application should be stopped")
		panic(err)
	}

	writer := zapcore.AddSync(logFile)
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, loggerLevelFile),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), loggerLevelConsole),
	)
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

}
