package zerolog

import (
	"io"
	"os"

	"github.com/rs/zerolog"

	"github.com/rnymphaea/chronoflow/auth/internal/config"
	"github.com/rnymphaea/chronoflow/auth/internal/logger"
)

type Logger struct {
	zl zerolog.Logger
}

func New(cfg config.LoggerConfig) *Logger {
	logLevel, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		logLevel = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(logLevel)

	var output io.Writer
	if cfg.Pretty {
		output = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: zerolog.TimeFormatUnix,
		}
	} else {
		output = os.Stdout
	}

	base := zerolog.New(output).With().Timestamp().CallerWithSkipFrameCount(3).Logger()

	return &Logger{base}
}

func (l *Logger) Debug(msg string, args ...any) {
	l.zl.Debug().Fields(parseArgs(args...)).Msg(msg)
}

func (l *Logger) Info(msg string, args ...any) {
	l.zl.Info().Fields(parseArgs(args...)).Msg(msg)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.zl.Warn().Fields(parseArgs(args...)).Msg(msg)
}

func (l *Logger) Error(msg string, args ...any) {
	l.zl.Error().Fields(parseArgs(args...)).Msg(msg)
}

func (l *Logger) With(args ...any) logger.Logger {
	return &Logger{
		zl: l.zl.With().Fields(parseArgs(args...)).Logger(),
	}
}

func parseArgs(args ...any) map[string]any {
	if len(args) == 0 {
		return nil
	}

	fields := make(map[string]any)
	for i := 0; i < len(args); i++ {
		key, ok := args[i].(string)
		if !ok {
			continue
		}

		if i+1 < len(args) {
			fields[key] = args[i+1]
			i++
		} else {
			fields[key] = nil
		}
	}

	return fields
}
