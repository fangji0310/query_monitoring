package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(environment string) *zap.Logger {
	config := zap.NewProductionConfig()
	config.DisableCaller = false
	config.DisableStacktrace = false
	config.Sampling = nil
	//config.Encoding = "console"
	if environment != "production" {
		config.Level.SetLevel(zapcore.DebugLevel)
	}
	return zap.Must(config.Build())
}
