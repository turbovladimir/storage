package mylog

import (
	"golang.org/x/exp/slog"
	"os"
)

func Init() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	return logger
}
