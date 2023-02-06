package logger

import "context"

// Level defines log levels.
type Level int8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
	NoLevel
	Disabled
	TraceLevel Level = -1
)

type LoggerI interface {
	Init()
	SetLevel(level Level)

	Debug(msg string, fields ...interface{})
	Info(msg string, fields ...interface{})
	Warn(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
	Fatal(msg string, fields ...interface{})
	Panic(msg string, fields ...interface{})
}

type Logger struct {
	logger LoggerI
}

type contextKey struct{}

func New(li LoggerI) *Logger {
	l := &Logger{
		logger: li,
	}

	l.logger.Init()

	return l
}

func (l *Logger) SetLevel(level Level) {
	if l.logger == nil {
		return
	}

	l.logger.SetLevel(level)
}

func (l *Logger) Debug(msg string, fields ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Error(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Fatal(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Panic(msg, fields...)
}

func (l *Logger) WithContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, contextKey{}, l)

	return ctx
}

func FromContext(ctx context.Context) *Logger {
	if ctx == nil {
		return nil
	}

	if li, ok := ctx.Value(contextKey{}).(Logger); ok {
		return &li
	}

	return nil
}
