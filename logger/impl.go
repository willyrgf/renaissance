package logger

func (l *logger) Tracef(format string, args ...interface{}) {
	l.provider.Tracef(format, args...)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.provider.Debugf(format, args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.provider.Infof(format, args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.provider.Warnf(format, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.provider.Errorf(format, args...)
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.provider.Fatalf(format, args...)
}

func (l *logger) Panicf(format string, args ...interface{}) {
	l.provider.Panicf(format, args...)
}
