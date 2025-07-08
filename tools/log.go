package tools

import (
	"os"

	"github.com/sirupsen/logrus"
)

/**
@description
@date: 03/13 14:18
@author Gk
**/

var (
	LogrusLogger = logrus.New()
	LogFile      = "tasks.log"
)

func init() {
	formatter := new(logrus.TextFormatter)
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"
	formatter.DisableQuote = true
	LogrusLogger.Formatter = formatter
	LogrusLogger.SetOutput(os.Stdout)
}
