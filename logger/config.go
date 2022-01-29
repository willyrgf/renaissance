package logger

import "flag"

type (
	flags struct {
		dev   bool
		trace bool
	}
	config struct {
		flags flags
	}
)

func newConfig() *config {
	return &config{
		flags: flags{
			dev:   false,
			trace: false,
		},
	}
}

func WithFlags(c *config) {
	c.flags.dev = *(flag.Bool("dev", true, "Set the environment to dev."))
	c.flags.trace = *(flag.Bool("trace", false, "Enable tracing."))
	flag.Parse()
}
