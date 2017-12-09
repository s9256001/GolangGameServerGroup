package servercommon

import (
	"fmt"

	"../base/logger"
)

const (
	// TracePrefix is the prefix of trace output
	TracePrefix = "TRACE"
	// DebugPrefix is the prefix of trace output
	DebugPrefix = "DEBUG"
	// InfoPrefix is the prefix of trace output
	InfoPrefix = "INFO"
	// ErrorPrefix is the prefix of trace output
	ErrorPrefix = "ERROR"
)

// ConsoleGameLogger is the game logger outputted to the console
type ConsoleGameLogger struct {
	*logger.GameLogger // base class
}

// OnLog is the function implementing the common underlying output operation
func (l *ConsoleGameLogger) OnLog(prefix string, message string) {
	message = fmt.Sprintf("| %-5s | %s", prefix, message)
	fmt.Printf(message)
}

// OnTrace is called when Trace()
func (l *ConsoleGameLogger) OnTrace(message string) {
	l.OnLog(TracePrefix, message)
}

// OnDebug is called when Debug()
func (l *ConsoleGameLogger) OnDebug(message string) {
	l.OnLog(DebugPrefix, message)
}

// OnInfo is called when Info()
func (l *ConsoleGameLogger) OnInfo(message string) {
	l.OnLog(InfoPrefix, message)
}

// OnError is called when Error()
func (l *ConsoleGameLogger) OnError(message string) {
	l.OnLog(ErrorPrefix, message)
}

// NewConsoleGameLogger is a constructor of ConsoleGameLogger
func NewConsoleGameLogger() *ConsoleGameLogger {
	ret := &ConsoleGameLogger{}
	ret.GameLogger = logger.NewGameLogger(ret)
	return ret
}
