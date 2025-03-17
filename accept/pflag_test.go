package accept_test

import (
	"context"
	"io"
	"log/slog"
	"testing"

	"github.com/UiP9AV6Y/go-slog-adapter/stdflags"

	"github.com/spf13/pflag"
	"gotest.tools/v3/assert"
)

func TestPFlag(t *testing.T) {
	subject := pflag.NewFlagSet("test", pflag.ContinueOnError)
	lvl := slog.LevelDebug
	fmt := stdflags.LogFormatText()
	lvlFlag := stdflags.NewLogLevelFlag(&lvl)
	fmtFlag := stdflags.NewLogFormatFlag(fmt)

	subject.AddGoFlag(lvlFlag)
	subject.AddGoFlag(fmtFlag)

	gotErr := subject.Parse([]string{"--log.level", "warn", "--log.format", "json"})
	assert.NilError(t, gotErr)

	assert.Equal(t, lvl.String(), slog.LevelWarn.String())
	assert.Equal(t, fmt.String(), stdflags.LogFormatJSON().String())

	haveContext := context.Background()
	gotHandler := stdflags.NewHandler(io.Discard, fmt, lvl)
	assert.Assert(t, gotHandler != nil)

	assert.Assert(t, !gotHandler.Enabled(haveContext, slog.LevelInfo))
	assert.Assert(t, gotHandler.Enabled(haveContext, slog.LevelWarn))
	assert.Assert(t, gotHandler.Enabled(haveContext, slog.LevelError))
}
