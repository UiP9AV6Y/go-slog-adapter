package stdflags

import (
	"flag"
	"io"
	"log/slog"
	"os"

	"github.com/UiP9AV6Y/go-slog-adapter"
)

const (
	FlagUsageLevel  = "Log level. Valid values include debug, info, warn, and error"
	FlagUsageFormat = "Log format. Valid values include text, and json"
)

const (
	FlagLogLevel  = "log.level"
	FlagLogFormat = "log.format"
)

const (
	EnvLogLevel  = "LOG_LEVEL"
	EnvLogFormat = "LOG_FORMAT"
)

// LogFlags is a data storage for [flag.FlagSet] parsing results.
type LogFlags struct {
	lvlValue slog.Level
	fmtValue LogFormatValue
	lvlEnv   string
	fmtEnv   string
	fs       *flag.FlagSet
}

func newLogFlags(fs *flag.FlagSet) *LogFlags {
	result := &LogFlags{
		lvlValue: slog.LevelInfo,
		fmtValue: LogFormatText(),
		lvlEnv:   EnvLogLevel,
		fmtEnv:   EnvLogFormat,
		fs:       fs,
	}

	return result
}

// NewLogFlags returns a [LogFlags] instance with the given
// flagset primed for populating its internal state.
func NewLogFlags(fs *flag.FlagSet) *LogFlags {
	result := newLogFlags(fs)

	fs.TextVar(&result.lvlValue, FlagLogLevel, result.lvlValue, FlagUsageLevel)
	fs.TextVar(result.fmtValue, FlagLogFormat, result.fmtValue, FlagUsageFormat)

	return result
}

// NewEnvLogFlags returns a [LogFlags] instance with the given
// flagset primed for populating its internal state. he flags usage
// description will include mentions of environment variables.
// If this is not desired, use [NewLogFlags].
//
// The environment variables will use the optional prefix as-is.
func NewEnvLogFlags(fs *flag.FlagSet, prefix string) *LogFlags {
	result := newLogFlags(fs)
	result.lvlEnv = prefix + EnvLogLevel
	result.fmtEnv = prefix + EnvLogFormat

	fs.TextVar(&result.lvlValue, FlagLogLevel, result.lvlValue, FlagUsageLevel+" [$"+result.lvlEnv+"]")
	fs.TextVar(result.fmtValue, FlagLogFormat, result.fmtValue, FlagUsageFormat+" [$"+result.fmtEnv+"]")

	return result
}

// ParseFunc uses the provided function to retrieve values
// for the previously provisioned flagset. Values are only
// forwarded if they are not empty. Returned errors originate
// from the flag parsing logic.
func (f *LogFlags) ParseFunc(get func(string) string) error {
	if lvl := get(f.lvlEnv); lvl != "" {
		if err := f.fs.Set(FlagLogLevel, lvl); err != nil {
			return err
		}
	}

	if fmt := get(f.fmtEnv); fmt != "" {
		if err := f.fs.Set(FlagLogFormat, fmt); err != nil {
			return err
		}
	}

	return nil
}

// ParseEnv calls [LogFlags.ParseFunc] with [os.Getenv].
func (f *LogFlags) ParseEnv() error {
	return f.ParseFunc(os.Getenv)
}

// Adapter returns a [slogadapter.SlogAdapter] with a newly created [slog.Logger].
// The provided handler options are optional.
func (f *LogFlags) Adapter(w io.Writer, opts *slog.HandlerOptions) *slogadapter.SlogAdapter {
	logger := slog.New(f.Handler(w, opts))

	return slogadapter.New(logger, f.lvlValue)
}

// Level returns the parsed level.
func (f *LogFlags) Level() slog.Level {
	return f.lvlValue
}

// Handler returns a [slog.Handler] from the internal state.
// If no handler options are provided, a default
// set will be created. In either case, the level
// field will be overwritten with the value of [LogFlags.Level].
func (f *LogFlags) Handler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	if opts == nil {
		opts = &slog.HandlerOptions{Level: f.lvlValue}
	} else {
		opts = &(*opts)
		opts.Level = f.lvlValue
	}

	return f.fmtValue.Handler(w, opts)
}
