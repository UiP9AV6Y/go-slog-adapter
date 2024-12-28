package slogadapter

import (
	"context"
)

// EnabledContextLogger is the same implementation as [EnabledLogger]
// with the addition of a custom context
type EnabledContextLogger interface {
	// EnabledContext is the same implementation as [EnabledLogger.Enabled]
	// with the addition of a custom context
	EnabledContext(ctx context.Context) bool
}

// PrintContextLogger is the same implementation as [PrintLogger]
// with the addition of a custom context
type PrintContextLogger interface {
	// PrintContext is the same implementation as [PrintLogger.Print]
	// with the addition of a custom context
	PrintContext(ctx context.Context, msg ...interface{})
}

// PrintFormatContextLogger is the same implementation as [PrintFormatLogger]
// with the addition of a custom context
type PrintFormatContextLogger interface {
	// PrintfContext is the same implementation as [PrintFormatLogger.Printf]
	// with the addition of a custom context
	PrintfContext(ctx context.Context, format string, args ...interface{})
}

// PrintLineContextLogger is the same implementation as [PrintLineLogger]
// with the addition of a custom context
type PrintLineContextLogger interface {
	// PrintlnContext is the same implementation as [PrintLineLogger.Println]
	// with the addition of a custom context
	PrintlnContext(ctx context.Context, msg ...interface{})
}

// LeveledContextLogger is the same implementation as [LeveledLogger]
// with the addition of a custom context
type LeveledContextLogger interface {
	// ErrorContext is the same implementation as [LeveledLogger.Error]
	// with the addition of a custom context
	ErrorContext(ctx context.Context, msg ...interface{})
	// InfoContext is the same implementation as [LeveledLogger.Info]
	// with the addition of a custom context
	InfoContext(ctx context.Context, msg ...interface{})
	// DebugContext is the same implementation as [LeveledLogger.Debug]
	// with the addition of a custom context
	DebugContext(ctx context.Context, msg ...interface{})
	// WarnContext is the same implementation as [LeveledLogger.Warn]
	// with the addition of a custom context
	WarnContext(ctx context.Context, msg ...interface{})
}

// LeveledFormatContextLogger is the same implementation as [LeveledFormatLogger]
// with the addition of a custom context
type LeveledFormatContextLogger interface {
	// ErrorfContext is the same implementation as [LeveledFormatLogger.Errorf]
	// with the addition of a custom context
	ErrorfContext(ctx context.Context, format string, args ...interface{})
	// InfofContext is the same implementation as [LeveledFormatLogger.Infof]
	// with the addition of a custom context
	InfofContext(ctx context.Context, format string, args ...interface{})
	// DebugfContext is the same implementation as [LeveledFormatLogger.Debugf]
	// with the addition of a custom context
	DebugfContext(ctx context.Context, format string, args ...interface{})
	// WarnfContext is the same implementation as [LeveledFormatLogger.Warnf]
	// with the addition of a custom context
	WarnfContext(ctx context.Context, format string, args ...interface{})
}

// LeveledLineContextLogger is the same implementation as [LeveledLineLogger]
// with the addition of a custom context
type LeveledLineContextLogger interface {
	// ErrorlnContext is the same implementation as [LeveledLineLogger.Errorln]
	// with the addition of a custom context
	ErrorlnContext(ctx context.Context, msg ...interface{})
	// InfolnContext is the same implementation as [LeveledLineLogger.Infoln]
	// with the addition of a custom context
	InfolnContext(ctx context.Context, msg ...interface{})
	// DebuglnContext is the same implementation as [LeveledLineLogger.Debugln]
	// with the addition of a custom context
	DebuglnContext(ctx context.Context, msg ...interface{})
	// WarnlnContext is the same implementation as [LeveledLineLogger.Warnln]
	// with the addition of a custom context
	WarnlnContext(ctx context.Context, msg ...interface{})
}
