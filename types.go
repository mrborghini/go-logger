package gologger

type LogLevel string

const (
	info    LogLevel = "\033[34m[INFO"
	warning LogLevel = "\033[33m[WARNING"
	error   LogLevel = "\033[31m[ERROR"
	debug   LogLevel = "\033[32m[DEBUG"
)
