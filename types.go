package logger

type logLevel string

const (
	info    logLevel = "\033[34m[INFO"
	warning logLevel = "\033[33m[WARNING"
	error   logLevel = "\033[31m[ERROR"
	debug   logLevel = "\033[32m[DEBUG"
)
