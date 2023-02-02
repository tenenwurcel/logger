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

	Debug(msg string, fields ...*Field)
	Info(msg string, fields ...*Field)
	Warn(msg string, fields ...*Field)
	Error(msg string, fields ...*Field)
	Fatal(msg string, fields ...*Field)
	Panic(msg string, fields ...*Field)
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

func (l *Logger) Debug(msg string, fields ...*Field) {
	if l.logger == nil {
		return
	}

	l.logger.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...*Field) {
	if l.logger == nil {
		return
	}

	l.logger.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...*Field) {
	if l.logger == nil {
		return
	}

	l.logger.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...*Field) {
	if l.logger == nil {
		return
	}

	l.logger.Error(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...*Field) {
	if l.logger == nil {
		return
	}

	l.logger.Fatal(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...*Field) {
	if l.logger == nil {
		return
	}

	l.logger.Panic(msg, fields...)
}

func (l *Logger) WithContext(ctx context.Context) context.Context {
	if l.logger == nil {
		return ctx
	}

	ctx = context.WithValue(ctx, contextKey{}, l.logger)

	return ctx
}

func FromContext(ctx context.Context) LoggerI {
	if ctx == nil {
		return nil
	}

	if li, ok := ctx.Value(contextKey{}).(LoggerI); ok {
		return li
	}

	return nil
}
