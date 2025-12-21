package log

import (
	"os"
	"strings"

	"github.com/YuanJun-93/CodeGenesis/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Init initializes a Zap logger based on the provided configuration.
func Init(c config.LogConf) *zap.Logger {
	// 1. Set Level
	var level zapcore.Level
	switch strings.ToLower(c.Level) {
	case "debug":
		level = zap.DebugLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	// 2. Encoder Configuration
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	var encoder zapcore.Encoder
	if strings.ToLower(c.Encoding) == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	// 3. Create Core (Write to Stdout for K8s/Docker)
	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(os.Stdout),
		level,
	)

	// 4. Build Logger
	logger := zap.New(core, zap.AddCaller())
	return logger
}
