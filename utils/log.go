package utils

import (
	"fmt"
	"os"
	"runtime"

	"github.com/Sirupsen/logrus"
)

var log *logrus.Logger

func Logger() *logrus.Logger {
	return log
}

func init() {
	log = logrus.New()
	//log.Formatter = new(logrus.JSONFormatter)
	log.Formatter = new(logrus.TextFormatter)
	//log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true
	log.Level = logrus.DebugLevel
	log.Out = os.Stdout
}

func Recover() {
	if r := recover(); r != nil {
		var err error
		switch entry := r.(type) {
		case error:
			err = entry
		case *logrus.Entry:
			logrus.WithFields(logrus.Fields{
				"omg":         true,
				"err_animal":  entry.Data["animal"],
				"err_size":    entry.Data["size"],
				"err_level":   entry.Level,
				"err_message": entry.Message,
				"number":      100,
			}).Error("The ice breaks!")
		default:
			err = fmt.Errorf("%v", r)
		}
		stack := make([]byte, 4<<10) // 4 KB
		length := runtime.Stack(stack, true)
		logrus.WithFields(logrus.Fields{
			"stack": string(stack[:length]),
		}).WithError(err).Error("[PANIC RECOVER]")
	}

}
