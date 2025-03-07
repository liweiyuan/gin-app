package log

import (
	"os"

	"gin-app/config"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger = logrus.New()

func init() {
	// 设置日志级别
	level, err := logrus.ParseLevel(config.GlobalConfig.Log.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	Logger.SetLevel(level)

	// 设置日志格式
	switch config.GlobalConfig.Log.Format {
	case "json":
		Logger.SetFormatter(&logrus.JSONFormatter{})
	default:
		Logger.SetFormatter(&logrus.TextFormatter{})
	}

	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		logrus.Fatalf("Could not create log directory: %v", err)
	}

	switch config.GlobalConfig.Log.Output {
	case "stdout":
		Logger.Out = os.Stdout
	case "file":
		Logger.Out = &lumberjack.Logger{
			Filename:   "logs/" + config.GlobalConfig.Log.Filename,
			MaxSize:    config.GlobalConfig.Log.MaxSize,
			MaxBackups: config.GlobalConfig.Log.MaxBackups,
			MaxAge:     config.GlobalConfig.Log.MaxAge,
			Compress:   config.GlobalConfig.Log.Compress,
		}
	default:
		Logger.Out = os.Stdout
	}
}
