package raft

import (
	"context"
	"io"
	"log"
	"log/slog"
	"strings"

	"github.com/hashicorp/go-hclog"
	"github.com/nonchan7720/user-flex-feature/pkg/utils/collection"
)

type wrapHcLogger struct {
	*slog.Logger
	name string
}

func newWrapHcLogger(name string) hclog.Logger {
	return &wrapHcLogger{
		Logger: slog.Default(),
		name: strings.Join(collection.Filter([]string{"user-flex-user-raft", name}, func(v string) bool {
			return v != ""
		}), "."),
	}
}

var (
	_ hclog.Logger = (*wrapHcLogger)(nil)
)

func (l *wrapHcLogger) Log(level hclog.Level, msg string, args ...interface{}) {
	var lv slog.Level
	switch level {
	case hclog.Trace:
		lv = slog.LevelError
	case hclog.Debug:
		lv = slog.LevelDebug
	case hclog.Info:
		lv = slog.LevelInfo
	case hclog.Warn:
		lv = slog.LevelWarn
	case hclog.Error:
		lv = slog.LevelError
	default:
		return
	}
	l.Logger.Log(context.Background(), lv, msg, args...)
}
func (l *wrapHcLogger) Trace(msg string, args ...interface{}) {
	l.Log(hclog.Trace, msg, args...)
}
func (l *wrapHcLogger) Debug(msg string, args ...interface{}) {
	l.Log(hclog.Debug, msg, args...)
}
func (l *wrapHcLogger) Info(msg string, args ...interface{}) {
	l.Log(hclog.Info, msg, args...)
}
func (l *wrapHcLogger) Warn(msg string, args ...interface{}) {
	l.Log(hclog.Warn, msg, args...)
}
func (l *wrapHcLogger) Error(msg string, args ...interface{}) {
	l.Log(hclog.Warn, msg, args...)
}

func (l *wrapHcLogger) IsTrace() bool {
	return slog.Default().Enabled(context.Background(), slog.LevelError)
}

func (l *wrapHcLogger) IsDebug() bool {
	return slog.Default().Enabled(context.Background(), slog.LevelDebug)
}

func (l *wrapHcLogger) IsInfo() bool {
	return slog.Default().Enabled(context.Background(), slog.LevelInfo)
}

func (l *wrapHcLogger) IsWarn() bool {
	return slog.Default().Enabled(context.Background(), slog.LevelWarn)
}

func (l *wrapHcLogger) IsError() bool {
	return slog.Default().Enabled(context.Background(), slog.LevelError)
}

func (l *wrapHcLogger) ImpliedArgs() []interface{} {
	return nil
}

func (l *wrapHcLogger) With(args ...interface{}) hclog.Logger {
	return &wrapHcLogger{
		Logger: slog.With(args...),
	}
}

func (l *wrapHcLogger) Name() string {
	return l.name
}

func (l *wrapHcLogger) Named(name string) hclog.Logger {
	return newWrapHcLogger(name)
}

func (l *wrapHcLogger) ResetNamed(name string) hclog.Logger {
	return newWrapHcLogger(l.name)
}

func (l *wrapHcLogger) SetLevel(level hclog.Level) {
}

func (l *wrapHcLogger) GetLevel() hclog.Level {
	return hclog.Info
}

func (l *wrapHcLogger) StandardLogger(opts *hclog.StandardLoggerOptions) *log.Logger {
	return nil
}

func (l *wrapHcLogger) StandardWriter(opts *hclog.StandardLoggerOptions) io.Writer {
	return nil
}
