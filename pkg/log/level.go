package log

//go:generate stringer -type LogLevel
const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

type LogLevel uint8
