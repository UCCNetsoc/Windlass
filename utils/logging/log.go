package log

import (
	"github.com/bwmarrin/lit"
)

const (
	callDepth = 3
)

type Fields map[string]interface{}

type Entry struct {
	fields Fields
}

func init() {
	lit.Prefix = ""
	lit.PrefixError = "ERROR"
	lit.PrefixDebug = "DEBUG"
	lit.PrefixWarning = "WARN "
	lit.PrefixInformational = "INFO "
	lit.LogLevel = lit.LogDebug
}

func WithFields(f Fields) *Entry {
	return &Entry{f}
}

func Debug(format string, a ...interface{}) {
	entry := &Entry{}
	entry.Debug(format, a...)
}

func (e *Entry) Debug(format string, a ...interface{}) {
	lit.Custom(lit.Writer, lit.LogDebug, callDepth, format, a...)
}

func Info(format string, a ...interface{}) {
	entry := &Entry{}
	entry.Info(format, a...)
}

func (e *Entry) Info(format string, a ...interface{}) {
	lit.Custom(lit.Writer, lit.LogInformational, callDepth, format, a...)
}

func Warn(format string, a ...interface{}) {
	entry := &Entry{}
	entry.Warn(format, a...)
}

func (e *Entry) Warn(format string, a ...interface{}) {
	lit.Custom(lit.Writer, lit.LogWarning, callDepth, format, a...)
}

func Error(format string, a ...interface{}) {
	entry := &Entry{}
	entry.Error(format, a...)
}

func (e *Entry) Error(format string, a ...interface{}) {
	lit.Custom(lit.Writer, lit.LogError, callDepth, format, a...)
}
