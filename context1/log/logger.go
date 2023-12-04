package log

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type Logger struct {
	level level
}

type level int

const (
	LevelDebug = level(iota)
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic

	// LevelUnknown specifies a level that the logger won't handle
	LevelUnknown = level(-255)
)

func NewConsoleLogger(level string) *Logger {
	l := &Logger{
		level: LevelFromString(level),
	}

	if level == "" {
		l.level = LevelInfo // hardcoded default value
	}

	if l.level == LevelUnknown {
		panic(fmt.Errorf("unknown log level '%v'", level))
	}

	return l
}

func (l *Logger) SetLevel(level string) error {
	lvl := LevelFromString(level)
	if lvl == LevelUnknown {
		return errors.New("unknown level")
	}

	l.level = lvl
	return nil
}

func (l *Logger) Debug(msg string, values ...any) {
	l.log(LevelDebug, msg, values)
}

func (l *Logger) Info(msg string, values ...any) {
	l.log(LevelInfo, msg, values)
}

func (l *Logger) Warn(msg string, values ...any) {
	l.log(LevelWarn, msg, values)
}

func (l *Logger) Error(msg string, values ...any) {
	l.log(LevelError, msg, values)
}

func (l *Logger) Fatal(msg string, values ...any) {
	l.log(LevelFatal, msg, values)
}

func (l *Logger) Panic(msg string, values ...any) {
	l.log(LevelPanic, msg, values)
}

func (l *Logger) log(level level, msg string, values ...any) {
	if l.level > level {
		return
	}

	msgExpand := msg
	if len(values) > 0 {
		msgExpand = fmt.Sprintf(msg, values)
	}

	log.Println(msgExpand)

	if level >= LevelPanic {
		os.Exit(1)
	}
}

const (
	debugLevel = "debug"
	infoLevel  = "info"
	warnLevel  = "warn"
	errorLevel = "error"
	fatalLevel = "fatal"
	panicLevel = "panic"
)

func LevelFromString(lvl string) level {
	switch strings.ToLower(lvl) {
	case debugLevel:
		return LevelDebug
	case infoLevel:
		return LevelInfo
	case warnLevel:
		return LevelWarn
	case errorLevel:
		return LevelError
	case fatalLevel:
		return LevelFatal
	case panicLevel:
		return LevelPanic
	}
	return LevelUnknown
}

const ContextKey = "__logger__"

func FromContext(ctx context.Context) *Logger {
	return ctx.Value(ContextKey).(*Logger)
}

func (l *Logger) ContextWithValue(ctx context.Context) context.Context {
	return context.WithValue(ctx, ContextKey, l)
}
