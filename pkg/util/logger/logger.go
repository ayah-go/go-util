package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.SugaredLogger

func InitLogger(baseName string) {
	var allCore []zapcore.Core

	writeSyncer := getLogWriter(baseName)
	encoder := getEncoder()

	//	打印到控制台
	allCore = append(allCore, zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel))
	//	输出到文件
	allCore = append(allCore, zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel))

	zapLog := zap.New(zapcore.NewTee(allCore...), zap.AddCaller())
	logger = zapLog.Sugar()
}

func L() *zap.SugaredLogger {
	return logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:   "time",
		LevelKey:  "level",
		NameKey:   "logger",
		CallerKey: "caller",
		// FunctionKey:   zapcore.OmitKey,
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder, // 指定颜色
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		// EncodeTime:     zapcore.ISO8601TimeEncoder,  // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.StringDurationEncoder,
		// EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
		EncodeCaller: zapcore.ShortCallerEncoder, // 路径编码器
	}

	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(baseName string) zapcore.WriteSyncer {
	filePath := "./logs/" + baseName + ".log"
	// if !exists(filePath) {
	// 	file, _ := os.Create(filePath)
	// 	defer file.Close()
	// }

	lumberJackLogger := &lumberjack.Logger{
		Filename:   filePath, // 日志文件的位置
		MaxSize:    1,        // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 5,        // 保留旧文件的最大个数
		MaxAge:     30,       // 保留旧文件的最大天数
		Compress:   false,    // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

/*
查看文件/文件夹是否存在
*/
func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//func Debug(msg ...interface{}) {
//	Logger.Debug(msg...)
//}

// // func Debugf(template string, args ...interface{}) {
// // 	sugarLogger.Debugf(template, args...)
// // }
//
//func Info(args ...interface{}) {
//	logger.L().Info(args...)
//}

//
// func Warn(args ...interface{}) {
// 	sugarLogger.Warn(args...)
// }
//
//func Error(args ...interface{}) {
//	Logger.Error(args...)
//}

//
// func DPanic(args ...interface{}) {
// 	sugarLogger.DPanic(args...)
// }
//
// func Panic(args ...interface{}) {
// 	sugarLogger.Panic(args...)
// }
//
// func Fatal(args ...interface{}) {
// 	sugarLogger.Fatal(args...)
// }
