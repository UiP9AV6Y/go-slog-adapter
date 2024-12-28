package slogadapter

import (
	"context"
	"fmt"
	"log/slog"
)

func (a *SlogAdapter) Enabled() bool {
	return a.logger.Enabled(context.Background(), a.level)
}

func (a *SlogAdapter) Print(msg ...interface{}) {
	a.log(context.Background(), a.level, fmt.Sprint(msg...))
}

func (a *SlogAdapter) Printf(format string, args ...interface{}) {
	a.log(context.Background(), a.level, fmt.Sprintf(format, args...))
}

func (a *SlogAdapter) Println(msg ...interface{}) {
	a.log(context.Background(), a.level, fmt.Sprint(msg...))
}

func (a *SlogAdapter) Error(msg ...interface{}) {
	a.log(context.Background(), slog.LevelError, fmt.Sprint(msg...))
}

func (a *SlogAdapter) Info(msg ...interface{}) {
	a.log(context.Background(), slog.LevelInfo, fmt.Sprint(msg...))
}

func (a *SlogAdapter) Debug(msg ...interface{}) {
	a.log(context.Background(), slog.LevelDebug, fmt.Sprint(msg...))
}

func (a *SlogAdapter) Warn(msg ...interface{}) {
	a.log(context.Background(), slog.LevelWarn, fmt.Sprint(msg...))
}

func (a *SlogAdapter) Errorf(format string, args ...interface{}) {
	a.log(context.Background(), slog.LevelError, fmt.Sprintf(format, args...))
}

func (a *SlogAdapter) Infof(format string, args ...interface{}) {
	a.log(context.Background(), slog.LevelInfo, fmt.Sprintf(format, args...))
}

func (a *SlogAdapter) Debugf(format string, args ...interface{}) {
	a.log(context.Background(), slog.LevelDebug, fmt.Sprintf(format, args...))
}

func (a *SlogAdapter) Warnf(format string, args ...interface{}) {
	a.log(context.Background(), slog.LevelWarn, fmt.Sprintf(format, args...))
}

func (a *SlogAdapter) Errorln(msg ...interface{}) {
	a.log(context.Background(), slog.LevelError, fmt.Sprint(msg...))
}

func (a *SlogAdapter) Infoln(msg ...interface{}) {
	a.log(context.Background(), slog.LevelInfo, fmt.Sprint(msg...))
}

func (a *SlogAdapter) Debugln(msg ...interface{}) {
	a.log(context.Background(), slog.LevelDebug, fmt.Sprint(msg...))
}

func (a *SlogAdapter) Warnln(msg ...interface{}) {
	a.log(context.Background(), slog.LevelWarn, fmt.Sprint(msg...))
}
