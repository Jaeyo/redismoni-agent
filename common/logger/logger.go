package logger

import (
	"redismoni-agent/common/config"
	"fmt"
	"log/syslog"
	"log"
	"sync"
)

var infoSysLogger *log.Logger
var loggers map[syslog.Priority]*log.Logger
var once sync.Once

func initLoggers() {
	loggers = make(map[syslog.Priority]*log.Logger)
	for _, priority := range []syslog.Priority{syslog.LOG_INFO, syslog.LOG_DEBUG, syslog.LOG_WARNING, syslog.LOG_ERR} {
		syslogger, err := syslog.NewLogger(priority, 0)
		if (err != nil) {
			fmt.Println(err)
			continue
		}
		loggers[priority] = syslogger
	}
}

func priorityToString(priority syslog.Priority) string {
	switch priority {
	case syslog.LOG_INFO:
		return "INFO"
	case syslog.LOG_DEBUG:
		return "DEBUG"
	case syslog.LOG_WARNING:
		return "WARN"
	case syslog.LOG_ERR:
		return "ERROR"
	default:
		return ""
	}
}

func sendLog(priority syslog.Priority, msg ...interface{}) {
	once.Do(func() {
		initLoggers()
	})

	if config.GetDebug() {
		fmt.Println(priorityToString(priority), msg)
	}

	if syslogger, ok := loggers[priority]; ok {
		syslogger.Println(msg...)
	}
}

func Info(msg ...interface{}) {
	sendLog(syslog.LOG_INFO, msg...)
}

func Debug(msg ...interface{}) {
	sendLog(syslog.LOG_DEBUG, msg...)
}

func Warn(msg ...interface{}) {
	sendLog(syslog.LOG_WARNING, msg...)
}

func Error(msg ...interface{}) {
	sendLog(syslog.LOG_ERR, msg...)
}
