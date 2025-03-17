package stdflags

import (
	"encoding"
	"errors"
	"flag"
	"io"
	"log/slog"
)

type logFormat string

// LogFormatValue combines various interfaces for use with commandline flag parsing
type LogFormatValue interface {
	flag.Value
	encoding.TextMarshaler
	encoding.TextUnmarshaler

	// Handler returns a [slog.Handler] instance based on the represented format.
	// See [slog.NewJSONHandler] and [slog.NewTextHandler] for information about
	// parameter handling.
	Handler(io.Writer, *slog.HandlerOptions) slog.Handler
}

// ErrInvalidFormat is the error returned during unmarshalling of log format values.
var ErrInvalidFormat = errors.New("unsupported log format")

var (
	logFormatText = logFormat("text")
	logFormatJSON = logFormat("json")
)

// LogFormatText represents the slog.TextHandler.
func LogFormatText() LogFormatValue {
	lf := logFormatText

	return &lf
}

// LogFormatJSON represents the slog.JSONHandler.
func LogFormatJSON() LogFormatValue {
	lf := logFormatJSON

	return &lf
}

func (f logFormat) String() string {
	return string(f)
}

func (f logFormat) MarshalText() ([]byte, error) {
	return []byte(f), nil
}

func (f *logFormat) UnmarshalText(data []byte) error {
	return f.Set(string(data))
}

func (f *logFormat) Set(v string) (err error) {
	if v == "json" {
		*f = logFormatJSON
	} else if v == "text" {
		*f = logFormatText
	} else {
		err = ErrInvalidFormat
	}

	return
}

func (f logFormat) Handler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	if f == logFormatJSON {
		return slog.NewJSONHandler(w, opts)
	}

	return slog.NewTextHandler(w, opts)
}

// NewHandler create a [slog.Handler] whose constructor options are only
// initialized with the provided level.
// for more complex scenarios use [LogFormatValue.Handler] directly.
func NewHandler(w io.Writer, fmt LogFormatValue, lvl slog.Level) slog.Handler {
	opts := handlerOptions(lvl, nil)

	return fmt.Handler(w, opts)
}
