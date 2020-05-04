package domain

type Logger interface {
	Trace() string
	Debug(message string, payload interface{})
	Info(message string, payload interface{})
	Notice(message string, payload interface{})
	Warning(message string, payload interface{})
	Error(message string, payload interface{})
	Critical(message string, payload interface{})
	Alert(message string, payload interface{})
	Emergency(message string, payload interface{})
}
