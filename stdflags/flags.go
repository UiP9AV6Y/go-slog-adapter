package stdflags

import (
	"flag"
	"log/slog"
)

type logLevelValue struct {
	p *slog.Level
}

func (v logLevelValue) Set(s string) error {
	return v.p.UnmarshalText([]byte(s))
}

func (v logLevelValue) String() string {
	return v.p.String()
}

// NewLogLevelFlag returns a [flag.Flag] instance
// wrapping the provided level pointer. It is primarily
// intended for integration with other logging frameworks.
func NewLogLevelFlag(l *slog.Level) *flag.Flag {
	val := logLevelValue{l}
	result := &flag.Flag{
		Name:     FlagLogLevel,
		Usage:    FlagUsageLevel,
		Value:    val,
		DefValue: l.String(),
	}

	return result
}

// NewLogFormatValue returns a [flag.Flag] instance
// wrapping the provided format value. It is primarily
// intended for integration with other logging frameworks.
func NewLogFormatFlag(f LogFormatValue) *flag.Flag {
	result := &flag.Flag{
		Name:     FlagLogFormat,
		Usage:    FlagUsageFormat,
		Value:    f,
		DefValue: f.String(),
	}

	return result
}
