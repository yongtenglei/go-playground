package logger

import (
	"my_bubble/configs"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Init(config *configs.LogConf, mode string) (err error) {
	writeSyncer := getLogwriter(
		config.Filename,
		config.MaxSize,
		config.MaxBackups,
		config.MaxAge,
	)

	// 自定义 log level
	encoder := getEncoder()
	var lvl = new(zapcore.Level)
	if err = lvl.UnmarshalText([]byte(config.Level)); err != nil {
		return
	}

	var core zapcore.Core
	if mode == "release" {
		// release mode 下 日志输出至文件
		core = zapcore.NewCore(encoder, writeSyncer, lvl)
	} else {
		// 其他模式下, 日志输出至文件与终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, lvl),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	}

	// Enable caller
	lg := zap.New(core, zap.AddCaller())

	zap.ReplaceGlobals(lg)

	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogwriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
