package ginterface

// IGameLogger is an interface of the game logger
type IGameLogger interface {
	// Trace provides formatted output of trace
	Trace(format string, v ...interface{})
	// Debug provides formatted output of debug
	Debug(format string, v ...interface{})
	// Info provides formatted output of information
	Info(format string, v ...interface{})
	// Error provides formatted output of error
	Error(format string, v ...interface{})
}

// IGameLoggerHook is an interface of hook of the game logger
type IGameLoggerHook interface {
	// OnTrace is called when Trace()
	OnTrace(message string)
	// OnDebug is called when Debug()
	OnDebug(message string)
	// OnInfo is called when Info()
	OnInfo(message string)
	// OnError is called when Error()
	OnError(message string)
}
