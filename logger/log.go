package logger

import (
	"os"
	"sync"

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
	provider *logrus.Logger
	config   *config
}

var (
	l      Log = new()
	doOnce sync.Once
)

func (l *logger) setup(c *config) {
	l.provider.SetOutput(os.Stdout)

	if c.flags.trace {
		l.provider.SetReportCaller(true)
		l.provider.Debug("init(): trace enabled")
	}

	if c.flags.dev {
		l.provider.SetLevel(logrus.DebugLevel)
		l.provider.Debug("init(): dev environment")
	}

	l.config = c
}

func new(opts ...Option) *logger {
	l := logger{
		provider: logrus.StandardLogger(),
	}

	c := newConfig()
	for _, opt := range opts {
		opt(c)
	}

	l.setup(c)

	return &l
}

func New(opts ...Option) {
	doOnce.Do(func() {
		l = new(opts...)
	})
}
