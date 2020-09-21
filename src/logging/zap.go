package logging

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := config.Build(zap.AddCaller())
	if err != nil {
		fmt.Println("error creating zap logger - returning NewNop")
		return zap.NewNop()
	}

	return logger
}
