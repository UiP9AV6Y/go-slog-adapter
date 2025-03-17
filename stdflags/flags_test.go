package stdflags_test

import (
	"log/slog"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/UiP9AV6Y/go-slog-adapter/stdflags"
)

func TestNewLogLevelFlag(t *testing.T) {
	have := slog.LevelDebug
	subject := stdflags.NewLogLevelFlag(&have)

	assert.Equal(t, subject.Name, stdflags.FlagLogLevel)
	assert.Equal(t, subject.Usage, stdflags.FlagUsageLevel)
	assert.Equal(t, subject.DefValue, slog.LevelDebug.String())
	assert.Equal(t, subject.Value.String(), slog.LevelDebug.String())

	assert.NilError(t, subject.Value.Set(slog.LevelError.String()))

	assert.Equal(t, have.String(), slog.LevelError.String())
}

func TestNewLogFormatFlag(t *testing.T) {
	have := stdflags.LogFormatJSON()
	subject := stdflags.NewLogFormatFlag(have)

	assert.Equal(t, subject.Name, stdflags.FlagLogFormat)
	assert.Equal(t, subject.Usage, stdflags.FlagUsageFormat)
	assert.Equal(t, subject.DefValue, stdflags.LogFormatJSON().String())
	assert.Equal(t, subject.Value.String(), stdflags.LogFormatJSON().String())

	assert.NilError(t, subject.Value.Set(stdflags.LogFormatText().String()))

	assert.Equal(t, have.String(), stdflags.LogFormatText().String())
}
