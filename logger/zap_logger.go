package logger

import (
	"myeduate/utils"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var sugarLogger *zap.SugaredLogger

func InitLogger(config utils.Config) *zap.SugaredLogger {
	writeSyncer := getLogWriter(config)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
	sugarLogger = logger.Sugar()
	return sugarLogger
}

func GetLogger() *zap.SugaredLogger {
	return sugarLogger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(config utils.Config) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.LogPath + config.LogFileName,
		MaxSize:    config.LogFileMaxSize,
		MaxBackups: config.LogFileMaxBackups,
		MaxAge:     config.LogFileMaxAge,
		Compress:   config.LogFileCompression,
	}
	return zapcore.AddSync(lumberJackLogger)
}
