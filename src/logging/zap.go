package logging

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strconv"
	"time"
)

func NewLogger() *zap.Logger {
	year, month, day := time.Now().Date()
	_ = os.Mkdir("logs", 0755)
	filename := "/logs/" + strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day) + ".log"
	_, _ = os.Create(filename)

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.OutputPaths = []string{
		"stdout",
		"." + filename,
	}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := config.Build(zap.AddCaller())
	if err != nil {
		fmt.Println("error creating zap logger - returning NewNop: ", err)
		return zap.NewNop()
	}

	return logger
}
