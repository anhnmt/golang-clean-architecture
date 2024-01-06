package logger

import (
	"log"
	"log/slog"
	"path/filepath"

	"github.com/anhnmt/golang-clean-architecture/pkg/config"
)

func New(cfg config.Log) {
	replace := func(groups []string, a slog.Attr) slog.Attr {
		// Remove the directory from the source's filename.
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			if source.File == "" {
				return slog.Attr{}
			}

			source.File = filepath.Base(source.File)
		}

		return a
	}

	var level slog.Level
	err := level.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		level = slog.LevelInfo
	}

	writer := log.Writer()

	opts := &slog.HandlerOptions{
		Level:       level,
		AddSource:   true,
		ReplaceAttr: replace,
	}

	// slog handler
	var handler slog.Handler

	if cfg.Format == "json" {
		handler = slog.NewJSONHandler(writer, opts)
	} else {
		handler = slog.NewTextHandler(writer, opts)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
}
