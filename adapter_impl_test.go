package slogadapter_test

import (
	"log/slog"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/golden"
)

func TestAdapter_Enabled(t *testing.T) {
	handler := newSlogNopHandler(slog.LevelInfo)
	logger := slog.New(handler)

	subject := newSlogAdapterLoggerSubject(logger, slog.LevelInfo)
	assert.Assert(t, subject.Enabled(), "Info adapter is enabled")

	subject = newSlogAdapterLoggerSubject(logger, slog.LevelWarn)
	assert.Assert(t, subject.Enabled(), "Warn adapter is enabled")

	subject = newSlogAdapterLoggerSubject(logger, slog.LevelDebug)
	assert.Assert(t, !subject.Enabled(), "Debug adapter is disabled")
}

func TestAdapter_PrintLogger(t *testing.T) {
	var buf strings.Builder
	subject := newSlogAdapterSubject(&buf, slog.LevelInfo)

	subject.Print("hello", "world")
	subject.Print("test", "test")
	subject.Print("one", 2, "three")

	golden.Assert(t, buf.String(), "print.golden")
}

func TestAdapter_PrintfLogger(t *testing.T) {
	var buf strings.Builder
	subject := newSlogAdapterSubject(&buf, slog.LevelInfo)

	subject.Printf("%s %s", "hello", "world")
	subject.Printf("%s %s", "test", "test")
	subject.Printf("%s%d%s", "one", 2, "three")

	golden.Assert(t, buf.String(), "printf.golden")
}

func TestAdapter_PrintlnLogger(t *testing.T) {
	var buf strings.Builder
	subject := newSlogAdapterSubject(&buf, slog.LevelInfo)

	subject.Println("hello", "world")
	subject.Println("test", "test")
	subject.Println("one", 2, "three")

	golden.Assert(t, buf.String(), "println.golden")
}

func TestAdapter_LeveledLogger(t *testing.T) {
	var buf strings.Builder
	subject := newSlogAdapterSubject(&buf, slog.LevelInfo)

	subject.Error("error", 4, "message")
	subject.Info("info", 2, "message")
	subject.Debug("debug", 1, "message")
	subject.Warn("warn", 3, "message")

	golden.Assert(t, buf.String(), "leveled_logger.golden")
}

func TestAdapter_LeveledFormatLogger(t *testing.T) {
	var buf strings.Builder
	subject := newSlogAdapterSubject(&buf, slog.LevelInfo)

	subject.Errorf("%s %d %s", "error", 4, "message")
	subject.Infof("%s %d %s", "info", 2, "message")
	subject.Debugf("%s %d %s", "debug", 1, "message")
	subject.Warnf("%s %d %s", "warn", 3, "message")

	golden.Assert(t, buf.String(), "leveled_format_logger.golden")
}

func TestAdapter_LeveledLineLogger(t *testing.T) {
	var buf strings.Builder
	subject := newSlogAdapterSubject(&buf, slog.LevelInfo)

	subject.Errorln("error", 4, "message")
	subject.Infoln("info", 2, "message")
	subject.Debugln("debug", 1, "message")
	subject.Warnln("warn", 3, "message")

	golden.Assert(t, buf.String(), "leveled_line_logger.golden")
}
