package oops

import (
	"fmt"
)

// a simple logging interface with a Printf method.
type Logger interface {
	Printf(string, ...interface{})
}

// a structured logging interface with different log levels.
type LeveledLogger interface {
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Warn(string, ...interface{})
	Error(string, ...interface{})
}

// print logs the error with a specified prefix using the provided logger.
// handles both Logger and default fmt.Printf cases.
func (e *Error) print(logger interface{}, prefix string) {
	switch v := logger.(type) {
	case Logger:
		v.Printf("[%s] %s: %v", prefix, e.Message, e.Args())
	default:
		fmt.Printf("[%s] %s: %v\n", prefix, e.Message, e.Args())
	}
}

// log the error using a custom logging function.
// the function should accept a message string followed by variadic arguments.
func (e *Error) Log(fn func(string, ...any)) {
	fn(e.Message, e.Args()...)
}

// log the error at debug level using the provided logger.
// if the logger implements LeveledLogger, it uses the Debug method;
// otherwise, it falls back to the print method with "DEBUG" prefix.
func (e *Error) LogDebug(logger interface{}) {
	switch v := logger.(type) {
	case LeveledLogger:
		v.Debug(e.Message, e.Args()...)
	default:
		e.print(v, "DEBUG")
	}
}

// log the error at info level using the provided logger.
// if the logger implements LeveledLogger, it uses the Info method;
// otherwise, it falls back to the print method with "INFO" prefix.
func (e *Error) LogInfo(logger interface{}) {
	switch v := logger.(type) {
	case LeveledLogger:
		v.Info(e.Message, e.Args()...)
	default:
		e.print(v, "INFO")
	}
}

// log the error at warning level using the provided logger.
// if the logger implements LeveledLogger, it uses the Warn method;
// otherwise, it falls back to the print method with "WARN" prefix.
func (e *Error) LogWarn(logger interface{}) {
	switch v := logger.(type) {
	case LeveledLogger:
		v.Warn(e.Message, e.Args()...)
	default:
		e.print(v, "WARN")
	}
}

// log the error at error level using the provided logger.
// if the logger implements LeveledLogger, it uses the Error method;
// otherwise, it falls back to the print method with "ERROR" prefix.
func (e *Error) LogError(logger interface{}) {
	switch v := logger.(type) {
	case LeveledLogger:
		v.Error(e.Message, e.Args()...)
	default:
		e.print(v, "ERROR")
	}
}
