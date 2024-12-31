package stdflags_test

import (
	"flag"
	"io"
	"log/slog"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/golden"

	"github.com/UiP9AV6Y/go-slog-adapter/stdflags"
)

func TestNewFlagSetError(t *testing.T) {
	var buf strings.Builder
	fs := flag.NewFlagSet("plain-flagset", flag.ContinueOnError)
	fs.SetOutput(&buf)
	_ = stdflags.NewLogFlags(fs)
	argv := []string{"-h"}

	err := fs.Parse(argv)
	assert.Error(t, err, flag.ErrHelp.Error())
	golden.Assert(t, buf.String(), "help_plain.golden")
}

func TestNewEnvFlagSetError(t *testing.T) {
	var buf strings.Builder
	fs := flag.NewFlagSet("env-flagset", flag.ContinueOnError)
	fs.SetOutput(&buf)
	_ = stdflags.NewEnvLogFlags(fs, "TEST_")
	argv := []string{"-h"}

	err := fs.Parse(argv)
	assert.Error(t, err, flag.ErrHelp.Error())
	golden.Assert(t, buf.String(), "help_env.golden")
}

func TestNewFlagSet(t *testing.T) {
	tests := map[string]struct {
		haveArgv    []string
		wantHandler string
		wantLevel   slog.Level
		wantError   string
	}{
		"empty": {
			haveArgv:    []string{},
			wantHandler: "*slog.TextHandler",
			wantLevel:   slog.LevelInfo,
		},
		"both": {
			haveArgv:    []string{"-log.format", "json", "-log.level", "debug"},
			wantHandler: "*slog.JSONHandler",
			wantLevel:   slog.LevelDebug,
		},
		"json": {
			haveArgv:    []string{"-log.format", "json"},
			wantHandler: "*slog.JSONHandler",
			wantLevel:   slog.LevelInfo,
		},
		"text": {
			haveArgv:    []string{"-log.format", "text"},
			wantHandler: "*slog.TextHandler",
			wantLevel:   slog.LevelInfo,
		},
		"markdown": {
			haveArgv:  []string{"-log.format", "markdown"},
			wantError: `invalid value "markdown" for flag -log.format: unsupported log format`,
		},
		"debug": {
			haveArgv:  []string{"-log.level", "debug"},
			wantLevel: slog.LevelDebug,
		},
		"DEBUG": {
			haveArgv:  []string{"-log.level", "DEBUG"},
			wantLevel: slog.LevelDebug,
		},
		"info+2": {
			haveArgv:  []string{"-log.level", "info+2"},
			wantLevel: slog.LevelInfo + 2,
		},
		"INFO+2": {
			haveArgv:  []string{"-log.level", "INFO+2"},
			wantLevel: slog.LevelInfo + 2,
		},
		"fatal": {
			haveArgv:  []string{"-log.level", "fatal"},
			wantError: `invalid value "fatal" for flag -log.level: slog: level string "fatal": unknown name`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var buf strings.Builder
			fs := flag.NewFlagSet(name, flag.ContinueOnError)
			fs.SetOutput(&buf)
			subject := stdflags.NewLogFlags(fs)

			err := fs.Parse(test.haveArgv)
			if test.wantError != "" {
				assert.ErrorContains(t, err, test.wantError)
				return
			}

			assert.NilError(t, err)
			assert.Assert(t, buf.Len() == 0, "NewFlagSet(%q) encounterd parsing error: %s", test.haveArgv, buf.String())

			gotLevel := subject.Level()
			assert.Assert(t, gotLevel == test.wantLevel, "NewFlagSet(%q) returned wrong slog level; got=%v, want=%v", test.haveArgv, gotLevel, test.wantLevel)

			if test.wantHandler != "" {
				gotHandler := subject.Handler(io.Discard, nil)
				gotName := reflect.TypeOf(gotHandler).String()
				assert.Assert(t, gotName == test.wantHandler, "NewFlagSet(%q) returned wrong slog handler; got=%v, want=%v", test.haveArgv, gotName, test.wantHandler)
			}
		})
	}
}

func TestNewEnvFlagSet(t *testing.T) {
	tests := map[string]struct {
		haveEnv     url.Values
		wantHandler string
		wantLevel   slog.Level
		wantError   string
	}{
		"empty": {
			haveEnv:     map[string][]string{},
			wantHandler: "*slog.TextHandler",
			wantLevel:   slog.LevelInfo,
		},
		"both": {
			haveEnv: map[string][]string{
				"LOG_FORMAT": []string{"json"},
				"LOG_LEVEL":  []string{"debug"},
			},
			wantHandler: "*slog.JSONHandler",
			wantLevel:   slog.LevelDebug,
		},
		"json": {
			haveEnv:     map[string][]string{"LOG_FORMAT": []string{"json"}},
			wantHandler: "*slog.JSONHandler",
			wantLevel:   slog.LevelInfo,
		},
		"text": {
			haveEnv:     map[string][]string{"LOG_FORMAT": []string{"text"}},
			wantHandler: "*slog.TextHandler",
			wantLevel:   slog.LevelInfo,
		},
		"markdown": {
			haveEnv:   map[string][]string{"LOG_FORMAT": []string{"markdown"}},
			wantError: `unsupported log format`,
		},
		"debug": {
			haveEnv:   map[string][]string{"LOG_LEVEL": []string{"debug"}},
			wantLevel: slog.LevelDebug,
		},
		"DEBUG": {
			haveEnv:   map[string][]string{"LOG_LEVEL": []string{"DEBUG"}},
			wantLevel: slog.LevelDebug,
		},
		"info+2": {
			haveEnv:   map[string][]string{"LOG_LEVEL": []string{"info+2"}},
			wantLevel: slog.LevelInfo + 2,
		},
		"INFO+2": {
			haveEnv:   map[string][]string{"LOG_LEVEL": []string{"INFO+2"}},
			wantLevel: slog.LevelInfo + 2,
		},
		"fatal": {
			haveEnv:   map[string][]string{"LOG_LEVEL": []string{"fatal"}},
			wantError: `slog: level string "fatal": unknown name`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var buf strings.Builder
			fs := flag.NewFlagSet(name, flag.ContinueOnError)
			fs.SetOutput(&buf)
			subject := stdflags.NewLogFlags(fs)

			err := subject.ParseFunc(test.haveEnv.Get)
			if test.wantError != "" {
				assert.ErrorContains(t, err, test.wantError)
				return
			}

			assert.NilError(t, err)
			assert.Assert(t, buf.Len() == 0, "NewEnvFlagSet(%q) encounterd parsing error: %s", test.haveEnv, buf.String())

			gotLevel := subject.Level()
			assert.Assert(t, gotLevel == test.wantLevel, "NewEnvFlagSet(%q) returned wrong slog level; got=%v, want=%v", test.haveEnv, gotLevel, test.wantLevel)

			if test.wantHandler != "" {
				gotHandler := subject.Handler(io.Discard, nil)
				gotName := reflect.TypeOf(gotHandler).String()
				assert.Assert(t, gotName == test.wantHandler, "NewEnvFlagSet(%q) returned wrong slog handler; got=%v, want=%v", test.haveEnv, gotName, test.wantHandler)
			}
		})
	}
}
