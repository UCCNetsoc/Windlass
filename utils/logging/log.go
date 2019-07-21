package log

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/lit"
)

const (
	callDepth = 3
)

type Fields map[string]interface{}

type Entry struct {
	fields Fields
}

var emptyEntry = &Entry{}

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

func (f Fields) format() string {
	if f == nil {
		return ""
	}

	var builder strings.Builder
	builder.WriteRune('\t')
	for k, v := range f {
		builder.WriteString(fmt.Sprintf("%s=%v", k, v))
	}
	builder.WriteRune('\n')
	return builder.String()
}

func Debug(format string, a ...interface{}) {
	emptyEntry.Debug(format, a...)
}

func (e *Entry) Debug(format string, a ...interface{}) {
	lit.Custom(os.Stdout, lit.LogDebug, callDepth, format, a...)
	fmt.Print(e.fields.format())
}

func Info(format string, a ...interface{}) {
	emptyEntry.Info(format, a...)
}

func (e *Entry) Info(format string, a ...interface{}) {
	lit.Custom(os.Stdout, lit.LogInformational, callDepth, format, a...)
	fmt.Print(e.fields.format())
}

func Warn(format string, a ...interface{}) {
	emptyEntry.Warn(format, a...)
}

func (e *Entry) Warn(format string, a ...interface{}) {
	lit.Custom(os.Stdout, lit.LogWarning, callDepth, format, a...)
	fmt.Print(e.fields.format())
}

func Error(err error, format string, a ...interface{}) {
	emptyEntry.Error(err, format, a...)
}

func (e *Entry) Error(err error, format string, a ...interface{}) {
	format += fmt.Sprintf(": %v", err)
	lit.Custom(os.Stderr, lit.LogError, callDepth, format, a...)
	fmt.Print(e.fields.format())
}
