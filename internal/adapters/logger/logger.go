package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitLogger(level slog.Level) {
	Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))
	slog.SetDefault(Logger)

}

func Log() *slog.Logger {
	return Logger
}

func LogErr(err error) {
	Log().Error("error", "error", err)
}
