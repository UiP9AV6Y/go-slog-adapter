package slogadapter

import (
	"context"
	"io"
	"log/slog"
	"runtime"
	"time"
)

// Logging adapter for structured logging. It implements
// a variety of log functions suitable for various custom
// logging contracts.
type SlogAdapter struct {
	logger *slog.Logger
	level  slog.Level
}

// New returns a [SlogAdapter] instance using the provided
// [slog.Logger] as data sink. The [slog.Level] is used as
// default value for logging functions without an explicit
// level semantic attached to it.
func New(logger *slog.Logger, level slog.Level) *SlogAdapter {
	return &SlogAdapter{
		logger: logger,
		level:  level,
	}
}

// NewLogger returns a [New] instance using the provided
// [slog.Logger] as data sink and [slog.LevelInfo] as
// default value for logging functions without an explicit
// level semantic attached to it.
func NewLogger(logger *slog.Logger) *SlogAdapter {
	return New(logger, slog.LevelInfo)
}

// NewLevel returns a [New] instance using the [slog.Default]
// instance as data sink. The [slog.Level] is used as
// default value for logging functions without an explicit
// level semantic attached to it.
func NewLevel(level slog.Level) *SlogAdapter {
	return New(slog.Default(), level)
}

// NewLevel returns a [New] instance using the [slog.Default]
// instance as data sink and [slog.LevelInfo] as
// default value for logging functions without an explicit
// level semantic attached to it.
func NewDefault() *SlogAdapter {
	return New(slog.Default(), slog.LevelInfo)
}

// NewText creates a [New] instance using a [slog.TextHandler]
// with the provided writer and level.
func NewText(writer io.Writer, level slog.Level) *SlogAdapter {
	opts := &slog.HandlerOptions{
		Level: level,
	}
	handler := slog.NewTextHandler(writer, opts)

	return New(slog.New(handler), level)
}

// NewJSON creates a [New] instance using a [slog.JSONHandler]
// with the provided writer and level.
func NewJSON(writer io.Writer, level slog.Level) *SlogAdapter {
	opts := &slog.HandlerOptions{
		Level: level,
	}
	handler := slog.NewJSONHandler(writer, opts)

	return New(slog.New(handler), level)
}

// Level returns the configured [slog.Level]
func (a *SlogAdapter) Level() slog.Level {
	return a.level
}

// Logger returns the internal [slog.Logger] instance
func (a *SlogAdapter) Logger() *slog.Logger {
	return a.logger
}

// log is the heart of the adapter. it is the central dispatcher for all
// interface implementations.
func (a *SlogAdapter) log(ctx context.Context, level slog.Level, msg string) {
	if !a.logger.Enabled(ctx, level) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(3, pcs[:]) // skip [Callers, log, API]
	r := slog.NewRecord(time.Now(), level, msg, pcs[0])
	_ = a.logger.Handler().Handle(ctx, r)
}
