package logger

import (
	"log/slog"
	"os"

	slogenv "github.com/cbrewster/slog-env"
	"github.com/phsym/console-slog"
)

type LoggerOption struct {
	UseConsoleHandler bool
}

func NewLogger(opt *LoggerOption) *slog.Logger {
	var handler slog.Handler
	if opt.UseConsoleHandler {
		handler = console.NewHandler(os.Stderr, nil)
	} else {
		handler = slog.NewJSONHandler(os.Stderr, nil)
	}
	// refer to GO_LOG env
	handler = slogenv.NewHandler(handler)
	logger := slog.New(handler)
	return logger
}
