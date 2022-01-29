package logger

func Tracef(format string, args ...interface{}) {
	l.Tracef(format, args...)
}

func Debugf(format string, args ...interface{}) {
	l.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	l.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	l.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	l.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	l.Panicf(format, args...)
}
