package logger

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelSuccess
	LevelWarn
	LevelError
	LevelPanic
)

var levelLabels = map[Level]string{
	LevelDebug:   "DEBUG",
	LevelInfo:    "INFO",
	LevelSuccess: "SUCCESS",
	LevelWarn:    "WARN",
	LevelError:   "ERROR",
	LevelPanic:   "PANIC",
}
