package logger

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	PanicLevel = uint8(logrus.PanicLevel)
	FatalLevel = uint8(logrus.FatalLevel)
	ErrorLevel = uint8(logrus.ErrorLevel)
	WarnLevel  = uint8(logrus.WarnLevel)
	InfoLevel  = uint8(logrus.InfoLevel)
	DebugLevel = uint8(logrus.DebugLevel)
)

// 创建日志记录器
func CreateLoggerOnce(level, filelevel uint8) {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	base := filepath.Base(os.Args[0])
	name := strings.TrimSuffix(base, filepath.Ext(base))
	path := "logs/" + name + "/" + tm.Format("20060102150405") + "/"

	once.Do(func() {
		globalLogger = &logger{
			file:      newLoggerOfLFShook(1048576000, 100, 365, path),
			console:   newLoggerOfConsole(),
			fileLevel: logrus.Level(filelevel),
		}
		globalLogger.console.SetLevel(logrus.Level(level))
		globalLogger.file.SetLevel(logrus.Level(filelevel))
	})
}

// 输出Debug日志
func Debug(v ...interface{}) {
	if globalLogger != nil {
		globalLogger.available(logrus.DebugLevel).Debug(v)
	}
}

// 格式化输出Debug日志
func Debugf(format string, params ...interface{}) {
	if globalLogger != nil {
		globalLogger.available(logrus.DebugLevel).Debugf(format, params...)
	}
}

// 输出Info日志
func Info(v ...interface{}) {
	if globalLogger != nil {
		globalLogger.available(logrus.InfoLevel).Info(v)
	}
}

// 格式化输出Info日志
func Infof(format string, params ...interface{}) {
	if globalLogger != nil {
		globalLogger.available(logrus.InfoLevel).Infof(format, params...)
	}
}

// 输出Warn日志
func Warn(v ...interface{}) {
	if globalLogger != nil {
		globalLogger.available(logrus.WarnLevel).Warn(v)
	}
}

// 格式化输出Warn日志
func Warnf(format string, params ...interface{}) {
	if globalLogger != nil {
		globalLogger.available(logrus.WarnLevel).Warnf(format, params...)
	}
}

// 输出Error日志
func Error(v ...interface{}) {
	if globalLogger != nil {
		globalLogger.available(logrus.ErrorLevel).Error(v)
	}
}

// 格式化输出Error日志
func Errorf(format string, params ...interface{}) {
	if globalLogger != nil {
		globalLogger.available(logrus.ErrorLevel).Errorf(format, params...)
	}
}

// 输出Fatal日志
func Fatal(v ...interface{}) {
	if globalLogger != nil {
		globalLogger.available(logrus.FatalLevel).Fatal(v)
	}
}

// 格式化输出Fatal日志
func Fatalf(format string, params ...interface{}) {
	if globalLogger != nil {
		globalLogger.available(logrus.FatalLevel).Fatalf(format, params...)
	}
}

// 输出Panic日志
func Panic(v ...interface{}) {
	if globalLogger != nil {
		globalLogger.available(logrus.PanicLevel).Panic(v)
	}
}

// 格式化输出Panic日志
func Panicf(format string, params ...interface{}) {
	if globalLogger != nil {
		globalLogger.available(logrus.PanicLevel).Panicf(format, params...)
	}
}

// 日志选项
type logger struct {
	file      *logrus.Logger // 文件日志
	console   *logrus.Logger // 控制台日志
	fileLevel logrus.Level   // 文件日志级别
}

// 获取可用的日志记录器
func (lg *logger) available(level logrus.Level) *logrus.Logger {
	if level <= lg.fileLevel {
		return lg.file
	}
	return lg.console
}

var once sync.Once
var globalLogger *logger

// 创建终端记录器
func newLoggerOfConsole() *logrus.Logger {
	lg := logrus.New()
	for _, level := range logrus.AllLevels {
		lg.Level |= level
	}
	lg.Formatter = &logrus.JSONFormatter{}
	return lg
}

// 创建文件记录器
func newLoggerOfLFShook(maxsize int, maxbackup int, maxage int, path string) *logrus.Logger {
	lg := logrus.New()
	writerMap := lfshook.WriterMap{}
	for _, level := range logrus.AllLevels {
		lg.Level |= level
		writer := &lumberjack.Logger{
			Filename:   path + level.String() + ".log",
			MaxSize:    maxsize,
			MaxBackups: maxbackup,
			MaxAge:     maxage,
		}
		writerMap[level] = writer
	}
	lg.Formatter = &logrus.JSONFormatter{}
	lg.Hooks.Add(lfshook.NewHook(writerMap, nil))
	return lg
}
