package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.Out = os.Stdout
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000000000",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "ts",
			logrus.FieldKeyLevel: "lv",
			logrus.FieldKeyMsg:   "msg",
			logrus.FieldKeyFunc:  "func",
		},
		DisableColors: true,
	})
}

func Get(name string) *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"module": name,
	})
}

func Debug(args ...interface{}) {
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

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}