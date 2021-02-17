package utils

import "github.com/sirupsen/logrus"

func NewLogger() *logrus.Logger {
	var log = logrus.New()
	return log
}
