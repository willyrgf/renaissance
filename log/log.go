package log

import (
	"github.com/sirupsen/logrus"
)

type Option func(*config)

type Log interface {
	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
}

type logger struct {
	provider logrus.Logger
	config   *config
}

var _ Log = (*logger)(nil)

func (l *logger) setup(c *config) {
	if *c.flags.trace {
		logrus.SetReportCaller(true)
		logrus.Debug("init(): trace enabled")
	}

	if *c.flags.dev {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("init(): dev environment")
	}

	l.config = c
}

func new(opts ...Option) *logger {
	var l logger

	c := newConfig()
	for _, opt := range opts {
		opt(c)
	}

	l.setup(c)

	return &l
}

func New(opts ...Option) Log {
	return new(opts...)
}
