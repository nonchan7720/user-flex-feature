package logging

import (
	"io"
	"log"
	"log/slog"
)

type writer struct{}

var (
	_ io.Writer = (*writer)(nil)
)

func (w *writer) Write(buf []byte) (int, error) {
	slog.Info(string(buf))
	return len(buf), nil
}

var StdLogger = log.New(&writer{}, "goff", log.LstdFlags)
