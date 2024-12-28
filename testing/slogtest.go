package testing

import (
	"log/slog"
	"path/filepath"
)

// DeterministicAttr removes the top-level time attribute
// and replaces the top-level source attribute with a
// value without any directories.
// It is intended to be used as a ReplaceAttr function,
// to make example output deterministic.
func DeterministicAttr(groups []string, a slog.Attr) slog.Attr {
	// Remove time from the output for predictable test output.
	if a.Key == slog.TimeKey && len(groups) == 0 {
		return slog.Attr{}
	}
	// Remove the directory from the source's filename.
	if a.Key == slog.SourceKey && len(groups) == 0 {
		source := a.Value.Any().(*slog.Source)
		source.File = filepath.Base(source.File)
	}
	return a
}
