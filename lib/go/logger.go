package logger

import (
	"github.com/sirupsen/logrus"
)

func Info(s string) {
	logrus.Info(s)
}
