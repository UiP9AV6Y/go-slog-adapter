package slogadapter_test

import (
	"io"
	"log/slog"

	"github.com/UiP9AV6Y/go-slog-adapter"
	"github.com/UiP9AV6Y/go-slog-adapter/testing"
)

func newSlogHandler(w io.Writer, level slog.Level) slog.Handler {
	opts := &slog.HandlerOptions{
		Level:       level,
		AddSource:   true,
		ReplaceAttr: testing.DeterministicAttr,
	}

	return slog.NewTextHandler(w, opts)
}

func newSlogNopHandler(level slog.Level) slog.Handler {
	opts := &slog.HandlerOptions{
		Level: level,
	}

	return slog.NewTextHandler(io.Discard, opts)
}

func newSlogAdapterLoggerSubject(logger *slog.Logger, level slog.Level) *slogadapter.SlogAdapter {
	return slogadapter.New(logger, level)
}

func newSlogAdapterSubject(w io.Writer, level slog.Level) *slogadapter.SlogAdapter {
	h := newSlogHandler(w, slog.LevelDebug) // ensure everything is reported

	return slogadapter.New(slog.New(h), level)
}
