package gologger

type LogLevel string

const (
	Info    LogLevel = "\033[34m[INFO"
	Warning LogLevel = "\033[33m[WARNING"
	Error   LogLevel = "\033[31m[ERROR"
	Debug   LogLevel = "\033[32m[DEBUG"
)
