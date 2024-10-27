package logger

import (
	"backend/config"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init() (err error) {
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(config.Conf.LogConfig.Level))
	if err != nil {
		return
	}
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	var core zapcore.Core
	if config.Conf.LogConfig.Mode == "dev" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else if config.Conf.LogConfig.Mode == "release" {
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}
	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
	return
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.Conf.LogConfig.Filename,
		MaxSize:    config.Conf.LogConfig.MaxSize,
		MaxAge:     config.Conf.LogConfig.MaxAge,
		MaxBackups: config.Conf.LogConfig.MaxBackups,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	return zapcore.NewConsoleEncoder(encoderConfig)
}
