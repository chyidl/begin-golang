package ch35

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

var _isDebug = true

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.Out = os.Stdout
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.00000000",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "ts",
			logrus.FieldKeyLevel: "lv",
			logrus.FieldKeyMsg: "msg",
			logrus.FieldKeyFunc: "func",
		},
		DisableColors: false,
	})
}

func Get(name string) *logrus.Entry {
	return logger.WithFields(logrus.Fields{"module": name})
}

func TestLogrus(t *testing.T) {
	logger := Get("TestLogrus")
	logger.Info("Say Hello", 1, 2)
}