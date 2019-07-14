package log

import (
	"encoding/json"
	"os"

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
	if e.fields != nil {
		b, _ := json.Marshal(e.fields)
		format += "\n\t" + string(b)
	}
	lit.Custom(os.Stdout, lit.LogDebug, callDepth, format, a...)
}

func Info(format string, a ...interface{}) {
	entry := &Entry{}
	entry.Info(format, a...)
}

func (e *Entry) Info(format string, a ...interface{}) {
	if e.fields != nil {
		b, _ := json.Marshal(e.fields)
		format += "\n\t" + string(b)
	}
	lit.Custom(os.Stdout, lit.LogInformational, callDepth, format, a...)
}

func Warn(format string, a ...interface{}) {
	entry := &Entry{}
	entry.Warn(format, a...)
}

func (e *Entry) Warn(format string, a ...interface{}) {
	if e.fields != nil {
		b, _ := json.Marshal(e.fields)
		format += "\n\t" + string(b)
	}
	lit.Custom(os.Stdout, lit.LogWarning, callDepth, format, a...)
}

func Error(format string, a ...interface{}) {
	entry := &Entry{}
	entry.Error(format, a...)
}

func (e *Entry) Error(format string, a ...interface{}) {
	if e.fields != nil {
		b, _ := json.Marshal(e.fields)
		format += "\n\t" + string(b)
	}
	lit.Custom(os.Stderr, lit.LogError, callDepth, format, a...)
}
