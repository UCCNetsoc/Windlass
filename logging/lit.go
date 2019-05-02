package log

import (
	"github.com/bwmarrin/lit"
)

const (
	callDepth = 3
)

func init() {
	lit.LogLevel = lit.LogDebug
}

func Debug(format string, a ...interface{}) {
	if lit.LogLevel == lit.LogDebug {
		lit.Custom(lit.Writer, lit.LogDebug, callDepth, format, a...)
	}
}

func Info(format string, a ...interface{}) {
	if lit.LogLevel >= lit.LogInformational {
		lit.Custom(lit.Writer, lit.LogInformational, callDepth, format, a...)
	}
}

func Warn(format string, a ...interface{}) {
	if lit.LogLevel >= lit.LogWarning {
		lit.Custom(lit.Writer, lit.LogWarning, callDepth, format, a...)
	}
}

func Error(format string, a ...interface{}) {
	if lit.LogLevel >= lit.LogError {
		lit.Custom(lit.Writer, lit.LogError, callDepth, format, a...)
	}
}
