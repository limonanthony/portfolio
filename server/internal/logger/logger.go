package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Logger struct {
	std   *log.Logger
	level Level
	color bool
}

var l = &Logger{
	std:   log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
	level: LevelInfo,
	color: true,
}

const baseCallDepth = 3

func SetOutput(w io.Writer) { l.std.SetOutput(w) }
func SetLevel(level Level)  { l.level = level }
func EnableColor(on bool)   { l.color = on }

func (lg *Logger) printf(level Level, format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	label := levelLabels[level]

	if lg.color {
		col := levelColors[level]
		_ = lg.std.Output(baseCallDepth, fmt.Sprintf("%s%s: %s%s", col, label, msg, colorReset))
	} else {
		_ = lg.std.Output(baseCallDepth, fmt.Sprintf("%s %s", label, msg))
	}
}

func (lg *Logger) print(level Level, msg string) {
	if level < lg.level {
		return
	}

	label := levelLabels[level]

	if lg.color {
		col := levelColors[level]
		_ = lg.std.Output(baseCallDepth, fmt.Sprintf("%s%s: %s%s", col, label, msg, colorReset))
	} else {
		_ = lg.std.Output(baseCallDepth, fmt.Sprintf("%s %s", label, msg))
	}
}

func Debugf(format string, args ...any)   { l.printf(LevelDebug, format, args...) }
func Infof(format string, args ...any)    { l.printf(LevelInfo, format, args...) }
func Successf(format string, args ...any) { l.printf(LevelSuccess, format, args...) }
func Warnf(format string, args ...any)    { l.printf(LevelWarn, format, args...) }
func Errorf(format string, args ...any)   { l.printf(LevelError, format, args...) }
func Panicf(format string, args ...any) {
	l.printf(LevelPanic, format, args...)
	panic(fmt.Sprintf(format, args...))
}

func Debug(msg string)   { l.print(LevelDebug, msg) }
func Info(msg string)    { l.print(LevelInfo, msg) }
func Success(msg string) { l.print(LevelSuccess, msg) }
func Warn(msg string)    { l.print(LevelWarn, msg) }
func Error(msg string)   { l.print(LevelError, msg) }
func Panic(msg string)   { l.print(LevelPanic, msg); panic(msg) }
