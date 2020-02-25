package log

import (
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
	formatter := &logrus.TextFormatter{
		TimestampFormat: "02-01-2006 15:04:05", // the "time" field configuratiom
		FullTimestamp:   true,
	}
	logrus.SetFormatter(formatter)
}
