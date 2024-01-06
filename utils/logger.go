// utils/logger.go
package utils

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger struct
type Logger struct {
	logger *zap.Logger
}

// SetupLogger initializes the logger
func SetupLogger() *Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := config.Build()

	return &Logger{
		logger: logger,
	}
}

// LogError logs an error message
func (l *Logger) LogError(message string, err error) {
	l.logger.Error(message, zap.Error(err))
}

// LogRequest logs information about the request
func (l *Logger) LogRequest(requestID, endpoint string) {
	l.logger.Info("Request",
		zap.String("RequestID", requestID),
		zap.String("Endpoint", endpoint),
	)
}

// GenerateUUID generates a unique identifier
func GenerateUUID() string {
	return uuid.New().String()
}
