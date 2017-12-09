package logger

import (
	"fmt"

	"../ginterface"
)

// GameLogger is an abstract class of the game logger
// It fits the interface of IGameLogger, and depends on the derived type to implement the interface of IGameLoggerHook
type GameLogger struct {
	ginterface.IGameLoggerHook // hook

	IsTraceEnabled bool // is the trace output enabled?
	IsDebugEnabled bool // is the debug output enabled?
	IsInfoEnabled  bool // is the information output enabled?
	IsErrorEnabled bool // is the error output enabled?
}

// Trace provides formatted output of trace
func (l *GameLogger) Trace(format string, v ...interface{}) {
	if l.IsTraceEnabled == false {
		return
	}
	message := fmt.Sprintf(format, v...)
	l.OnTrace(message)
}

// Debug provides formatted output of debug
func (l *GameLogger) Debug(format string, v ...interface{}) {
	if l.IsDebugEnabled == false {
		return
	}
	message := fmt.Sprintf(format, v...)
	l.OnDebug(message)
}

// Info provides formatted output of information
func (l *GameLogger) Info(format string, v ...interface{}) {
	if l.IsInfoEnabled == false {
		return
	}
	message := fmt.Sprintf(format, v...)
	l.OnInfo(message)
}

// Error provides formatted output of error
func (l *GameLogger) Error(format string, v ...interface{}) {
	if l.IsErrorEnabled == false {
		return
	}
	message := fmt.Sprintf(format, v...)
	l.OnError(message)
}

// NewGameLogger is a constructor of GameLogger
func NewGameLogger(hook ginterface.IGameLoggerHook) *GameLogger {
	ret := &GameLogger{
		IGameLoggerHook: hook,

		IsTraceEnabled: true,
		IsDebugEnabled: true,
		IsInfoEnabled:  true,
		IsErrorEnabled: true,
	}
	return ret
}
