package log

import "github.com/NooBeeID/go-logging/logger"

var Log logger.Logger

func init() {
	Log = logger.NewLog()
	Log.SetReportCaller(true)
}
