package logger

// Level defines log levels.
type Level int8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	TraceLevel Level = -1
)

func stringToLevel(s string) Level {
	switch s {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	case "trace":
		return TraceLevel
	default:
		return InfoLevel
	}
}
