package logrec

import (
	log "codehub.huawei.com/hemyzhao/lib/hemyutils/logsdk"
)

// Debug debug
func Debug(flag bool, format string, v ...interface{}) {
	if flag {
		log.Debug(format, v...)
	}
}

// Info information
func Info(flag bool, format string, v ...interface{}) {
	if flag {
		log.Info(format, v...)
	}
}

// Warn for log
func Warn(flag bool, format string, v ...interface{}) {
	if flag {
		log.Warn(format, v...)
	}
}

// Error error
func Error(flag bool, format string, v ...interface{}) {
	if flag {
		log.Error(format, v...)
	}
}

// Critical critical
func Critical(flag bool, format string, v ...interface{}) {
	if flag {
		log.Critical(format, v...)
	}
}
