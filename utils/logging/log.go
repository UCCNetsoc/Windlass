package log

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bwmarrin/lit"
)

const (
	callDepth = 3
)

type Fields map[string]interface{}

type Entry struct {
	fields Fields
	err    error
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
	return &Entry{fields: f}
}

func (e *Entry) WithFields(f Fields) *Entry {
	if e.fields == nil {
		e.fields = make(Fields)
	}
	for k, v := range f {
		e.fields[k] = v
	}
	return e
}

func (f Fields) format() string {
	if f == nil || len(f) == 0 {
		return ""
	}

	keys := make([]string, 0, len(f))
	for k := range f {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	var builder strings.Builder
	builder.WriteRune('\t')

	var iterCount int
	for _, k := range keys {
		v := f[k]
		s := fmt.Sprintf("%s=%v", k, v)
		if iterCount < len(f)-1 {
			s += " "
		}
		builder.WriteString(s)
		iterCount++
	}
	builder.WriteRune('\n')
	return builder.String()
}

func Debug(format string, a ...interface{}) {
	emptyEntry.Debug(format, a...)
}

func (e *Entry) Debug(format string, a ...interface{}) {
	builder := new(strings.Builder)
	lit.Custom(builder, lit.LogDebug, callDepth, format, a...)
	builder.WriteString(e.fields.format())
	fmt.Print(builder.String())
}

func Info(format string, a ...interface{}) {
	emptyEntry.Info(format, a...)
}

func (e *Entry) Info(format string, a ...interface{}) {
	builder := new(strings.Builder)
	lit.Custom(builder, lit.LogInformational, callDepth, format, a...)
	builder.WriteString(e.fields.format())
	fmt.Print(builder.String())
}

func Warn(format string, a ...interface{}) {
	emptyEntry.Warn(format, a...)
}

func (e *Entry) Warn(format string, a ...interface{}) {
	builder := new(strings.Builder)
	lit.Custom(builder, lit.LogWarning, callDepth, format, a...)
	builder.WriteString(e.fields.format())
	fmt.Print(builder.String())
}

func Error(format string, a ...interface{}) {
	emptyEntry.Error(format, a...)
}

func (e *Entry) Error(format string, a ...interface{}) {
	if e.err != nil {
		e.fields["error"] = e.err
	}
	builder := new(strings.Builder)
	lit.Custom(builder, lit.LogError, callDepth, format, a...)
	builder.WriteString(e.fields.format())
	fmt.Print(builder.String())
}

func WithError(err error) *Entry {
	return &Entry{err: err}
}

func (e *Entry) WithError(err error) *Entry {
	e.err = err
	return e
}
