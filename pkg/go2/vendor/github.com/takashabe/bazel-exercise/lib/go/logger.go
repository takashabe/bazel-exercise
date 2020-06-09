package logger

import (
	"github.com/sirupsen/logrus"
)

func Info(format, s string) {
	logrus.Infof(format, s)
}
