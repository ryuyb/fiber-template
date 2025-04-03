package middleware

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"live-pilot/pkg/conf"
	"os"
	"strings"
)

func NewLogger(appConfig conf.AppConfig) *fiberzap.LoggerConfig {
	zapLevel := getLogLevel(appConfig.Log.Level)

	encoderConfig := getEncoderConfig(appConfig)
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	coreConfigs := []fiberzap.CoreConfig{
		{
			Encoder:      zapcore.NewConsoleEncoder(encoderConfig),
			WriteSyncer:  zapcore.AddSync(os.Stdout),
			LevelEncoder: zap.NewAtomicLevelAt(zapLevel),
		},
	}
	if appConfig.Log.File.Enable {
		fileConfig := fiberzap.CoreConfig{
			Encoder:      zapcore.NewJSONEncoder(getEncoderConfig(appConfig)),
			WriteSyncer:  zapcore.AddSync(buildLumberjackLogger(appConfig.Log.File)),
			LevelEncoder: zap.NewAtomicLevelAt(zapLevel),
		}
		coreConfigs = append(coreConfigs, fileConfig)
	}
	logger := fiberzap.NewLogger(fiberzap.LoggerConfig{
		CoreConfigs: coreConfigs,
		ZapOptions: []zap.Option{
			zap.AddCaller(),
			zap.AddCallerSkip(3),
			zap.AddStacktrace(zapcore.ErrorLevel),
		},
	})

	return logger
}

func getLogLevel(level string) zapcore.Level {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func getEncoderConfig(appConfig conf.AppConfig) zapcore.EncoderConfig {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	if appConfig.IsProdEnv() {
		encoderConfig = zap.NewProductionEncoderConfig()
	}
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return encoderConfig
}

func buildLumberjackLogger(fileCfg conf.FileLog) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   fileCfg.Path,
		MaxSize:    fileCfg.MaxSize,
		MaxBackups: fileCfg.MaxBackups,
		MaxAge:     fileCfg.MaxAge,
		Compress:   fileCfg.Compress,
	}
}
