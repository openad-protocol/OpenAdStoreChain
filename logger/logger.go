package logger

import (
	"AdServerCollector/conf"
	"context"

	"github.com/zxysilent/logs"
)

var logger *logs.Logger

//logger.Debug("你好啊")
//logger.Info("你好啊")
//logger.Warn("你好啊")
//loggerError("你好啊")

func init() {
	logger = logs.New(nil)
	logger.SetSkip(1)
	if conf.ISTEST {
		logger.SetFile("./logs/debug.log")
		ctx := logs.TraceCtx(context.Background())
		logger.Ctx(ctx)
		logger.SetLevel(logs.LDEBUG)
		logger.SetCaller(true)
		logger.SetMaxSize(2048)
		// 设置同时显示到控制台
		// 默认只输出到文件
		logger.SetCons(true)
	} else {
		logger.SetFile("./logs/debug.log")
		ctx := logs.TraceCtx(context.Background())
		logger.Ctx(ctx)
		logger.SetMaxSize(2048)
		logger.SetLevel(logs.LDEBUG)
		logger.SetCaller(true)
		// 设置同时显示到控制台
		// 默认只输出到文件
		logger.SetCons(true)
	}
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debug(args...)
}
func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}
