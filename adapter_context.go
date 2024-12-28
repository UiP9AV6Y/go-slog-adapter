package slogadapter

import (
	"context"
	"fmt"
	"log/slog"
)

func (a *SlogAdapter) EnabledContext(ctx context.Context) bool {
	return a.logger.Enabled(ctx, a.level)
}

func (a *SlogAdapter) PrintContext(ctx context.Context, msg ...interface{}) {
	a.log(ctx, a.level, fmt.Sprint(msg...))
}

func (a *SlogAdapter) PrintfContext(ctx context.Context, format string, args ...interface{}) {
	a.log(ctx, a.level, fmt.Sprintf(format, args...))
}

func (a *SlogAdapter) PrintlnContext(ctx context.Context, msg ...interface{}) {
	a.log(ctx, a.level, fmt.Sprint(msg...))
}

func (a *SlogAdapter) ErrorContext(ctx context.Context, msg ...interface{}) {
	a.log(ctx, slog.LevelError, fmt.Sprint(msg...))
}

func (a *SlogAdapter) InfoContext(ctx context.Context, msg ...interface{}) {
	a.log(ctx, slog.LevelInfo, fmt.Sprint(msg...))
}

func (a *SlogAdapter) DebugContext(ctx context.Context, msg ...interface{}) {
	a.log(ctx, slog.LevelDebug, fmt.Sprint(msg...))
}

func (a *SlogAdapter) WarnContext(ctx context.Context, msg ...interface{}) {
	a.log(ctx, slog.LevelWarn, fmt.Sprint(msg...))
}

func (a *SlogAdapter) ErrorfContext(ctx context.Context, format string, args ...interface{}) {
	a.log(ctx, slog.LevelError, fmt.Sprintf(format, args...))
}

func (a *SlogAdapter) InfofContext(ctx context.Context, format string, args ...interface{}) {
	a.log(ctx, slog.LevelInfo, fmt.Sprintf(format, args...))
}

func (a *SlogAdapter) DebugfContext(ctx context.Context, format string, args ...interface{}) {
	a.log(ctx, slog.LevelDebug, fmt.Sprintf(format, args...))
}

func (a *SlogAdapter) WarnfContext(ctx context.Context, format string, args ...interface{}) {
	a.log(ctx, slog.LevelWarn, fmt.Sprintf(format, args...))
}

func (a *SlogAdapter) ErrorlnContext(ctx context.Context, msg ...interface{}) {
	a.log(ctx, slog.LevelError, fmt.Sprint(msg...))
}

func (a *SlogAdapter) InfolnContext(ctx context.Context, msg ...interface{}) {
	a.log(ctx, slog.LevelInfo, fmt.Sprint(msg...))
}

func (a *SlogAdapter) DebuglnContext(ctx context.Context, msg ...interface{}) {
	a.log(ctx, slog.LevelDebug, fmt.Sprint(msg...))
}

func (a *SlogAdapter) WarnlnContext(ctx context.Context, msg ...interface{}) {
	a.log(ctx, slog.LevelWarn, fmt.Sprint(msg...))
}
