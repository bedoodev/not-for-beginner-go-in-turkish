package main

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Println("Error in initializing Zap logger.")
	}
	defer logger.Sync()

	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // Kendi formatın
	cfg.DisableStacktrace = true
	logger, err = cfg.Build()
	if err != nil {
		log.Println("Error in building Zap logger.")
	}

	logger.Info("User logged in.", zap.String("username", "Bedirhan"), zap.String("method", "GET"))
	logger.Error("Error in processing request.", zap.String("error", "Database connection failed."))
	logger.Warn("Slow query detected.", zap.Duration("duration", 1_000_000_000))
}
