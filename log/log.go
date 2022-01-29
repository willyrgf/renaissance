package log

import (
	"github.com/sirupsen/logrus"
)

type Option func(*config)

type logger struct {
	provider logrus.Logger
}

type Log interface {
}

func (l *logger) setup(c *config) {
	if *c.flags.trace {
		logrus.SetReportCaller(true)
		logrus.Debug("init(): trace enabled")
	}

	if *c.flags.dev {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("init(): dev environment")
	}

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
